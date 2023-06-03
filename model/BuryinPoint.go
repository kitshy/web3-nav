package model

/*
*
埋点
*/
type BuryingPoint struct {
	BaseModel
	ClickCode       string `json:"clickCode"`       // 埋点事件
	BusinessId      int    `json:"businessId"`      // 埋点业务id
	BusinessContent string `json:"businessContent"` // 埋点业务内容
}

/*
*
添加数据
*/
func AddBuryingPoint(dto BuryingPoint) {
	db.Create(&dto)
}

/*
*
更新埋点
*/
func UpdateBuryingPoint(dto BuryingPoint) {
	db.Model(&dto).Updates(&dto)
}
