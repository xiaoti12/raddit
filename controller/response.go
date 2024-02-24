package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"raddit/models"
)

type ResponseData struct {
	Code ResponseCode `json:"code"`
	Msg  any          `json:"msg"`  // may be string or json object
	Data any          `json:"data"` // may be string or json object
}

type SwaggerCommunityListResponse struct {
	Code ResponseCode             `json:"code"`
	Msg  string                   `json:"msg"`
	Data []*models.CommunityBasic `json:"data"`
}

type SwaggerCommunityDetailResponse struct {
	Code ResponseCode            `json:"code"`
	Msg  string                  `json:"msg"`
	Data *models.CommunityDetail `json:"data"`
}

type SwaggerPostListResponse struct {
	Code ResponseCode         `json:"code"`
	Msg  string               `json:"msg"`
	Data []*models.PostDetail `json:"data"`
}

type SwaggerPostDetailResponse struct {
	Code ResponseCode       `json:"code"`
	Msg  string             `json:"msg"`
	Data *models.PostDetail `json:"data"`
}

func RespondSuccess(c *gin.Context, data any) {
	respData := ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, respData)

}

func RespondError(c *gin.Context, code ResponseCode) {
	respData := ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, respData)
}
func RespondErrorWithMsg(c *gin.Context, code ResponseCode, msg any) {
	respData := ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, respData)
}
