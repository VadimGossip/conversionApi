package config

import (
	"github.com/VadimGossip/conversionApi/internal/domain"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func setFromEnv(cfg *domain.Config) error {
	if err := envconfig.Process("api_http", &cfg.ApiHttpServer); err != nil {
		return err
	}
	if err := envconfig.Process("api_metrics_http", &cfg.ApiMetricsHttpServer); err != nil {
		return err
	}
	if err := envconfig.Process("ampq_server", &cfg.AMPQServerConfig); err != nil {
		return err
	}

	return nil
}

func Init() (*domain.Config, error) {

	var cfg domain.Config
	if err := setFromEnv(&cfg); err != nil {
		return nil, err
	}
	//temp
	cfg.AMPQServerConfig.ConvQueueName = "ConvWorkQueue"
	cfg.AMPQServerConfig.Url = "amqp://guest:guest@localhost:5672/"
	cfg.ApiHttpServer.Port = 8080
	cfg.ApiMetricsHttpServer.Port = 9090

	logrus.Infof("Config %v", cfg)
	return &cfg, nil
}
