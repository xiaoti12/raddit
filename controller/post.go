package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"raddit/models"
	"raddit/service"
	"strconv"
	"time"
)

// CreatePostHandler create new post
// @Summary create new post
// @Tags Post
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization: Bearer {token}"
// @Param json body models.Post true "create post params"
// @Success 200 {object} ResponseData
// @Router /create/post [post]
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

// GetPostDetailHandler get post detail by id
// @Summary get post detail by id
// @Tags Post
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization: Bearer {token}"
// @Param id path int true "post id"
// @Success 200 {object} SwaggerPostDetailResponse
// @Router /post/{id} [get]
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

// GetPostListHandler get post info list
// @Summary get post info list
// @Tags Post
// @Produce json
// @Param Authorization header string true "Authorization: Bearer {token}"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} SwaggerPostListResponse
// @Router /list/post [get]
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

// GetOrderedPostListHandler get post info list by order(time or score)
// @Summary get post info list by order
// @Tags Post
// @Produce json
// @Param Authorization header string true "Authorization: Bearer {token}"
// @Param page query int false "page"
// @Param size query int false "size"
// @Param order_type query string false "order_type: time/score"
// @Success 200 {object} SwaggerPostListResponse
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
	var posts []*models.PostDetail
	if p.CommunityID == nil {
		posts, err = service.GetOrderedPostList(p)
	} else {
		posts, err = service.GetOrderedPostListByCommunity(p)
	}
	if err != nil {
		zap.L().Error("get post list failed", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, posts)
}
