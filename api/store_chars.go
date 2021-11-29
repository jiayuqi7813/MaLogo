package api

import (
	"github.com/gin-gonic/gin"
	"main/dbconf"
)

type Charts struct{
	Sid int     `json:"sid"`
	Cid int     `json:"cid"`
	Uid int     `json:"uid"`
	Creator string `json:"creator"`
	Version string  `json:"version"`
	Level string    `json:"level"`
	Type int        `json:"type"`
	Size int    `json:"size"`
	Mode int    `json:"mode"`
}

var Sid int

//谱面返回
func Scharts(c *gin.Context){
	link()//调用连接
	DB :=dbconf.Dbmalogo
	//自动连接对应
	DB.SingularTable(true)
	var Cs []Charts
	Sid := c.Query("sid")
	if err := DB.Where("sid = ?",Sid).Find(&Cs).Error;err !=nil{
		c.JSON(200,gin.H{
			"code":-2,
		})

	}else {
		c.JSON(200,gin.H{
			"code": 0,
			"hasMore": true,
			"next": 0,
			"data": Cs,
		})
	}
}
