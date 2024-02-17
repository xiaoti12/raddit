package middlewares

import (
	"github.com/gin-gonic/gin"
	"raddit/controller"
	"raddit/pkg/jwt"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.RespondError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// format: Authorization: Bearer xxx.yyy.zzz
		authParts := strings.SplitN(authHeader, " ", 2)
		if !(len(authParts) == 2 && authParts[0] == "Bearer") {
			controller.RespondError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(authParts[1])
		if err != nil {
			controller.RespondError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
