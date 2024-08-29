package v1

import (
	"net/http"

	"food-delivery/pkg/logger"
	"github.com/gin-gonic/gin"
)

type fooHandler struct {
	log logger.Logger
}

func NewAuthHandler(
	g *gin.RouterGroup,
	l logger.Logger,
) {
	handler := &fooHandler{
		log: l,
	}
	foo := g.Group("/foo")

	foo.GET("/login", handler.foo)
}

func (f *fooHandler) foo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is the foo endpoint!",
	})
}
