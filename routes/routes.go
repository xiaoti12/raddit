package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"raddit/controller"
	"raddit/logger"
)

func SetRouteEngine() *gin.Engine {
	err := controller.CustomValidator()
	if err != nil {
		fmt.Println(err)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello I'm gin\n")
	})

	r.POST("/register", controller.RegisterHandler)

	return r
}
