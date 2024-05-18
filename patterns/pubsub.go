package patterns

import (
	"fmt"
	"sync"
)

type Message struct {
	Topic string
	Data  interface{}
}

type Publisher struct {
	Data chan<- Message
}

func (p *Publisher) Publish(topic string, data interface{}) {
	p.Data <- Message{Topic: topic, Data: data}
}

type Subscriber struct {
	Topic string
	Data  <-chan Message
}

func (s *Subscriber) Subscribe(wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range s.Data {
		if msg.Topic == s.Topic {
			fmt.Printf("Subscriber received: %v\n", msg.Data)
		}
	}
}

type PubHub struct {
	channels chan Message
}

func NewPubHub() *PubHub {
	return &PubHub{
		channels: make(chan Message),
	}
}

func (ph *PubHub) RegisterPublisher() *Publisher {
	ch := make(chan Message)
	ph.channels = ch
	return &Publisher{Data: ch}
}

func (ph *PubHub) RegisterSubscriber(topic string) *Subscriber {
	return &Subscriber{Topic: topic, Data: ph.channels}
}

func (ph *PubHub) Close() {
	close(ph.channels)
}
