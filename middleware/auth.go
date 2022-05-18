package middleware

import (
	"github.com/aiio/core"
	"github.com/aiio/core/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth 通用权限判断
func Auth(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("authorization")
		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, core.AuthErr("authorization is empty"))
			ctx.Abort()
			return
		}
		claims, err := token.VerifyToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, core.AuthErr(err.Error()))
			ctx.Abort()
			return
		}
		if !strings.EqualFold(claims.Role, role) {
			ctx.JSON(http.StatusUnauthorized, core.AuthErr("用户权限不正确:"+claims.Role))
			ctx.Abort()
			return
		}
		ctx.Set("uid", claims.UID)
		ctx.Next()
	}
}
