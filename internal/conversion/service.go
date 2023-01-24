package conversion

import (
	"github.com/streadway/amqp"
	"math/rand"
	"strconv"
)

type Service interface {
	PublishMsg() error
}

type service struct {
	convQueueName string
	convQueueChan *amqp.Channel
}

var _ Service = (*service)(nil)

func NewService(convQueueName string, convQueueChan *amqp.Channel) *service {
	return &service{convQueueName: convQueueName, convQueueChan: convQueueChan}
}

func (s *service) PublishMsg() error {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("some message id = " + strconv.Itoa(rand.Int())),
	}

	// Attempt to publish a message to the queue.
	if err := s.convQueueChan.Publish(
		"ConvWorkExchange",
		"ConvWorkQueue",
		false,
		false,
		message,
	); err != nil {
		return err
	}

	return nil
}
