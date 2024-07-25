package skill

import (
	"log"
	"os"

	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer() *Producer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{os.Getenv("BROKERS")}, config)
	if err != nil {
		log.Fatalln(err)
	}
	return &Producer{producer: producer}
}

func (p *Producer) Publish(topic string, msg []byte) error {
	bytemsg := &sarama.ProducerMessage{Topic: topic, Value: sarama.ByteEncoder(string(msg))}
	_, _, err := p.producer.SendMessage(bytemsg)
	if err != nil {
		return err
	}
	return nil
}
