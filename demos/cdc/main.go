package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/segmentio/kafka-go"
	"log"
)

type OrderEvent struct {
	Payload struct {
		After *struct {
			ID       int    `json:"id"`
			Item     string `json:"item"`
			Quantity int    `json:"quantity"`
			Status   string `json:"status"`
		} `json:"after"`
		Op string `json:"op"`
	} `json:"payload"`
}

func main() {
	ctx := context.Background()

	es, err := elasticsearch.NewDefaultClient() //11
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Kafka consumer
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"host.docker.internal:9092"},
		Topic:    "pgserver1.public.orders",
		GroupID:  "go-consumer",
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	fmt.Println("Listening for Kafka messages...")

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("could not read message: %v", err)
		}

		var evt OrderEvent
		if err := json.Unmarshal(m.Value, &evt); err != nil {
			log.Printf("unmarshal error: %v", err)
			continue
		}

		if evt.Payload.After != nil {
			docID := fmt.Sprintf("%d", evt.Payload.After.ID)
			body, _ := json.Marshal(evt.Payload.After)

			// index to ES
			res, err := es.Index(
				"orders",
				bytes.NewReader(body),
				es.Index.WithDocumentID(docID),
				es.Index.WithRefresh("true"),
			)
			if err != nil {
				log.Printf("ES index error: %v", err)
				continue
			}
			res.Body.Close()

			fmt.Printf("Indexed order %s to ES\n", docID)
		} else if evt.Payload.Op == "d" {
			// handle delete in ES
		}
	}
}
