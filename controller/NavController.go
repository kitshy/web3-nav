package controller

import (
	"ginweb/common"
	"ginweb/common/e"
	"ginweb/dto"
	"ginweb/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
*
通过板块id 获取板块下导航列表
get nav list by categoryId
*/
func GetNavListByCategoryId(c *gin.Context) {

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

	navDtos, err := service.GetNavListByCategoryId(id)
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}

	rs.ResponseSuccess(navDtos)
}

/*
*
通过导航id 获取导航详情
get nav details by id
*/
func GetNavDetailById(c *gin.Context) {

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
	navDto, err := service.GetNavDetailById(id)
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}

	// m埋点
	service.BuryingPoint(e.NAV, id, "", "")
	rs.ResponseSuccess(navDto)
}

func GetNavPage(c *gin.Context) {

	rs := common.Gin{C: c}
	var page common.Page
	c.ShouldBindJSON(&page)
	if page.Size == 0 || page.Current == 0 {
		page.Size = 10
		page.Current = 1
	}

	pg, err := service.GetNavPage(page)
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}

	rs.ResponseSuccess(pg)
}

/*
*
新增编辑导航
*/
func AddEditNav(c *gin.Context) {
	rs := common.Gin{C: c}
	var dto dto.NavDto
	c.ShouldBindJSON(&dto)

	if dto.Id == 0 {
		//  新增
		if dto.NavName == "" || dto.LinkUrl == "" || dto.ImageUrl == "" {
			rs.ResponseFail("参数错误", nil)
			return
		}
		err := service.AddNav(dto)
		if err != nil {
			rs.ResponseFail("fail", nil)
			return
		}
	} else {
		// 编辑  -- 只可以编辑导航内容 不可以编辑
		err := service.UpdateNav(dto)
		if err != nil {
			rs.ResponseFail("fail", nil)
			return
		}

	}

	rs.ResponseSuccess("")
}
