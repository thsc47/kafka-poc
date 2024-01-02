package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {

}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-kafka-1:9092",
	}
	producer, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Println(err.Error())
	}
	return producer
}
