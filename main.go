package main

import (
	"ginweb/config"
	"ginweb/model"
	"ginweb/router"
	"ginweb/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	config.SetUpConfig()
	model.SetUpDb()
	utils.InitRedisPool()
}

func main() {
	r := gin.Default()
	/**
	埋点
	*/
	r = router.CommonSetupRouter(r)
	/**
	首页
	*/
	r = router.IndexSetupRouter(r)
	/**
	登录 用户信息
	*/
	r = router.UserSetupRouter(r)
	r.Run(":8080")
}
