package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

var (
	topic   = "123"
	groupID = "my-group"
)

//func main() {
//	var wg sync.WaitGroup
//
//	// Запуск продюсера
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		producer := NewKafkaProducer([]string{"localhost:9092"}, topic)
//		defer producer.Close()
//
//		for i := 0; i < 10; i++ {
//			message := fmt.Sprintf("message %d", i)
//			if err := producer.SendMessage(message); err != nil {
//				log.Println("не удалось отправить сообщение:", err)
//			}
//			time.Sleep(1 * time.Second) // имитируем задержку между сообщениями
//		}
//	}()
//
//	// Запуск консьюмера
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		consumer := NewKafkaConsumer([]string{"localhost:9092"}, topic, groupID)
//		defer consumer.Close()
//
//		consumer.ConsumeMessages()
//	}()
//
//	wg.Wait()
//}

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
func (c *KafkaConsumer) ConsumeMessages() {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("не удалось прочитать сообщение:", err)
			continue
		}
		fmt.Printf("получено: %s\n", string(msg.Value))
	}
}

// Close - закрывает консьюмер Kafka
func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}
