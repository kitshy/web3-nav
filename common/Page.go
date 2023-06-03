package common

type Page struct {
	Size    int         `json:"size"`
	Current int         `json:"current"`
	KeyWord string      `json:"keyWord"`
	Data    interface{} `json:"data"`
}
