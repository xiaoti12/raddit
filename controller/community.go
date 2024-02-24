package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"raddit/service"
	"strconv"
)

// CommunityListHandler get list of community basic info
// @Summary get community basic info list
// @Tags Community
// @Produce json
// @Param Authorization header string true "Authorization: Bearer {token}"
// @Success 200 {object} SwaggerCommunityListResponse
// @Router /list/community [get]
func CommunityListHandler(c *gin.Context) {
	communityList, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("get community info list error", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}
	RespondSuccess(c, communityList)
}

// CommunityDetailHandler get community detail info by id
// @Summary get community detail info by id
// @Tags Community
// @Produce json
// @Param id path int true "community id"
// @Param Authorization header string true "Authorization: Bearer {token}"
// @Success 200 {object} SwaggerCommunityDetailResponse
// @Router /community/{id} [get]
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
