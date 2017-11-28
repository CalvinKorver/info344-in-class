package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// This particular channel is read-only, we can not put things into it
// That is what the left arrow means (w/o means R & W access)
func listen(msgs <-chan amqp.Delivery) {
	log.Println("Listening for b o u n d l e s s messages")
	for msg := range msgs {
		log.Println(string(msg.Body))
	}

}

func main() {
	mqAddr := os.Getenv("MQADDR")
	if len(mqAddr) == 0 {
		mqAddr = "localhost:5672"
	}
	mqURL := fmt.Sprintf("amqp://%s", mqAddr)
	conn, err := amqp.Dial(mqURL)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("error creating channel: %v", err)
	}

	q, err := channel.QueueDeclare("testQ", false, false, false, false, nil)

	// In memory data channel that we can put stuff into and take out
	msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)

	// When we use go before a command, it starts a new goroutine
	go listen(msgs)

	// There is nothing in it
	neverEnd := make(chan bool)

	// I want to read a boolean out of this channel
	// If we try to read from a channel and there is nothin in the channel Go will block until we can read something out. So it will sit there and block
	<-neverEnd

}
