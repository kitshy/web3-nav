package router

import (
	"ginweb/controller"
	"ginweb/exception"
	"ginweb/middleware"
	"github.com/gin-gonic/gin"
)

func UserSetupRouter(r *gin.Engine) *gin.Engine {
	r.Use(exception.HandlerException)

	admin := r.Group("/admin")
	admin.Use(middleware.AdminAuthToken())
	{
		// 新增编辑管理员
		admin.POST("/addEditUser", controller.AddEditUser)
		// 文件上传
		admin.POST("/uploadFile")

		// 新闻，信息，轮播图
		news := admin.Group("/news")
		{
			// 分页获取信息
			news.POST("/getNewsPage", controller.AdminGetPageNews)
			// 新增编辑信息
			news.POST("/addEditNews", controller.AddEditNews)
		}
		// 分类
		category := admin.Group("/category")
		{
			// 获取父分类
			category.GET("/getParentCategory", controller.GetParentCategory)
			// 通过父分类获取子分类
			category.GET("/getChildrenCategoryByParentId", controller.GetChildrenCategoryByParentId)
			// 获取首页左侧板块列表 子列表嵌套
			category.GET("/getCategoryList", controller.GetCategoryList)
			// 新增编辑分类
			category.POST("/addEditParentCategory", controller.AddEditCategory)
		}
		// 导航
		nav := admin.Group("/nav")
		{
			//  分页获取导航
			nav.POST("/getNavPage", controller.GetNavPage)
			// 新增编辑导航
			nav.POST("/addEditNav", controller.AddEditNav)
		}
	}

	return r

}
