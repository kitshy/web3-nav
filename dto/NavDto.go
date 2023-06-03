package dto

import "time"

type NavDto struct {
	Id          int       `json:"id"`
	NavName     string    `json:"navName"`
	Type        string    `json:"type"`
	Describe    string    `json:"describe"`
	LinkUrl     string    `json:"linkUrl"`
	ImageUrl    string    `json:"ImageUrl"`
	UpdatedTime time.Time `json:"updatedTime"`
	CreatedTime time.Time `json:"createdTime"`
	Sort        int       `json:"sort"`
	Click       int       `json:"click"`
	CategoryId  int       `json:"categoryId"`
	Size        int       `json:"size"`
	current     int       `json:"current"`
}
