package middleware

import (
	"github.com/gin-gonic/gin"
	"go_study/src/utils"
	"net/http"
	"strings"
)

func JwtCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		//从请求头中获取token
		token := context.Request.Header.Get("Authorization")
		if token == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "请求未携带token，无权限访问",
			})
			context.Abort()
			return
		}
		//以空格分割token
		tokenArr := strings.SplitN(token, " ", 2)
		if len(tokenArr) != 2 || tokenArr[0] != "Bearer" {
			context.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "请求头中Authorization格式有误",
			})
			context.Abort()
			return
		}
		//对token进行验证
		parseToken, err := utils.ParseToken(tokenArr[1])
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  err.Error(),
			})
			context.Abort()
			return
		}
		//将解析出来的信息保存到上下文中
		context.Set("userId", parseToken.UserId)
		context.Set("userName", parseToken.UserName)
		context.Next()
	}
}
