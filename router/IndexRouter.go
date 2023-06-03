package router

import (
	"ginweb/controller"
	"ginweb/exception"
	"github.com/gin-gonic/gin"
)

func IndexSetupRouter(r *gin.Engine) *gin.Engine {
	r.Use(exception.HandlerException)

	// 首页通知
	r.GET("/getNotice", controller.GetNoticeMsg)

	news := r.Group("/news")
	{
		// 获取首页新闻资讯轮播
		news.GET("/getRotationChart", controller.GetNewsBanner)
		// 分页获取新闻资讯
		news.POST("/getNews", controller.GetPageNews)
		// 获取新闻详情
		news.GET("/getNewsById", controller.GetNewsDetailById)
	}
	category := r.Group("/category")
	{
		// 获取首页所有板块下所有导航  （加缓存）
		category.GET("/getCategoryListAndNav", controller.GetCategoryListAndNav)
	}
	nav := r.Group("/nav")
	{
		// 通过板块id 获取板块下导航列表
		nav.GET("/getNavListByCategoryId", controller.GetNavListByCategoryId)
		// 通过导航id 获取导航详情
		nav.GET("/getNavDetailByNavId", controller.GetNavDetailById)
	}
	return r

}
