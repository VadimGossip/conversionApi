package conversion

import (
	"github.com/VadimGossip/conversionApi/internal/conversion"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConvController interface {
	ConversionRequest(c *gin.Context)
}

type convController struct {
	convService conversion.Service
}

var _ ConvController = (*convController)(nil)

func NewConvController(convService conversion.Service) *convController {
	return &convController{convService: convService}
}

func (r *convController) ConversionRequest(c *gin.Context) {
	if err := r.convService.PublishMsg(); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
