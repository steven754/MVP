package middleware

import (
	"MVP/pkg/e"
	"MVP/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func TokenVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization") //从请求的header中获取toekn字符串
		if token == "" {
			util.ResponseWithJson(e.ERROR_AUTH_TOKEN, "", c)
			c.Abort()
			return
		} else {
			claims, err := util.ParseToken(token) //token校验,claims的内容是自定义的荷载，可以根据里面的id取出用户信息
			if err != nil {                       //token校验失败，返回错误信息
				util.ResponseWithJson(e.ERROR_AUTH_CHECK_TOKEN_FAIL, "", c)
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt { //token过期，返回错误信息
				util.ResponseWithJson(e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, "", c)
				c.Abort()
				return
			} else { //token正确，可以进行后续的操作。设置用户的ID和手机号，供后续方法使用
				c.Set("ID", claims.ID)
				//c.Set("Mobile",claims.Mobile)
				c.Next()
			}
		}
	}
}
