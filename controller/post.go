package controller

import (
	"fmt"
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

	uid, err := GetCtxUserID(c)
	fmt.Println("uid: ", uid)
	if err != nil {
		zap.L().Error("get user id failed", zap.Error(err))
		RespondError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = uid

	err = service.CreatePost(p)
	if err != nil {
		zap.L().Error("create post failed", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, nil)
}
