package conversion

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConvController interface {
	ConversionRequest(c *gin.Context)
}

type convController struct {
}

var _ ConvController = (*convController)(nil)

func NewConvController() *convController {
	return &convController{}
}

func (r *convController) ConversionRequest(c *gin.Context) {
	c.Status(http.StatusOK)
}
