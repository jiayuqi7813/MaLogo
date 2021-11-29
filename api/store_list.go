package api

import (
	"github.com/gin-gonic/gin"
	"main/dbconf"
)

//歌曲列表数据库类型

type Songlist struct{
	Sid int `json:"sid"`
	Cover string `json:"cover"`
	Length int `json:"length"`
	Bpm float64 `json:"bpm"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Mode int `json:"mode"`
	Time int64 `json:"time"`
}

//连接数据库
func link() {
	err := dbconf.Inimysql()
	if err != nil{
		panic(err)
	}
}
//歌曲返回，暂时还没写筛选功能

func Slist(c *gin.Context){
	link()//调用连接
	DB :=dbconf.Dbmalogo
	//自动连接对应
	DB.SingularTable(true)
	DB.AutoMigrate(&Songlist{})
	var Sl []Songlist
	//筛选数据，使其可用
	if err := DB.Find(&Sl).Error;err !=nil{
		c.JSON(200,gin.H{
			"code":-2,
		})
	}else {
		//自动补全cover链接，后期更换算法
		for i:=0;i<len(Sl);i++{
			Sl[i].Cover = "http://"+"127.0.0.1:9000"+Sl[i].Cover
		}
		//返回数据
		c.JSON(200,gin.H{
			"code": 0,
			"hasMore": true,
			"next": 0,
			"data": Sl,
		})
	}
}