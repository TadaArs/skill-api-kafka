package skill

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
	handler *SkillHandler
}

type kafkaMsg struct {
	Action string `json:"action"`
	Key    string `json:"key,omitempty"`
	Data   Skill    `json:"data,omitempty"`
}

func NewConsumer(topic string, storage SkillStorage) *Consumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer([]string{os.Getenv("BROKERS")}, config)
	if err != nil {
		log.Fatalln(err)
	}

	return &Consumer{consumer: consumer, handler: NewSkillHandler(&storage)}
}



func (c *Consumer) Consume(topic string) error {
	partitionConsumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln(err)
	}
	defer partitionConsumer.Close()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var message kafkaMsg
			if err := json.Unmarshal(msg.Value, &message); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}
			c.handler.ExtractMsg(message);
			consumed++
			log.Printf("Consumed message: %s", msg.Value)
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %v", err)
		case <-signals:
			log.Println("Interrupt is detected")
			break ConsumerLoop
		}
	}
	log.Printf("Consumed: %d messages", consumed)
	return nil
}


func (c *Consumer) Close() {
	if err := c.consumer.Close(); err != nil {
		log.Printf("Failed to close consumer: %v", err)
	}
}