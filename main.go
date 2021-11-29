package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"main/api"
	"main/dbconf"
)

func main(){
	err := dbconf.Inimysql()
	if err != nil{
		panic(err)
	}
	r :=gin.Default()
	r.Static("/file","file")
	r.Static("/pic","pic")
	apiStore := r.Group("/api/store")
	{
		apiStore.GET("/info",api.Sinfo) //服务器信息api
		apiStore.GET("/list",api.Slist)	//歌曲目录api
		apiStore.GET("/charts",api.Scharts)	//根据歌曲后的谱面api
		apiStore.GET("/download",api.Sdownload) //谱面下载api
		apiUpload := apiStore.Group("/upload")	//文件上传路由组
		{
			apiUpload.POST("/sign",api.Usign)				//1.上传文件签名api
			apiUpload.POST("/finish",api.Ufinish)						//3.三阶段验证
		}
		apiStore.POST("/upload",dbconf.Fupload)				//2.文件上传阶段

	}


	r.Run(":9000")
}





