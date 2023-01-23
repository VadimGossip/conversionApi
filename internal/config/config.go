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

	logrus.Infof("Config %v", cfg)
	return &cfg, nil
}

//func init() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}
