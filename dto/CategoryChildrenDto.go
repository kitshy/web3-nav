package dto

import "time"

/*
*
子导航
*/
type CategoryChildrenDto struct {
	Id           int       `json:"id"`
	CategoryId   int       `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	Type         string    `json:"type"`
	Describe     string    `json:"describe"`
	LinkUrl      string    `json:"linkUrl"`
	IconUrl      string    `json:"iconUrl"`
	Sort         int       `json:"sort"`
	UpdatedTime  time.Time `json:"updatedTime"`
	CreatedTime  time.Time `json:"CreatedTime"`
}
