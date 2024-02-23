package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"raddit/models"
	"raddit/service"
	"strconv"
	"time"
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
	if err != nil {
		zap.L().Error("get user id failed", zap.Error(err))
		RespondError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = uid
	p.CreateTime = time.Now()

	err = service.CreatePost(p)
	if err != nil {
		zap.L().Error("create post failed", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		zap.L().Error("get post detail with invalid params", zap.Error(err), zap.String("id", pidStr))
		RespondError(c, CodeInvalidParams)
		return
	}
	postData, err := service.GetPostDetail(int64(pid))
	if err != nil {
		zap.L().Error("get post detail failed", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, postData)
}

func GetPostListHandler(c *gin.Context) {
	page, size := GetPageSize(c)
	posts, err := service.GetPostList(page, size)
	if err != nil {
		zap.L().Error("get post list failed", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, posts)
}

func GetOrderedPostListHandler(c *gin.Context) {
	// set default value
	p := &models.PostListParams{
		Page:      1,
		Size:      10,
		OrderType: models.OrderByTime,
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("get post list with invalid params", zap.Error(err))
		RespondError(c, CodeInvalidParams)
		return
	}
	posts, err := service.GetOrderedPostList(p)
	if err != nil {
		zap.L().Error("get post list failed", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, posts)
}
