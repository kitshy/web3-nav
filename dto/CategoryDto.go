package dto

import (
	"time"
)

/*
*
导航
*/
type CategoryDto struct {
	Childrens    []CategoryChildrenDto `json:"childrens"`
	Id           int                   `json:"id"`
	CategoryName string                `json:"categoryName"`
	Type         string                `json:"type"`
	Describe     string                `json:"describe"`
	LinkUrl      string                `json:"linkUrl"`
	IconUrl      string                `json:"iconUrl"`
	Sort         int                   `json:"sort"`
	UpdatedTime  time.Time             `json:"updatedTime"`
	CreatedTime  time.Time             `json:"createdTime"`
}
