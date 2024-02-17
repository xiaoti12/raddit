package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code ResponseCode `json:"code"`
	Msg  any          `json:"msg"`
	Data any          `json:"data"`
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
