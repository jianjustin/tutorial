package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// SSEEvent represents a server-sent event
type SSEEvent struct {
	ID    string      `json:"id,omitempty"`
	Event string      `json:"event,omitempty"`
	Data  interface{} `json:"data"`
	Retry int         `json:"retry,omitempty"`
}

// Client represents a connected SSE client
type Client struct {
	ID       string
	Channel  chan SSEEvent
	Request  *http.Request
	Writer   http.ResponseWriter
	LastSeen time.Time
}

// SSEHub manages all connected clients and broadcasts events
type SSEHub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan SSEEvent
	mutex      sync.RWMutex
}

// NewSSEHub creates a new SSE hub
func NewSSEHub() *SSEHub {
	return &SSEHub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan SSEEvent),
	}
}

// Run starts the SSE hub
func (h *SSEHub) Run() {
	go func() {
		for {
			select {
			case client := <-h.register:
				h.mutex.Lock()
				h.clients[client.ID] = client
				h.mutex.Unlock()
				log.Printf("Client %s connected. Total clients: %d", client.ID, len(h.clients))

				// Send welcome message
				welcomeEvent := SSEEvent{
					ID:    fmt.Sprintf("%d", time.Now().UnixNano()),
					Event: "connected",
					Data:  map[string]string{"message": "Connected to SSE server", "clientId": client.ID},
				}
				select {
				case client.Channel <- welcomeEvent:
				default:
					// Client channel is full, skip
				}

			case client := <-h.unregister:
				h.mutex.Lock()
				if _, ok := h.clients[client.ID]; ok {
					delete(h.clients, client.ID)
					close(client.Channel)
				}
				h.mutex.Unlock()
				log.Printf("Client %s disconnected. Total clients: %d", client.ID, len(h.clients))

			case event := <-h.broadcast:
				h.mutex.RLock()
				for _, client := range h.clients {
					select {
					case client.Channel <- event:
					default:
						// Client channel is full, remove client
						delete(h.clients, client.ID)
						close(client.Channel)
					}
				}
				h.mutex.RUnlock()
				log.Printf("Broadcasted event '%s' to %d clients", event.Event, len(h.clients))
			}
		}
	}()
}

// BroadcastEvent sends an event to all connected clients
func (h *SSEHub) BroadcastEvent(event SSEEvent) {
	select {
	case h.broadcast <- event:
	default:
		log.Println("Broadcast channel is full, dropping event")
	}
}

// GetClientCount returns the number of connected clients
func (h *SSEHub) GetClientCount() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.clients)
}

// formatSSEEvent formats an event for SSE transmission
func formatSSEEvent(event SSEEvent) string {
	var result string

	if event.ID != "" {
		result += fmt.Sprintf("id: %s\n", event.ID)
	}

	if event.Event != "" {
		result += fmt.Sprintf("event: %s\n", event.Event)
	}

	if event.Retry > 0 {
		result += fmt.Sprintf("retry: %d\n", event.Retry)
	}

	// Format data as JSON
	dataBytes, err := json.Marshal(event.Data)
	if err != nil {
		dataBytes = []byte(fmt.Sprintf(`{"error": "Failed to marshal data: %v"}`, err))
	}

	result += fmt.Sprintf("data: %s\n\n", string(dataBytes))

	return result
}

// sseHandler handles SSE connections
func (h *SSEHub) sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Cache-Control")

	// Create client
	clientID := fmt.Sprintf("client_%d", time.Now().UnixNano())
	client := &Client{
		ID:       clientID,
		Channel:  make(chan SSEEvent, 10), // Buffer for 10 events
		Request:  r,
		Writer:   w,
		LastSeen: time.Now(),
	}

	// Register client
	h.register <- client

	// Ensure client is unregistered when connection closes
	defer func() {
		h.unregister <- client
	}()

	// Handle client disconnection
	notify := r.Context().Done()

	// Send events to client
	for {
		select {
		case event := <-client.Channel:
			// Write event to client
			eventData := formatSSEEvent(event)
			if _, err := fmt.Fprint(w, eventData); err != nil {
				log.Printf("Error writing to client %s: %v", client.ID, err)
				return
			}

			// Flush the data immediately
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}

			client.LastSeen = time.Now()

		case <-notify:
			// Client disconnected
			log.Printf("Client %s disconnected", client.ID)
			return

		case <-time.After(30 * time.Second):
			// Send heartbeat
			heartbeat := SSEEvent{
				ID:    fmt.Sprintf("%d", time.Now().UnixNano()),
				Event: "heartbeat",
				Data:  map[string]interface{}{"timestamp": time.Now().Unix()},
			}

			select {
			case client.Channel <- heartbeat:
			default:
				// Channel is full, client might be slow
				log.Printf("Client %s channel is full, disconnecting", client.ID)
				return
			}
		}
	}
}

// Message represents a message to be sent
type Message struct {
	Event   string      `json:"event"`
	Data    interface{} `json:"data"`
	EventID string      `json:"eventId,omitempty"`
}

// sendMessageHandler handles sending messages via HTTP POST
func (h *SSEHub) sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create SSE event
	event := SSEEvent{
		ID:    message.EventID,
		Event: message.Event,
		Data:  message.Data,
	}

	if event.ID == "" {
		event.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	// Broadcast event
	h.BroadcastEvent(event)

	// Response
	response := map[string]interface{}{
		"success":     true,
		"message":     "Event sent successfully",
		"eventId":     event.ID,
		"clientCount": h.GetClientCount(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// statusHandler returns server status
func (h *SSEHub) statusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"status":      "running",
		"clientCount": h.GetClientCount(),
		"timestamp":   time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// indexHandler serves the HTML demo page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>SSE Demo</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 800px; margin: 0 auto; }
        .status { padding: 10px; margin: 10px 0; border-radius: 5px; }
        .connected { background-color: #d4edda; color: #155724; }
        .disconnected { background-color: #f8d7da; color: #721c24; }
        .events { border: 1px solid #ccc; height: 300px; overflow-y: auto; padding: 10px; margin: 10px 0; }
        .event { margin: 5px 0; padding: 5px; background-color: #f8f9fa; border-left: 3px solid #007bff; }
        .controls { margin: 20px 0; }
        input, button { margin: 5px; padding: 8px; }
        button { background-color: #007bff; color: white; border: none; border-radius: 3px; cursor: pointer; }
        button:hover { background-color: #0056b3; }
    </style>
</head>
<body>
    <div class="container">
        <h1>SSE (Server-Sent Events) Demo</h1>
        
        <div id="status" class="status disconnected">Disconnected</div>
        
        <div class="controls">
            <button onclick="connect()">Connect</button>
            <button onclick="disconnect()">Disconnect</button>
            <button onclick="clearEvents()">Clear Events</button>
        </div>
        
        <div class="controls">
            <input type="text" id="eventType" placeholder="Event Type" value="message">
            <input type="text" id="eventData" placeholder="Event Data" value="Hello World!">
            <button onclick="sendMessage()">Send Message</button>
        </div>
        
        <h3>Events:</h3>
        <div id="events" class="events"></div>
        
        <div class="controls">
            <button onclick="sendPredefinedMessages()">Send Test Messages</button>
        </div>
    </div>

    <script>
        let eventSource = null;
        let eventCount = 0;

        function updateStatus(connected, message) {
            const statusEl = document.getElementById('status');
            statusEl.className = 'status ' + (connected ? 'connected' : 'disconnected');
            statusEl.textContent = message || (connected ? 'Connected' : 'Disconnected');
        }

        function addEvent(event) {
            const eventsEl = document.getElementById('events');
            const eventEl = document.createElement('div');
            eventEl.className = 'event';
            
            const timestamp = new Date().toLocaleTimeString();
            eventEl.innerHTML = '<strong>' + timestamp + ' [' + (event.type || 'message') + ']:</strong> ' + 
                               (typeof event.data === 'string' ? event.data : JSON.stringify(event.data));
            
            eventsEl.appendChild(eventEl);
            eventsEl.scrollTop = eventsEl.scrollHeight;
            eventCount++;
        }

        function connect() {
            if (eventSource) {
                eventSource.close();
            }

            updateStatus(false, 'Connecting...');
            
            eventSource = new EventSource('/events');
            
            eventSource.onopen = function(event) {
                updateStatus(true, 'Connected');
                addEvent({type: 'system', data: 'Connected to server'});
            };
            
            eventSource.onmessage = function(event) {
                try {
                    const data = JSON.parse(event.data);
                    addEvent({type: 'message', data: data});
                } catch (e) {
                    addEvent({type: 'message', data: event.data});
                }
            };
            
            eventSource.addEventListener('connected', function(event) {
                try {
                    const data = JSON.parse(event.data);
                    addEvent({type: 'connected', data: data});
                } catch (e) {
                    addEvent({type: 'connected', data: event.data});
                }
            });
            
            eventSource.addEventListener('heartbeat', function(event) {
                try {
                    const data = JSON.parse(event.data);
                    addEvent({type: 'heartbeat', data: 'Heartbeat: ' + new Date(data.timestamp * 1000).toLocaleTimeString()});
                } catch (e) {
                    addEvent({type: 'heartbeat', data: event.data});
                }
            });
            
            eventSource.addEventListener('notification', function(event) {
                try {
                    const data = JSON.parse(event.data);
                    addEvent({type: 'notification', data: data});
                } catch (e) {
                    addEvent({type: 'notification', data: event.data});
                }
            });
            
            eventSource.onerror = function(event) {
                updateStatus(false, 'Connection error');
                addEvent({type: 'error', data: 'Connection error occurred'});
            };
        }

        function disconnect() {
            if (eventSource) {
                eventSource.close();
                eventSource = null;
            }
            updateStatus(false, 'Disconnected');
            addEvent({type: 'system', data: 'Disconnected from server'});
        }

        function clearEvents() {
            document.getElementById('events').innerHTML = '';
            eventCount = 0;
        }

        async function sendMessage() {
            const eventType = document.getElementById('eventType').value || 'message';
            const eventData = document.getElementById('eventData').value || 'Hello World!';
            
            try {
                const response = await fetch('/send', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        event: eventType,
                        data: eventData
                    })
                });
                
                const result = await response.json();
                if (result.success) {
                    addEvent({type: 'system', data: 'Message sent: ' + eventData});
                } else {
                    addEvent({type: 'error', data: 'Failed to send message'});
                }
            } catch (error) {
                addEvent({type: 'error', data: 'Error sending message: ' + error.message});
            }
        }

        async function sendPredefinedMessages() {
            const messages = [
                {event: 'notification', data: {title: 'System Alert', message: 'Server maintenance scheduled'}},
                {event: 'message', data: 'This is a test message'},
                {event: 'notification', data: {title: 'Update Available', message: 'New version 2.0 is available'}},
                {event: 'message', data: {user: 'admin', text: 'System is running smoothly'}},
                {event: 'notification', data: {title: 'Welcome', message: 'Welcome to SSE Demo!'}}
            ];
            
            for (let i = 0; i < messages.length; i++) {
                setTimeout(async () => {
                    try {
                        await fetch('/send', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify(messages[i])
                        });
                    } catch (error) {
                        console.error('Error sending message:', error);
                    }
                }, i * 1000);
            }
        }

        // Auto-connect when page loads
        window.onload = function() {
            connect();
        };

        // Cleanup when page unloads
        window.onbeforeunload = function() {
            if (eventSource) {
                eventSource.close();
            }
        };
    </script>
</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func main() {
	// Initialize SSE hub
	hub := NewSSEHub()
	hub.Run()

	// Setup routes
	r := mux.NewRouter()

	// SSE endpoint
	r.HandleFunc("/events", hub.sseHandler).Methods("GET")

	// API endpoints
	r.HandleFunc("/send", hub.sendMessageHandler).Methods("POST")
	r.HandleFunc("/status", hub.statusHandler).Methods("GET")

	// Serve demo page
	r.HandleFunc("/", indexHandler).Methods("GET")

	// Start background message generator (for demo purposes)
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		messageCount := 1
		for range ticker.C {
			event := SSEEvent{
				ID:    fmt.Sprintf("%d", time.Now().UnixNano()),
				Event: "notification",
				Data: map[string]interface{}{
					"title":     "Auto Message",
					"message":   fmt.Sprintf("This is automated message #%d", messageCount),
					"timestamp": time.Now().Unix(),
				},
			}
			hub.BroadcastEvent(event)
			messageCount++
		}
	}()

	// Start server
	port := ":8080"
	log.Printf("SSE server starting on port %s", port)
	log.Printf("Open http://localhost%s in your browser to see the demo", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
