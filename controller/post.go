package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"raddit/models"
	"raddit/service"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("create post with invalid params", zap.Error(err))
		RespondError(c, CodeInvalidParams)
		return
	}
	// TODO: add getting userID from context
	err = service.CreatePost(p)
	if err != nil {
		zap.L().Error("create post failed", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, nil)
}
