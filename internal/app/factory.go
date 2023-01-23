package app

import (
	"github.com/VadimGossip/conversionApi/internal/conversion"
	"github.com/VadimGossip/conversionApi/internal/domain"
)

type Factory struct {
	queueAdapter *QueueAdapter

	convService conversion.Service
}

var factory *Factory

func newFactory(cfg *domain.Config, queueAdapter *QueueAdapter) *Factory {
	factory = &Factory{queueAdapter: queueAdapter}
	factory.convService = conversion.NewService(cfg.AMPQServerConfig.ConvQueueName, factory.queueAdapter.convQueueChan)
	return factory
}
