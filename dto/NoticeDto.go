package dto

import "time"

type NoticeDto struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	Click       int       `json:"click"`
	UpdatedTime time.Time `json:"updatedTime"`
	CreatedTime time.Time `json:"createdTime"`
}
