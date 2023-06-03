package service

import (
	"ginweb/common"
	"ginweb/dto"
	"ginweb/model"
)

/*
*
通过板块Id获取导航详情
*/
func GetNavListByCategoryId(id int) ([]*dto.NavDto, error) {

	dtos, err := model.GetNavListByCategoryId(id)
	if err != nil {
		return nil, err
	}

	var navDtos = make([]*dto.NavDto, 0)

	for _, nav := range dtos {
		navDto := &dto.NavDto{
			Id:          nav.Id,
			NavName:     nav.NavName,
			Type:        nav.Type,
			Describe:    nav.Describe,
			LinkUrl:     nav.LinkUrl,
			ImageUrl:    nav.ImageUrl,
			UpdatedTime: nav.UpdatedTime,
			CreatedTime: nav.CreatedTime,
			Sort:        nav.Sort,
		}
		navDtos = append(navDtos, navDto)
	}

	return navDtos, nil

}

/*
*
通过导航id获取导航详情
*/
func GetNavDetailById(id int) (*dto.NavDto, error) {
	nav, err := model.GetNavDetailById(id)
	if err != nil || nav == nil {
		return nil, err
	}
	navDto := &dto.NavDto{
		Id:          nav.Id,
		NavName:     nav.NavName,
		Type:        nav.Type,
		Describe:    nav.Describe,
		LinkUrl:     nav.LinkUrl,
		ImageUrl:    nav.ImageUrl,
		UpdatedTime: nav.UpdatedTime,
		CreatedTime: nav.CreatedTime,
	}
	return navDto, nil

}

func GetNavPage(page common.Page) (*common.Page, error) {
	navs, _, err := model.GetNavPage(page)
	if err != nil {
		return nil, err
	}
	page.Data = navs
	return &page, err
}

func AddNav(dto dto.NavDto) error {
	mo := model.Nav{
		NavName:  dto.NavName,
		ImageUrl: dto.ImageUrl,
		LinkUrl:  dto.LinkUrl,
		Describe: dto.Describe,
		Type:     dto.Type,
	}

	m, err := model.AddNav(mo)
	if err != nil {
		return err
	}
	var cateNav model.CategoryNav
	if dto.CategoryId == 0 {
		cateNav = model.CategoryNav{
			CategoryId: dto.CategoryId,
			NavId:      m.Id,
			Sort:       dto.Sort,
		}
		if dto.Sort == 0 {
			cateNav.Sort = 999
		}
		err := model.AddCategoryNav(cateNav)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateNav(dto dto.NavDto) error {
	mo := model.Nav{
		NavName:  dto.NavName,
		ImageUrl: dto.ImageUrl,
		LinkUrl:  dto.LinkUrl,
		Describe: dto.Describe,
		Type:     dto.Type,
	}
	return model.UpdateNav(mo)
}
