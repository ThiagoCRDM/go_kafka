package producer

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewKafkaProducer() *kafka.Producer {
	configmap := &kafka.ConfigMap{
		"bootstrap.servers":   "gokafka_kafka_1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}

	p, err := kafka.NewProducer(configmap)

	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(message string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	msg := &kafka.Message{
		Value:          []byte(message),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(msg, deliveryChan)
	if err != nil {
		return err
	}
	producer.Flush(5000)
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Error to send")
			} else {
				fmt.Println("Message sent successfully", ev.TopicPartition)
			}
		}
	}
}
