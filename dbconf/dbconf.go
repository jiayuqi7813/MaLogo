package dbconf

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//定义全局数据库
var (
	Dbmalogo *gorm.DB
)
//定义全局数据库连接
func Inimysql()(err error){
	cfg, err := goconfig.LoadConfigFile("conf/config.ini")
	if err != nil{
		panic("错误")
	}
	group, err := cfg.GetSection("mysql")
	dsn := fmt.Sprintf("%s:%s@%s?charset=utf8mb4&parseTime=True&loc=Local",group["username"],group["password"],group["url"])
	Dbmalogo ,err = gorm.Open("mysql",dsn)
	err = Dbmalogo.DB().Ping()
	return
}