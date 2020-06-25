package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// Queue instantiates a Queue to connect to
type Queue struct {
	l *log.Logger
}

// Connect returns a amqp connection (to be closed), a channel, and error
func (q *Queue) Connect(l *log.Logger) (*amqp.Connection, *amqp.Channel, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		l.Fatal(err)
	}

	ch, err := connection.Channel()
	if err != nil {
		l.Fatal(err)
	}

	select {}

	return connection, ch, err
}

// NewConnection returns a reference to a Queue on which the Connect method can be called
func NewConnection(l *log.Logger) *Queue {
	return &Queue{l}
}
