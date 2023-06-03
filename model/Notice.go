package model

import (
	"errors"
	"ginweb/dto"
	"gorm.io/gorm"
)

/*
*
消息通知
*/
type Notice struct {
	BaseModel
	Content string `json:"content"`
	Click   int    `json:"click"`
}

/*
*
get recent one notice
*/
func GetRecentNotice() (dto *dto.NoticeDto, err error) {
	err = db.Model(&Notice{}).Where("is_deleted = 0 order by created_time").Limit(1).Scan(&dto).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return dto, err
}
