package model

import (
	"errors"
	"ginweb/common"
	"ginweb/dto"
	"gorm.io/gorm"
)

/*
*
导航
*/
type Nav struct {
	BaseModel
	NavName  string `json:"nav_name"`
	Type     string `json:"type"`
	Describe string `json:"describe"`
	LinkUrl  string `json:"link_url"`
	ImageUrl string `json:"image_url"`
	Sort     int    `json:"sort"`
	Click    int    `json:"click"`
}

/*
*
根据板块id列表获取所有导航
*/
func GetNavListByCategoryIds(ids []int) (navDtos []*dto.NavDto, err error) {
	err = db.Raw("SELECT nav.*,ca.sort,ca.category_id FROM nav nav "+
		"INNER JOIN category_nav ca ON nav.id = ca.nav_id and ca.is_deleted = 0 "+
		"WHERE nav.is_deleted = 0 and ca.category_id in (?) ORDER BY ca.sort ASC", ids).Scan(&navDtos).Error
	if err != nil {
		return nil, err
	}
	return navDtos, err
}

/*
*
根据板块id获取所有导航
*/
func GetNavListByCategoryId(id int) (navs []*Nav, err error) {
	err = db.Raw("SELECT nav.*,ca.sort FROM nav nav "+
		"INNER JOIN category_nav ca ON nav.id = ca.nav_id and ca.is_deleted = 0 "+
		"WHERE nav.is_deleted = 0 and ca.category_id = ? ORDER BY ca.sort ASC", id).Scan(&navs).Error
	if err != nil {
		return nil, err
	}
	return navs, err
}

/*
*
根据id获取详情
*/
func GetNavDetailById(id int) (nav *Nav, err error) {
	err = db.Model(&Nav{}).Where("is_deleted = 0 and id = ?", id).Scan(&nav).Error
	if err != nil {
		return nil, err
	}
	return nav, err
}

func GetNavPage(page common.Page) (navs []*Nav, total int, error error) {
	DB := db.Model(&News{}).Where("is_deleted = 0")
	if page.KeyWord != "" {
		DB = DB.Where("nav_name like %?%", page.KeyWord)
	}
	DB = DB.Where("is_top = 0 order by created_time")
	err := DB.Limit(page.Size).Offset((page.Current - 1) * page.Size).Scan(&navs).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	return navs, total, err
}

func AddNav(nav Nav) (Nav, error) {
	err := db.Create(&nav).Error
	return nav, err
}

func UpdateNav(nav Nav) error {
	return db.Model(&nav).Updates(&nav).Error
}
