package api

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"main/dbconf"
	"os"
	"path"
	"strconv"
	"strings"
	"github.com/zyx4843/gojson"
	"github.com/jfreymuth/oggvorbis"
)

var Uid int
var Finaldir string
var Length int64
var Creator string
var Version string
var  Mode	string
var Time	string
var  Title string
var Artist string
var  Bpm	string

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
func Ufinish(c *gin.Context){
	link()
	DB :=dbconf.Dbmalogo
	//自动连接对应
	DB.SingularTable(true)
	//参数
	uuid := c.Query("uid")
	Uid,_ = strconv.Atoi(uuid)
	sid := c.PostForm("sid")
	cid := c.PostForm("cid")
	ssid ,_ := strconv.Atoi(sid) //强制类型strings转int
	ccid,_ := strconv.Atoi(cid)
	size := c.PostForm("size")
	names := c.PostForm("name")
	//hashes :=c.PostForm("hash")
	name := strings.Split(names, ",")
	//hash := strings.Split(hashes,",")
	dir2 := "./file/_song_"+sid+"_/"+cid+"/"	//经典位置
	//图片文件后缀名
	arr := [...]string{".jpg",".png",".jpeg"}
	set := make(map[string]struct{})
	for _, value := range arr{
		set[value] = struct{}{}
	}


	for i:=0;i<len(name);i++{
		fixes := path.Ext(name[i])
		fileme := fmt.Sprintf("%s%s",dir2,name[i])
		if _, ok := set[fixes];ok {			//判断是否是图片（此处已知bug，图片无法复制过去）
			filedir := fmt.Sprintf("%s%s",dir2,name[i])		//连接原始路径文件
			Finaldir = fmt.Sprintf("/pic/%s",name[i])		//目标最终文件
			CopyFile(Finaldir,filedir)									//复制
		}
		if fixes==".mc"{				//对铺面文件进行操作

			r, err := zip.OpenReader(fileme)			//本质是压缩包，需要解压里面的内容才可以继续判断
			if err != nil {
				log.Fatal(err)
			}
			defer r.Close()
			for _, f := range r.File {				//虽然是for，但是只运行一次，防止错误
				tmpfile := "tmp.txt"
				rc, _ := f.Open()
				w, err := os.Create(tmpfile)
				if err != nil {
					return
				}
				//内容复制出来
				_, err = io.Copy(w, rc)
				if err != nil {
					return
				}
				w.Close()
				rc.Close()
				//读取复制出的内容
				cc, err := ioutil.ReadFile(tmpfile)
				sss := bytes.TrimPrefix(cc, []byte("\xef\xbb\xbf"))//消除bom
				fuck := string(sss) 								//被sbjson搞崩溃的文明命名
				Creator = gojson.Json(fuck).Get("meta").Get("creator").Tostring()
				Version = gojson.Json(fuck).Get("meta").Get("version").Tostring()
				Mode = gojson.Json(fuck).Get("meta").Get("mode").Tostring()
				Time = gojson.Json(fuck).Get("meta").Get("time").Tostring()
				Title = gojson.Json(fuck).Get("meta").Get("song").Get("title").Tostring()
				Artist = gojson.Json(fuck).Get("meta").Get("song").Get("artist").Tostring()
				Bpm = gojson.Json(fuck).Get("meta").Get("song").Get("bpm").Tostring()
				os.Remove(tmpfile)		//删除缓存文件
			}
		}
		if fixes == ".ogg"{				//对音频文件的长度进行判断
			ogg, _ := os.Open(fileme)
			reader, _:= oggvorbis.NewReader(ogg)
			len :=reader.Length()
			rate := int64(reader.SampleRate())
			Length = len/rate
		}

	}
	imode,_ := strconv.Atoi(Mode)			//类型转换
	itime,_ := strconv.Atoi(Time)			//类型转换
	isize,_ := strconv.Atoi(size)
	songl:= Songlist{						//数据库模型
		Sid:    ssid,
		Cover:  Finaldir,
		Length: int(Length),
		Bpm:    0,
		Title:  Title,
		Artist: Artist,
		Mode:   imode,
		Time: int64(itime),
	}
	charl := Charts{
		Sid:     ssid,
		Cid:     ccid,
		Uid:     Uid,
		Creator: Creator,
		Version: Version,
		Level:   "1",
		Type:    0,
		Size:    isize,
		Mode:    imode,
	}
	if err := DB.Create(&songl).Error;err!=nil { //添加songlist数据
		c.JSON(200, gin.H{
			"code": -1,
		})
	}
	if err := DB.Create(&charl).Error;err!=nil { //添加charts数据
		c.JSON(200, gin.H{
			"code": -2,
		})
	}
	c.JSON(200,gin.H{
		"code": 0,
	})
}


