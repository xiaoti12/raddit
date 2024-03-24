package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"raddit/controller"
	_ "raddit/docs"
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
			username := c.MustGet(controller.CtxUserName).(string)
			controller.RespondSuccess(c, gin.H{"username": username})
		})
		v.GET("/list/community", controller.CommunityListHandler)

		v.GET("/post/:id", controller.GetPostDetailHandler)
		v.POST("/vote", controller.VotePostHandler)

		v.Use(middlewares.RateLimitMiddleware())
		{
			v.GET("/community/:id", controller.CommunityDetailHandler)
			v.POST("/create/post", controller.CreatePostHandler)
			v.GET("/list/post", controller.GetPostListHandler)
			v.GET("/list/post/order", controller.GetOrderedPostListHandler)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
