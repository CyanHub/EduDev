package main

import (
	"log"

	"github.com/IBM/sarama"
)
func main() {
	producer, err := sarama.NewSyncProducer([]string{"47.98.239.196:9096", "47.98.239.197:9092", "47.98.239.198:9094"}, nil)
	if err!= nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "sync-topic",
		Value: sarama.StringEncoder("hello kafka"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("failed to send message: %s", err)
		return
	}
	println("Message sent to partition", partition, "at offset", offset)
}