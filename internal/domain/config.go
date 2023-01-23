package domain

type NetServerConfig struct {
	Port int
}

type AMPQServerConfig struct {
	Url           string
	ConvQueueName string
}

type Config struct {
	ApiHttpServer        NetServerConfig
	ApiMetricsHttpServer NetServerConfig
	AMPQServerConfig     AMPQServerConfig
}
