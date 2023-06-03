package common

import (
	"ginweb/common/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) ResponseSuccess(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  "success",
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) ResponseFail(msg string, data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.ERROR,
		Msg:  msg,
		Data: data,
	})
	return
}
