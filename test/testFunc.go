package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "my-topic"
	groupID := "my-group"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		GroupID: groupID,
	})

	defer func() {
		if err := reader.Close(); err != nil {
			log.Fatal("не удалось закрыть reader:", err)
		}
	}()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("не удалось прочитать сообщение:", err)
		}
		fmt.Printf("получено: %s\n", string(msg.Value))
	}
}
