package dbconf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// PathExists
/*
   判断文件或文件夹是否存在
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
   如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Fupload(c *gin.Context){
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":-2,
		})
		return
	}
	sid := c.PostForm("sid")
	cid := c.PostForm("cid")
	name := c.PostForm("name")
	hash := c.PostForm("hash")
	dir1 := "./file/_song_"+sid+"_/"			//歌曲目录
	dir2 := "./file/_song_"+sid+"_/"+cid+"/"	//谱面目录
	flag1,err := PathExists(dir1)
	flag2,err := PathExists(dir2)
	if flag1!=true{
		os.Mkdir(dir1,os.ModePerm)
	}
	if flag2 !=true{
		os.Mkdir(dir2,os.ModePerm)
	}
	//输出上传文件正常名称
	log.Println(file.Filename)
	//将md5名称作为文件保存目录位置
	dst := fmt.Sprintf("%s%s",dir2,name)
	c.SaveUploadedFile(file,dst)
	c.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("'%s' uploaded!", file.Filename),
		"hash":hash,
	})


}