package consumer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewConsumer() {
	configmap := &kafka.ConfigMap{
		"bootstrap.servers": "gokafka_kafka_1:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(configmap)

	if err != nil {
		fmt.Println("Erro Consumer", err.Error())
	}

	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
