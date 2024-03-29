package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func GetCtxUserID(c *gin.Context) (int64, error) {
	ctxUID, ok := c.Get(CtxUserID)
	if !ok {
		return 0, ErrorNotLogin
	}
	userID, ok := ctxUID.(int64)
	if !ok {
		return 0, ErrorNotLogin
	}
	return userID, nil
}

func GetPageSize(c *gin.Context) (int, int) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		zap.L().Warn("invalid query param `page`", zap.Error(err), zap.String("page", pageStr))
		page = 1
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		zap.L().Warn("invalid query param `size`", zap.Error(err), zap.String("size", sizeStr))
		size = 10
	}
	return page, size
}
