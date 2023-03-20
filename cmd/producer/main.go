package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/thiagoCRDM/gokafka/packages/producer"
)

func main() {

	deliveryChan := make(chan kafka.Event)

	produ := producer.NewKafkaProducer()
	producer.Publish("Transferred", "teste", produ, []byte("transfer"), deliveryChan)

	go producer.DeliveryReport(deliveryChan)

	// e := <-deliveryChan

	// msg := e.(*kafka.Message)

	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("Error to send")
	// } else {
	// 	fmt.Println("Message sent successfully", msg.TopicPartition)
	// }
}
