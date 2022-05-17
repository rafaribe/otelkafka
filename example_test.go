package otelkafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ExampleNewConsumer() {
	c, err := NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()
}

func ExampleNewProducer() {
	p, err := NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()
}
