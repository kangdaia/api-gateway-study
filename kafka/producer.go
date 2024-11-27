package kafka

import (
	"api-gateway-study/config"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	_allAcks = "all"
)

type Producer struct {
	cfg      config.Producer
	producer *kafka.Producer
}

func NewProducer(cfg config.Producer) Producer {
	url := cfg.URL
	id := cfg.ClientID
	acks := cfg.Acks

	if acks == "" {
		acks = _allAcks
	}

	conf := &kafka.ConfigMap{
		"bootstrap.servers": url,
		"client.id":         id,
		"acks":              acks,
	}

	producer, err := kafka.NewProducer(conf)
	if err != nil {
		panic(err.Error())
	}

	return Producer{
		cfg:      cfg,
		producer: producer,
	}
}

func (p Producer) SendEvent(v []byte) {
	topic := p.cfg.Topic
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: v,
	}, nil)
	if err != nil {
		log.Println("Failed to send topic", string(v))
	} else {
		log.Println("Success to send topic", string(v))
	}
}
