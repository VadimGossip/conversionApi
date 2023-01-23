package app

import "github.com/VadimGossip/conversionApi/internal/conversion"

type Factory struct {
	queueAdapter *QueueAdapter

	queueService conversion.Service
}

var factory *Factory

func newFactory(queueAdapter *QueueAdapter) *Factory {
	factory = &Factory{queueAdapter: queueAdapter}
	factory.queueService = conversion.NewService(factory.queueAdapter.convQueueChan)
	return factory
}
