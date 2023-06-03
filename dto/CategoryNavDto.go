package dto

import "time"

type CategoryNavDto struct {
	CategoryChildrenList []*CategoryNavCategoryList `json:"categoryChildrenList"`
	Id                   int                        `json:"id"`
	CategoryName         string                     `json:"categoryName"`
	Type                 string                     `json:"type"`
	Describe             string                     `json:"describe"`
	LinkUrl              string                     `json:"linkUrl"`
	IconUrl              string                     `json:"iconUrl"`
	Sort                 int                        `json:"sort"`
	UpdatedTime          time.Time                  `json:"updatedTime"`
	CreatedTime          time.Time                  `json:"createdTime"`
}

type CategoryNavCategoryList struct {
	NavList      []*CategoryNavList `json:"navList"`
	Id           int                `json:"id"`
	CategoryName string             `json:"categoryName"`
	Type         string             `json:"type"`
	Describe     string             `json:"describe"`
	LinkUrl      string             `json:"linkUrl"`
	IconUrl      string             `json:"iconUrl"`
	Sort         int                `json:"sort"`
	UpdatedTime  time.Time          `json:"updatedTime"`
	CreatedTime  time.Time          `json:"createdTime"`
}

type CategoryNavList struct {
	Id          int       `json:"id"`
	NavName     string    `json:"navName"`
	Type        string    `json:"type"`
	Describe    string    `json:"describe"`
	LinkUrl     string    `json:"linkUrl"`
	ImageUrl    string    `json:"ImageUrl"`
	UpdatedTime time.Time `json:"updatedTime"`
	CreatedTime time.Time `json:"createdTime"`
	Sort        int       `json:"sort"`
	CategoryId  int       `json:"categoryId"`
	Size        int       `json:"size"`
	current     int       `json:"current"`
}
