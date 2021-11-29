package api

import (
	"github.com/gin-gonic/gin"
	"main/dbconf"
	"net/http"
	"strconv"
	"strings"
)

type Data struct{
	Sid int 	`json:"sid"`
	Cid int		`json:"cid"`
	Hash string	`json:"hash"`
	Name string	`json:"name"`
}

var Host string = "http://127.0.0.1:9000/api/store/upload"			//host

func Usign(c *gin.Context){
	link()
	DB :=dbconf.Dbmalogo
	//自动连接对应
	DB.SingularTable(true)

	sid := c.PostForm("sid")
	cid := c.PostForm("cid")
	names := c.PostForm("name")
	hashes :=c.PostForm("hash")
	name := strings.Split(names, ",")
	hash := strings.Split(hashes,",")
	var P_data []Data
	for i:=0;i<len(name);i++{
		ssid ,_ := strconv.Atoi(sid) //强制类型strings转int
		ccid,_ := strconv.Atoi(cid)
		filePath := "/file/_song_"+sid+"_/"+cid+"/"+name[i] //存入数据库目录
		items := Items{Cid:ccid,Name: name[i],Hash: hash[i],File: filePath}
		if err := DB.Create(&items).Error;err!=nil{
			c.JSON(200,gin.H{
				"code":-2,
			})
		}
		P_data =append(P_data,Data{				//数据
			Sid:  ssid,
			Cid:  ccid,
			Hash: hash[i],
			Name: name[i],
		})

	}
	c.JSON(http.StatusOK,gin.H{				//回显
		"code": 0,
		"errorIndex": -1,
		"errorMsg": "string",
		"host": Host,
		"meta": P_data,
	})

}