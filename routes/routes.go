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

	v := r.Group("/api")
	v.POST("/register", controller.RegisterHandler)
	v.POST("/login", controller.LoginHandler)

	v.Use(middlewares.JWTAuthMiddleware())
	{
		v.GET("/home", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			c.JSON(http.StatusOK, gin.H{
				"code": controller.CodeSuccess,
				"msg":  controller.CodeSuccess.Msg(),
				"data": gin.H{"username": username},
			})
		})
		v.GET("/communitylist", controller.CommunityListHandler)
		v.GET("/community/:id", controller.CommunityDetailHandler)

		v.POST("/create/post", controller.CreatePostHandler)
		v.GET("/post/:id", controller.GetPostDetailHandler)
		v.GET("/postlist", controller.GetPostListHandler)
		v.POST("/vote", controller.VotePostHandler)
	}
	return r
}
