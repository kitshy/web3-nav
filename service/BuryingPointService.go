package service

import (
	"ginweb/model"
)

/*
*
埋点
*/
func AddBuryingPoint(dto model.BuryingPoint) {
	model.AddBuryingPoint(dto)
}

/*
*
更新埋点
*/
func UpdateBuryingPoint(dto model.BuryingPoint) {
	model.UpdateBuryingPoint(dto)
}

/*
*
埋点
*/
func BuryingPoint(clickCode string, businessId int, businessContent string, userId string) {
	point := model.BuryingPoint{
		BusinessContent: businessContent,
		BusinessId:      businessId,
		ClickCode:       clickCode,
	}
	model.AddBuryingPoint(point)
}
