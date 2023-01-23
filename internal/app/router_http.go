package app

import (
	"github.com/VadimGossip/conversionApi/internal/api/server/conversion"
	"github.com/gin-gonic/gin"
)

func initMetricsHttpRouter(app *App) {
	s := app.metricsHttpServer
	s.Use(gin.Recovery())
	healthController := conversion.NewHealthController(app.appStartedAt)
	s.GET("/healthcheck", healthController.HealthRequest)
}

func initHttpRouter(app *App) {
	s := app.apiHttpServer
	s.Use(gin.Recovery())
	s.Use(gin.Logger())

	convController := conversion.NewConvController()
	s.GET("/api/sms/received", convController.ConversionRequest)
	s.POST("/api/sms/received", convController.ConversionRequest)
}
