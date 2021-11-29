package api

import (
	"github.com/gin-gonic/gin"
	"main/dbconf"
)

type Items struct{
	Cid int `json:"cid"`
	Name string	`json:"name"`
	Hash string `json:"hash"`
	File string `json:"file"`
}


func Sdownload(c *gin.Context){
	link()
	DB :=dbconf.Dbmalogo
	//自动连接对应
	DB.SingularTable(true)
	var Item []Items
	cid := c.Query("cid")
	ssdi := Sid
	if err := DB.Where("cid = ?",cid).Find(&Item).Error;err !=nil{
		c.JSON(200,gin.H{
			"code":-2,
		})

	}else{
		for i:=0;i<len(Item);i++{
			Item[i].File = "http://"+"127.0.0.1:9000"+Item[i].File
		}
		c.JSON(200,gin.H{
			"code": 0,
			"items": Item,
			"sid": ssdi,
			"cid": cid,
		})
	}
}