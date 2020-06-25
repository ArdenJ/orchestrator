package queue

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func Sender() {
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

	q, err := ch.QueueDeclare(
		"Test Queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)

	// Create messages
	err = ch.Publish(
		"",
		"Test Queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello nerds"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
