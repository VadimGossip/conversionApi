package app

import (
	"github.com/VadimGossip/conversionApi/internal/config"
	"github.com/VadimGossip/conversionApi/internal/domain"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

type App struct {
	*Factory
	name              string
	appStartedAt      time.Time
	cfg               *domain.Config
	apiHttpServer     *HttpServer
	metricsHttpServer *HttpServer
}

func NewApp(name string, appStartedAt time.Time) *App {
	return &App{
		name:         name,
		appStartedAt: appStartedAt,
	}
}

func (app *App) Run() {
	cfg, err := config.Init()
	if err != nil {
		logrus.Fatalf("Config initialization error %s", err)
	}
	app.cfg = cfg

	qAdapter := NewQueueAdapter(app.cfg)
	if err := qAdapter.Connect(); err != nil {
		logrus.Fatalf("Fail to connect ampq %s", err)
	}

	app.Factory = newFactory(app.cfg, qAdapter)

	go func() {
		app.apiHttpServer = NewHttpServer(app.cfg.ApiHttpServer.Port)
		initHttpRouter(app)
		if err := app.apiHttpServer.Run(); err != nil {
			logrus.Fatalf("error occured while running drs http server: %s", err.Error())
		}
	}()

	go func() {
		app.metricsHttpServer = NewHttpServer(app.cfg.ApiMetricsHttpServer.Port)
		initMetricsHttpRouter(app)
		if err := app.metricsHttpServer.Run(); err != nil {
			logrus.Fatalf("error occured while running drs http server: %s", err.Error())
		}
	}()

	logrus.Print("Conversion api service started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := app.apiHttpServer.Shutdown(); err != nil {
		logrus.Errorf("Error occured on http server for conversion api shutting down: %s", err.Error())
	}

	if err := app.metricsHttpServer.Shutdown(); err != nil {
		logrus.Errorf("Error occured on metrics http server for conversion api shutting down: %s", err.Error())
	}

	logrus.Print("Conversion api service stopped")
}
