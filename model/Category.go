package model

/*
*
导航
*/
type Category struct {
	BaseModel
	CategoryName string `json:"category_name"`
	Type         string `json:"type"`
	Describe     string `json:"describe"`
	LinkUrl      string `json:"link_url"`
	IconUrl      string `json:"icon_url"`
	Sort         int    `json:"sort"`
}

// 获取所有分类列表
func GetAllCategoryList() (dtos []*Category, err error) {
	err = db.Model(&Category{}).Where("is_deleted = 0 order by sort asc").Scan(&dtos).Error
	if err != nil {
		return nil, err
	}
	return dtos, err
}

func AddCategory(category Category) error {
	return db.Create(&category).Error
}

func UpdateCategory(category Category) error {
	return db.Model(&category).Updates(&category).Error
}
