package model

import (
	"errors"
	"ginweb/common"
	"ginweb/dto"
	"gorm.io/gorm"
)

/*
*
新闻
*/
type News struct {
	BaseModel
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
	LinkUrl string `json:"link_url"`
	IsTop   string `json:"is_top"`
	Click   int    `json:"click"`
	Source  string `json:"source"`
}

func GetNewsDetailById(id int) (dto *dto.NewsDto, err error) {
	err = db.Select(&News{}).Where("is_deleted = 0 and id = ?", id).Scan(&dto).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return dto, err
}

/*
*
get recent one news
*/
func GetNewsBannerList() (dtos []*dto.NewsDto, err error) {
	err = db.Model(&News{}).Where("is_deleted = 0 and is_top = 1 order by created_time").Scan(&dtos).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return dtos, err
}

/*
*
get news by page
*/
func GetPageNews(page common.Page) (dtos []*dto.NewsDto, total int, error error) {
	err = db.Model(&News{}).Where("is_deleted = 0 and is_top = 0 order by created_time").Limit(page.Size).Offset((page.Current - 1) * page.Size).Scan(&dtos).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	return dtos, total, err
}

/*
*
get news by page
*/
func AdminGetPageNews(page common.Page) (dtos []*dto.NewsDto, total int, error error) {
	DB := db.Model(&News{}).Where("is_deleted = 0")

	if page.KeyWord != "" {
		DB = DB.Where("title like '%?%'", page.KeyWord)
	}

	DB = DB.Where("order by is_top, created_time")
	err = DB.Limit(page.Size).Offset((page.Current - 1) * page.Size).Scan(&dtos).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	return dtos, total, err
}

func AddNews(news News) error {
	return db.Create(&news).Error
}

func UpdateNews(news News) error {
	return db.Model(&news).Updates(&news).Error
}
