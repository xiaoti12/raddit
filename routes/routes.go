package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouteEngine() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello I'm gin")
	})
	return r
}
