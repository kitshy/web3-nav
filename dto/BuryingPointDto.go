package dto

import (
	"time"
)

/*
*
埋点
*/
type BuryingPointDto struct {
	Id              int       `json:"id"`
	ClickCode       string    `json:"clickCode"`       // 埋点事件
	BusinessId      int       `json:"businessId"`      // 埋点业务id
	BusinessContent string    `json:"businessContent"` // 埋点业务内容
	UpdatedTime     time.Time `json:"updatedTime"`
	CreatedTime     time.Time `json:"createdTime"`
}
