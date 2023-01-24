package app

import (
	"github.com/VadimGossip/conversionApi/internal/domain"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type QueueAdapter struct {
	cfg            *domain.Config
	convRabbitConn *amqp.Connection
	convQueueChan  *amqp.Channel
}

func NewQueueAdapter(cfg *domain.Config) *QueueAdapter {
	dba := &QueueAdapter{}
	dba.cfg = cfg
	return dba
}

func (q *QueueAdapter) Connect() error {
	connectRabbitMQ, err := amqp.Dial(q.cfg.AMPQServerConfig.Url)
	if err != nil {
		return err
	}
	q.convRabbitConn = connectRabbitMQ

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return err
	}
	q.convQueueChan = channelRabbitMQ

	if err = channelRabbitMQ.ExchangeDeclare(
		"ConvWorkExchange",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	_, err = channelRabbitMQ.QueueDeclare(
		q.cfg.AMPQServerConfig.ConvQueueName,
		true,
		false,
		false,
		false,
		amqp.Table{"x-dead-letter-exchange": "ConvRetryExchange"},
	)
	if err != nil {
		return err
	}

	err = channelRabbitMQ.QueueBind(
		q.cfg.AMPQServerConfig.ConvQueueName,
		"",
		"ConvWorkExchange",
		false,
		nil,
	)

	if err = channelRabbitMQ.ExchangeDeclare(
		"ConvRetryExchange",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	_, err = channelRabbitMQ.QueueDeclare(
		"ConvRetryQueue",
		true,
		false,
		false,
		false,
		amqp.Table{"x-dead-letter-exchange": "ConvWorkExchange", "x-message-ttl": 10000},
	)
	if err != nil {
		return err
	}

	err = channelRabbitMQ.QueueBind(
		"ConvRetryQueue",
		"",
		"ConvRetryExchange",
		false,
		nil,
	)

	return nil
}

func (q *QueueAdapter) Close() {
	if err := q.convQueueChan.Close(); err != nil {
		logrus.Errorf("Error occured on convQueueChan close: %s", err.Error())
	}

	if err := q.convRabbitConn.Close(); err != nil {
		logrus.Errorf("Error occured on convRabbitConn close: %s", err.Error())
	}
}
