package dto

import "time"

type NewsDto struct {
	Id          int       `json:"id"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	LinkUrl     string    `json:"linkUrl"`
	Click       int       `json:"click"`
	Source      string    `json:"source"`
	UpdatedTime time.Time `json:"updatedTime"`
	CreatedTime time.Time `json:"createdTime"`
}
