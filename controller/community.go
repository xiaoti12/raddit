package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"raddit/service"
	"strconv"
)

// CommunityListHandler return list of models.CommunityBasic struct info
func CommunityListHandler(c *gin.Context) {
	communityList, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("get community info list error", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, communityList)
}

// CommunityDetailHandler return models.CommunityDetail struct data
func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondError(c, CodeInvalidParams)
		return
	}
	detailData, err := service.GetCommunityDetail(int64(id))
	if err != nil {
		zap.L().Error("get community detail info error", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, detailData)
}
