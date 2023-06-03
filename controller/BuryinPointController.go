package controller

import (
	"ginweb/common"
	"ginweb/common/e"
	"ginweb/model"
	"ginweb/service"
	"github.com/gin-gonic/gin"
)

/*
*
埋点 统计
*/
func AddBuryingPoint(c *gin.Context) {

	rs := common.Gin{C: c}
	var pointDto model.BuryingPoint
	c.ShouldBindJSON(&pointDto)

	if pointDto.ClickCode == "" {
		rs.ResponseSuccess("")
		return
	}
	// 如果是 导航，新闻，通知 事件 判断是否有业务id
	if (pointDto.ClickCode == e.NAV || pointDto.ClickCode == e.NEWS ||
		pointDto.ClickCode == e.NOTICE) && pointDto.BusinessId == 0 {
		rs.ResponseSuccess("")
		return
	}
	// 如果是首页搜索事件 需要有搜索关键词
	if pointDto.ClickCode == e.HOME_SEARCH {
		if pointDto.BusinessContent == "" {
			rs.ResponseSuccess("")
			return
		}
		if len(pointDto.BusinessContent) > 512 {
			rs.ResponseSuccess("")
			return
		}
	}
	service.AddBuryingPoint(pointDto)
	rs.ResponseSuccess("success")

}

/*
*
更新
*/
func UpdateBuryingPoint(c *gin.Context) {
	rs := common.Gin{C: c}
	var pointDto model.BuryingPoint
	c.ShouldBindJSON(&pointDto)

	if pointDto.ClickCode == "" {
		rs.ResponseSuccess("")
		return
	}
	// 如果是 导航，新闻，通知 事件 判断是否有业务id
	if (pointDto.ClickCode == e.NAV || pointDto.ClickCode == e.NEWS ||
		pointDto.ClickCode == e.NOTICE) && pointDto.BusinessId == 0 {
		rs.ResponseSuccess("")
		return
	}
	// 如果是首页搜索事件 需要有搜索关键词
	if pointDto.ClickCode == e.HOME_SEARCH {
		if pointDto.BusinessContent == "" {
			rs.ResponseSuccess("")
			return
		}
		if len(pointDto.BusinessContent) > 512 {
			rs.ResponseSuccess("")
			return
		}
	}
	service.UpdateBuryingPoint(pointDto)
	rs.ResponseSuccess("success")
}
