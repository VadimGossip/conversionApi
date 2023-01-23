package conversion

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HealthController interface {
	HealthRequest(ctx *gin.Context)
}

type healthController struct {
	appStartedAt time.Time
}

func NewHealthController(appStartedAt time.Time) *healthController {
	return &healthController{appStartedAt: appStartedAt}
}

func (h *healthController) HealthRequest(c *gin.Context) {
	c.JSON(http.StatusOK, HealthCheckResponse{
		Status:       "UP",
		AppStartedAt: h.appStartedAt,
	})
}
