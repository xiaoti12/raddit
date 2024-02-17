package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"raddit/controller"
	"raddit/logger"
	"raddit/middlewares"
)

func SetRouteEngine(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
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

	r.POST("/login", controller.LoginHandler)

	r.GET("/home", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{
			"code": controller.CodeSuccess,
			"msg":  controller.CodeSuccess.Msg(),
			"data": gin.H{"username": username},
		})
	})

	return r
}
