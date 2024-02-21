package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"raddit/models"
	"raddit/service"
)

func VotePostHandler(c *gin.Context) {
	p := new(models.VoteParams)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("invalid vote params", zap.Error(err))
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			RespondErrorWithMsg(c, CodeInvalidParams, editValidatorError(err.Error()))
			return
		}
		// not a params validator error
		RespondError(c, CodeInvalidParams)
		return
	}

	err = service.VotePost(p)
	if err != nil {
		zap.L().Error("vote post error", zap.Error(err))
		RespondError(c, CodeServerError)
		return
	}

	RespondSuccess(c, nil)
}
