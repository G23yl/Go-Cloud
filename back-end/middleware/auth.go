// 设置身份验证中间件
package middleware

import (
	"disk/model/response"
	"disk/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从头部Authorization字段获取token
		tokenString := ctx.Request.Header.Get("Authorization")
		// 解析token
		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			response.UnauthorizedError(ctx, response.ErrorUnauthorized)
			ctx.Abort()
			return
		}
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
