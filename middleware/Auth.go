package middleware

import (
	"encoding/json"
	"ginweb/common"
	"ginweb/common/e"
	"ginweb/model"
	"ginweb/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
普通用户权限
*/
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		rs := common.Gin{C: c}
		accessToken := c.GetHeader("access_token")
		if accessToken == "" {
			rs.Response(http.StatusOK, e.CODE_TOKEN_EXPIRE, "身份认证失败", nil)
			c.Abort()
			return
		}
		// GET USER INFO
		str, err := utils.GetRedis(accessToken)
		if err != nil || str == "" {
			rs.Response(http.StatusOK, e.CODE_TOKEN_EXPIRE, "身份认证失败", nil)
			c.Abort()
			return
		}

		// 存上下文
		var user model.SysUser
		json.Unmarshal([]byte(str), &user)

		c.Request = c.Request.WithContext(utils.NewContext(c.Request.Context(), &user))
		c.Next()
	}
}

/*
*
后台管理权限
*/
func AdminAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		rs := common.Gin{C: c}
		accessToken := c.GetHeader("access_token")
		if accessToken == "" {
			rs.Response(http.StatusOK, e.CODE_TOKEN_EXPIRE, "身份认证失败", nil)
			c.Abort()
			return
		}
		// GET USER INFO
		str, err := utils.GetRedis(accessToken)
		if err != nil || str == "" {
			rs.Response(http.StatusOK, e.CODE_TOKEN_EXPIRE, "身份认证失败", nil)
			c.Abort()
			return
		}
		// 存上下文
		var user model.SysUser
		json.Unmarshal([]byte(str), &user)
		if user.Type != "1" {
			rs.Response(http.StatusOK, e.CODE_TOKEN_EXPIRE, "身份认证失败", nil)
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(utils.NewContext(c.Request.Context(), &user))
		c.Next()
	}
}
