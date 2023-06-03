package controller

import (
	"ginweb/common"
	"ginweb/common/e"
	"ginweb/model"
	"ginweb/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
*
首页通知
home page notice
*/
func GetNoticeMsg(c *gin.Context) {
	rs := common.Gin{C: c}
	noticeDto, err := service.GetRecentNotice()
	if err != nil {
		rs.ResponseFail("", noticeDto)
		return
	}
	rs.ResponseSuccess(noticeDto)
}

/*
*
获取新闻资讯banner图轮播
get news banner rotation chart
*/
func GetNewsBanner(c *gin.Context) {
	rs := common.Gin{C: c}
	news, err := service.GetNewsBanner()
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}
	rs.ResponseSuccess(news)
}

/*
*
新闻资讯 根据分页获取
get news by page
*/
func GetPageNews(c *gin.Context) {
	rs := common.Gin{C: c}
	var page common.Page
	c.ShouldBindJSON(&page)
	if page.Size == 0 || page.Current == 0 {
		page.Size = 10
		page.Current = 1
	}
	result, err := service.GetPageNews(page)
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}
	rs.ResponseSuccess(result)
}

/*
*
获取新闻详情
*/
func GetNewsDetailById(c *gin.Context) {
	rs := common.Gin{C: c}
	idStr, ok := c.GetQuery("id")
	if !ok {
		rs.ResponseFail("参数不能为空", nil)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rs.ResponseFail("参数错误", nil)
		return
	}
	dto, err := service.GetNewsDetailById(id)
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}

	// m埋点
	service.BuryingPoint(e.NEWS, id, "", "")
	rs.ResponseSuccess(dto)
}

/*
*
新闻资讯 根据分页获取
get news by page
*/
func AdminGetPageNews(c *gin.Context) {
	rs := common.Gin{C: c}
	var page common.Page
	c.ShouldBindJSON(&page)
	if page.Size == 0 || page.Current == 0 {
		page.Size = 10
		page.Current = 1
	}
	result, err := service.AdminGetPageNews(page)
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}
	rs.ResponseSuccess(result)
}

/*
*
新增编辑新闻
*/
func AddEditNews(c *gin.Context) {
	rs := common.Gin{C: c}
	var model model.News
	c.ShouldBindJSON(&model)

	if model.Id == 0 {
		if model.Type == "" || model.Source == "" || model.Title == "" {
			rs.ResponseFail("参数错误", nil)
			return
		}
		if model.IsTop == "" {
			model.IsTop = "0"
		}
		model.Click = 0
		//  新增
		err := service.AddNews(model)
		if err != nil {
			rs.ResponseFail("fail", nil)
			return
		}
	} else {
		// 编辑
		model.Click = 0
		err := service.UpdateNews(model)
		if err != nil {
			rs.ResponseFail("fail", nil)
			return
		}
	}

	rs.ResponseSuccess("")
}
