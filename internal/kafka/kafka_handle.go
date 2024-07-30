package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// KafkaProducer - структура для продюсера Kafka
type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer - создает новый продюсер Kafka
func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

// SendMessage - отправляет сообщение в Kafka
func (p *KafkaProducer) SendMessage(message string) error {
	return p.writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(message),
	})
}

// Close - закрывает продюсер Kafka
func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}

// KafkaConsumer - структура для консьюмера Kafka
type KafkaConsumer struct {
	reader *kafka.Reader
}

// NewKafkaConsumer - создает новый консьюмер Kafka
func NewKafkaConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
		}),
	}
}

// ConsumeMessages - потребляет сообщения из Kafka
func (c *KafkaConsumer) ConsumeMessages() (string, error) {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("не удалось прочитать сообщение:", err)
			continue
		}
		fmt.Printf("получено: %s\n", string(msg.Value))
		return string(msg.Value), nil
	}
}

// Close - закрывает консьюмер Kafka
func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}
