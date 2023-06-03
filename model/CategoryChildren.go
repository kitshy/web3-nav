package model

import "ginweb/dto"

/*
*
子导航
*/
type CategoryChildren struct {
	BaseModel
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Type         string `json:"type"`
	Describe     string `json:"describe"`
	LinkUrl      string `json:"link_url"`
	IconUrl      string `json:"icon_url"`
	Sort         int    `json:"sort"`
}

// 通过父id获取子分类
func GetCategoryByParentIds(ids []int) (dtos []*dto.CategoryChildrenDto, err error) {
	err = db.Model(&CategoryChildren{}).Where("category_id in (?) and is_deleted = 0 order by sort asc,created_time DESC", ids).Scan(&dtos).Error
	if err != nil {
		return nil, err
	}
	return dtos, err
}

func AddChildrenCategory(children CategoryChildren) error {
	return db.Create(&children).Error
}

func UpdateChildrenCategory(children CategoryChildren) error {
	return db.Model(&children).Updates(&children).Error
}
