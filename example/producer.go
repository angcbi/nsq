package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:5000", config)
	if err != nil {
		log.Fatal(err)
	}

	topicName := "topic"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	for i := 0; i < 10; i++ {
		messageBody := []byte(fmt.Sprintf("hello %d", i))
		err = producer.Publish(topicName, messageBody)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 2)
	}


	// Gracefully stop the producer when appropriate (e.g. before shutting down the service)
	producer.Stop()
}