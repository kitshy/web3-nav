package router

import (
	"ginweb/controller"
	"ginweb/exception"
	"ginweb/middleware"
	"github.com/gin-gonic/gin"
)

func CommonSetupRouter(r *gin.Engine) *gin.Engine {
	r.Use(exception.HandlerException)
	// 首页埋点
	r.POST("/buryingPoint", controller.AddBuryingPoint)
	// 登录
	r.POST("/login", controller.Login)
	// 注册普通用户
	r.POST("/registerUser", controller.AddUser)

	user := r.Group("user")
	{
		// 鉴权 认证
		user.Use(middleware.AuthToken())
		// 获取用户信息
		user.GET("/getUserInfo", controller.GetUserInfo)
		// 编辑普通用户
		user.POST("/editUser", controller.EditUser)
	}

	return r

}
