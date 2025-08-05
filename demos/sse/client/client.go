package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// SSEClient represents a simple SSE client
type SSEClient struct {
	URL        string
	httpClient *http.Client
}

// NewSSEClient creates a new SSE client
func NewSSEClient(url string) *SSEClient {
	return &SSEClient{
		URL: url,
		httpClient: &http.Client{
			Timeout: 0, // No timeout for SSE connections
		},
	}
}

// Connect connects to the SSE endpoint and starts listening for events
func (c *SSEClient) Connect() error {
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set SSE headers
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	log.Printf("Connected to SSE server at %s", c.URL)

	// Read events
	scanner := bufio.NewScanner(resp.Body)
	var eventType, eventID, eventData strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// Empty line indicates end of event
			if eventData.Len() > 0 {
				c.handleEvent(eventType.String(), eventID.String(), eventData.String())
			}
			eventType.Reset()
			eventID.Reset()
			eventData.Reset()
			continue
		}

		if strings.HasPrefix(line, "event:") {
			eventType.WriteString(strings.TrimSpace(line[6:]))
		} else if strings.HasPrefix(line, "id:") {
			eventID.WriteString(strings.TrimSpace(line[3:]))
		} else if strings.HasPrefix(line, "data:") {
			if eventData.Len() > 0 {
				eventData.WriteString("\n")
			}
			eventData.WriteString(strings.TrimSpace(line[5:]))
		} else if strings.HasPrefix(line, "retry:") {
			// Handle retry if needed
			log.Printf("Server suggests retry: %s", strings.TrimSpace(line[6:]))
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	return nil
}

// handleEvent processes received events
func (c *SSEClient) handleEvent(eventType, eventID, data string) {
	timestamp := time.Now().Format("15:04:05")

	if eventType == "" {
		eventType = "message"
	}

	fmt.Printf("[%s] Event: %s", timestamp, eventType)
	if eventID != "" {
		fmt.Printf(" (ID: %s)", eventID)
	}
	fmt.Printf("\nData: %s\n\n", data)
}

func main() {
	serverURL := "http://localhost:8080/events"

	log.Printf("Starting SSE client...")
	log.Printf("Connecting to: %s", serverURL)

	client := NewSSEClient(serverURL)

	// Connect with retry logic
	for {
		err := client.Connect()
		if err != nil {
			log.Printf("Connection failed: %v", err)
			log.Printf("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		// If we reach here, connection was closed
		log.Printf("Connection closed, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}
