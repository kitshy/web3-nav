package controller

import (
	"ginweb/common"
	"ginweb/dto"
	"ginweb/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
*
获取首页左侧板块列表 子列表嵌套
get home page left category list and child list
*/
func GetCategoryList(c *gin.Context) {

	rs := common.Gin{C: c}
	dtos, err := service.GetGategoryList()
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}
	rs.ResponseSuccess(dtos)

}

/*
*
获取首页所有板块下所有导航
get all nav in under category
*/
func GetCategoryListAndNav(c *gin.Context) {
	rs := common.Gin{C: c}
	dtos, err := service.GetCategoryListAndNav()
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}
	rs.ResponseSuccess(dtos)
}

/*
*
获取父板块列表
*/
func GetParentCategory(c *gin.Context) {
	rs := common.Gin{C: c}
	dtos, err := service.GetParentCategory()
	if err != nil {
		rs.ResponseFail("", nil)
		return
	}
	rs.ResponseSuccess(dtos)
}

/*
*
通过父id获取子板块列表
*/
func GetChildrenCategoryByParentId(c *gin.Context) {
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

	dtos, err := service.GetChildrenCategoryByParentId(id)

	if err != nil {
		rs.ResponseFail("", nil)
		return
	}
	rs.ResponseSuccess(dtos)
}

/*
*
新增编辑板块
*/
func AddEditCategory(c *gin.Context) {
	rs := common.Gin{C: c}
	var dto dto.CategoryChildrenDto
	c.ShouldBindJSON(&dto)

	if dto.Id == 0 {
		//  新增
		if dto.CategoryName == "" {
			rs.ResponseFail("参数错误", nil)
			return
		}
		err := service.AddCategory(dto)
		if err != nil {
			rs.ResponseFail("", nil)
			return
		}
	} else {
		// 编辑
		if dto.CategoryId == 0 {
			rs.ResponseFail("父id必须传", nil)
			return
		}
		err := service.EditCategory(dto)
		if err != nil {
			rs.ResponseFail("", nil)
			return
		}
	}

	rs.ResponseSuccess("")
}
