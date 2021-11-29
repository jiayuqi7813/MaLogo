package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//服务器信息，welcome可任意修改
func Sinfo(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"api": 202108,
		"min": 202103,
		"welcome": "Welcome maLOGO",
	})
}