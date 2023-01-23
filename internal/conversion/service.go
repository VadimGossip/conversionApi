package conversion

import "github.com/streadway/amqp"

type Service interface {
	PublishMsg() error
}

type service struct {
	convQueueChan *amqp.Channel
}

var _ Service = (*service)(nil)

func NewService(convQueueChan *amqp.Channel) *service {
	return &service{convQueueChan: convQueueChan}
}

func (s *service) PublishMsg() error {
	return nil
}
