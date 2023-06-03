package service

import (
	"ginweb/dto"
	"ginweb/model"
)

/*
*
获取板块tree
*/
func GetGategoryList() ([]*dto.CategoryDto, error) {

	dtos, err := model.GetAllCategoryList()
	if err != nil {
		return nil, err
	}
	var ids = make([]int, 0)
	for _, dto := range dtos {
		ids = append(ids, dto.Id)
	}
	childrenDtos, err := model.GetCategoryByParentIds(ids)
	if err != nil {
		return nil, err
	}

	var categorys = make([]*dto.CategoryDto, 0)

	for _, v := range dtos {
		var childrens = make([]dto.CategoryChildrenDto, 0)
		for _, child := range childrenDtos {
			if child.CategoryId == v.Id {
				childrens = append(childrens, *child)
			}
		}
		categorys = append(categorys, &dto.CategoryDto{
			Id:           v.Id,
			CategoryName: v.CategoryName,
			Describe:     v.Describe,
			Type:         v.Type,
			LinkUrl:      v.LinkUrl,
			IconUrl:      v.IconUrl,
			Childrens:    childrens,
			Sort:         v.Sort,
			CreatedTime:  v.CreatedTime,
			UpdatedTime:  v.UpdatedTime,
		})
	}

	return categorys, err

}

/*
*
获取板块tree 及下面导航
*/
func GetCategoryListAndNav() ([]*dto.CategoryNavDto, error) {

	//  获取主分类
	dtos, err := model.GetAllCategoryList()
	if err != nil {
		return nil, err
	}
	// 获取子分类
	var ids = make([]int, 0)
	for _, dto := range dtos {
		ids = append(ids, dto.Id)
	}
	childrenDtos, err := model.GetCategoryByParentIds(ids)
	if err != nil {
		return nil, err
	}
	// 获取子分类下导航
	var childrenIds = make([]int, 0)
	for _, dto := range childrenDtos {
		childrenIds = append(childrenIds, dto.Id)
	}
	navDtos, err := model.GetNavListByCategoryIds(childrenIds)
	if err != nil {
		return nil, err
	}

	// 组装数据 遍历主分类
	var categoryNavs = make([]*dto.CategoryNavDto, 0)
	for _, v := range dtos {

		// 遍历子分类
		var childrens = make([]*dto.CategoryNavCategoryList, 0)
		for _, child := range childrenDtos {
			if child.CategoryId == v.Id {

				// 子分类导航 将将导航放进子分类
				var navDtoList = make([]*dto.CategoryNavList, 0)
				for _, navDto := range navDtos {
					if navDto.CategoryId == child.Id {
						navDtoList = append(navDtoList, &dto.CategoryNavList{
							Id:       navDto.Id,
							NavName:  navDto.NavName,
							Sort:     navDto.Sort,
							Type:     navDto.Type,
							LinkUrl:  navDto.LinkUrl,
							ImageUrl: navDto.ImageUrl,
						})
					}
				}

				//  组装子分类列表
				childrens = append(childrens, &dto.CategoryNavCategoryList{
					CategoryName: child.CategoryName,
					Describe:     child.Describe,
					Type:         child.Type,
					LinkUrl:      child.LinkUrl,
					IconUrl:      child.IconUrl,
					Sort:         child.Sort,
					NavList:      navDtoList,
					CreatedTime:  child.CreatedTime,
					UpdatedTime:  child.UpdatedTime,
				})
			}
		}

		// 将子分类放进主分类
		categoryNavs = append(categoryNavs, &dto.CategoryNavDto{
			Id:                   v.Id,
			CategoryName:         v.CategoryName,
			Describe:             v.Describe,
			Type:                 v.Type,
			LinkUrl:              v.LinkUrl,
			IconUrl:              v.IconUrl,
			CategoryChildrenList: childrens,
			Sort:                 v.Sort,
			CreatedTime:          v.CreatedTime,
			UpdatedTime:          v.UpdatedTime,
		})
	}

	return categoryNavs, err
}

func GetParentCategory() ([]*dto.CategoryDto, error) {
	dtos, err := model.GetAllCategoryList()
	if err != nil {
		return nil, err
	}
	var categorys = make([]*dto.CategoryDto, 0)

	for _, v := range dtos {
		categorys = append(categorys, &dto.CategoryDto{
			Id:           v.Id,
			CategoryName: v.CategoryName,
			Describe:     v.Describe,
			Type:         v.Type,
			LinkUrl:      v.LinkUrl,
			IconUrl:      v.IconUrl,
			Sort:         v.Sort,
			CreatedTime:  v.CreatedTime,
			UpdatedTime:  v.UpdatedTime,
		})
	}
	return categorys, nil
}

func GetChildrenCategoryByParentId(id int) ([]*dto.CategoryChildrenDto, error) {
	ids := []int{id}
	childrenDtos, err := model.GetCategoryByParentIds(ids)
	if err != nil {
		return nil, err
	}
	return childrenDtos, err
}

func AddCategory(dto dto.CategoryChildrenDto) error {
	if dto.CategoryId == 0 {
		// 父类
		mo := model.Category{
			CategoryName: dto.CategoryName,
			Describe:     dto.Describe,
			Type:         dto.Type,
			LinkUrl:      dto.LinkUrl,
			IconUrl:      dto.IconUrl,
			Sort:         dto.Sort,
		}
		if mo.Sort == 0 {
			mo.Sort = 999
		}
		return model.AddCategory(mo)
	} else {
		// zilei
		// 父类
		mo := model.CategoryChildren{
			CategoryName: dto.CategoryName,
			Describe:     dto.Describe,
			Type:         dto.Type,
			LinkUrl:      dto.LinkUrl,
			IconUrl:      dto.IconUrl,
			Sort:         dto.Sort,
			CategoryId:   dto.CategoryId,
		}
		if mo.Sort == 0 {
			mo.Sort = 999
		}
		return model.AddChildrenCategory(mo)
	}
}

func EditCategory(dto dto.CategoryChildrenDto) error {
	if dto.CategoryId == 0 {
		// 父类
		mo := model.Category{
			CategoryName: dto.CategoryName,
			Describe:     dto.Describe,
			Type:         dto.Type,
			LinkUrl:      dto.LinkUrl,
			IconUrl:      dto.IconUrl,
			Sort:         dto.Sort,
		}
		if mo.Sort == 0 {
			mo.Sort = 999
		}
		mo.Id = dto.Id
		return model.UpdateCategory(mo)
	} else {
		// zilei
		// 父类
		mo := model.CategoryChildren{
			CategoryName: dto.CategoryName,
			Describe:     dto.Describe,
			Type:         dto.Type,
			LinkUrl:      dto.LinkUrl,
			IconUrl:      dto.IconUrl,
			Sort:         dto.Sort,
			CategoryId:   dto.CategoryId,
		}
		mo.Id = dto.Id
		if mo.Sort == 0 {
			mo.Sort = 999
		}
		return model.UpdateChildrenCategory(mo)
	}
}
