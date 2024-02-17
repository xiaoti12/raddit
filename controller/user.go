package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"raddit/dao/mysql"
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
		zap.L().Error("invalid register params", zap.Error(err))
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			RespondErrorWithMsg(c, CodeInvalidParams, editValidatorError(err.Error()))
			return
		}
		// not a params validator error
		RespondError(c, CodeInvalidParams)
		return
	}
	err = service.Register(&params)
	if err != nil {
		zap.L().Error("register error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			RespondError(c, CodeUserExist)
		} else {
			RespondError(c, CodeServerError)
		}
		return
	}
	zap.L().Info("register success", zap.String("username", params.Username))
	RespondSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	var params models.LoginParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		zap.L().Error("invalid login params", zap.Error(err))
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			RespondErrorWithMsg(c, CodeInvalidParams, editValidatorError(err.Error()))
			return
		}
		// not a params validator error
		RespondError(c, CodeInvalidParams)
		return
	}
	var token string
	token, err = service.Login(&params)
	if err != nil {
		zap.L().Error("login error", zap.Error(err))
		switch {
		case errors.Is(err, mysql.ErrorUserNotExist):
			RespondError(c, CodeUserNotExist)
		case errors.Is(err, mysql.ErrorInvalidPassword):
			RespondError(c, CodeInvalidPassword)
		default:
			RespondError(c, CodeServerError)
		}
		return
	}
	RespondSuccess(c, token)
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
