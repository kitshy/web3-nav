package exception

import (
	"ginweb/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandlerException(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case *common.Response:
				log.Printf("panic: %v\n", t.Msg)
				if t.Code == 0 {
					t.Code = http.StatusBadRequest
				}
				if t.Msg == "" {
					t.Msg = "服务器内部异常"
				}
				c.JSON(http.StatusOK, t)
			default:
				log.Printf("panic: internal error")
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "服务器内部异常",
				})
			}
			c.Abort()
		}
	}()
	c.Next()
}
