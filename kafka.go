package main

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func NewAsyncEventProducer(brokerList []string) sarama.AsyncProducer {
	// Use async producer for high throughput (AP)
	// Message loss a possiblity.
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for leader to ack
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start kafka producer:", err)
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write event:", err)
		}
	}()

	return producer
}
