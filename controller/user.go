package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"raddit/models"
	"raddit/service"
	"reflect"
	"regexp"
	"strings"
)

func RegisterHandler(c *gin.Context) {
	var params models.RegisterParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			c.JSON(http.StatusOK, gin.H{
				"msg": editValidatorError(err.Error()),
			})
			return
		}
		// not a params validator error
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	service.Register()
	c.JSON(http.StatusOK, gin.H{
		"msg": "register success",
	})

}

func CustomValidator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("validator not found")
	}

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return nil
}

func editValidatorError(errStr string) map[string]string {
	re := regexp.MustCompile(`Key: '\w+\.([^']+)'.*Error:(.+)`)
	match := re.FindStringSubmatch(errStr)
	if len(match) != 3 {
		fmt.Println("Invalid validator error format")
		return nil
	}
	result := map[string]string{
		match[1]: match[2],
	}
	// 直接返回map，交由c.JSON处理
	return result
}
