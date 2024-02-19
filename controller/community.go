package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"raddit/service"
)

// CommunityHandler return list of (community_id, community_name)
func CommunityHandler(c *gin.Context) {
	communityList, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("get community info list error", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, communityList)
}
