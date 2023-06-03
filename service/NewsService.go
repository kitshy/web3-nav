package service

import (
	"ginweb/common"
	"ginweb/dto"
	"ginweb/model"
)

func GetRecentNotice() (*dto.NoticeDto, error) {
	dto, err := model.GetRecentNotice()
	if err != nil {
		return nil, err
	}
	return dto, err
}

func GetNewsBanner() ([]*dto.NewsDto, error) {
	dtos, err := model.GetNewsBannerList()
	if err != nil {
		return nil, err
	}
	return dtos, err
}

func GetPageNews(page common.Page) (*common.Page, error) {
	dtos, _, err := model.GetPageNews(page)
	if err != nil {
		return nil, err
	}
	page.Data = dtos
	return &page, err
}

func GetNewsDetailById(id int) (*dto.NewsDto, error) {
	dto, err := model.GetNewsDetailById(id)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func AdminGetPageNews(page common.Page) (*common.Page, error) {
	dtos, _, err := model.AdminGetPageNews(page)
	if err != nil {
		return nil, err
	}
	page.Data = dtos
	return &page, err
}

func AddNews(news model.News) error {
	return model.AddNews(news)
}

func UpdateNews(news model.News) error {
	return model.UpdateNews(news)
}
