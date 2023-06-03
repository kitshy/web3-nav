package model

/*
*
板块与导航关系
*/
type CategoryNav struct {
	BaseModel
	CategoryId int `json:"category_id"`
	NavId      int `json:"nav_id"`
	Sort       int `json:"sort"`
}

func AddCategoryNav(nav CategoryNav) error {
	return db.Create(&nav).Error
}

func Deleted(categoryId int, navId int) error {
	return db.Delete(&CategoryNav{
		CategoryId: categoryId,
		NavId:      navId,
	}).Error
}
