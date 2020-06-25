package queue

import (
	"fmt"
	"log"
	"os"
)

func Consumer() {
	// Logger
	l := log.New(os.Stdout, "api", log.LstdFlags)

	connection, ch, err := NewConnection(l).Connect(l)
	if err != nil {
		l.Fatal(err)
	}
	defer func() {
		err = connection.Close()
		if err != nil {
			l.Fatal(err)
		}
	}()

	msgs, err := ch.Consume(
		"Test Queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		l.Fatal(err)
	}

	forever := make(chan bool, 1)
	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to RMQ")
	fmt.Println("Waiting for messages...")
	<-forever
}
