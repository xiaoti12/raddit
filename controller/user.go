package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"raddit/models"
	"raddit/service"
)

func RegisterHandler(c *gin.Context) {
	var params models.RegisterParams
	err := c.ShouldBind(&params)
	if err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "invalid params",
		})
	}

	service.Register()
	c.JSON(http.StatusOK, gin.H{
		"msg": "register success",
	})

}
