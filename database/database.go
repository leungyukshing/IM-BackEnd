package database

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

var (
	user   = beego.AppConfig.String("mysqluser")
	pass   = beego.AppConfig.String("mysqlpass")
	url    = beego.AppConfig.String("mysqlurl")
	port   = beego.AppConfig.String("mysqlport")
	dbname = beego.AppConfig.String("mysqldbname")
)

var Server *gorm.DB

func InitMySQL()  {
	dsn := user + ":" + pass + "@tcp(" + url + ":" + port + ")/" + dbname + "?charset=utf8" + "&parseTime=True&loc=Local"
	Server, err := gorm.Open("mysql", dsn)
	if err == nil {
		fmt.Println("open db sucess")
	} else {
		panic(err)
	}
	Server.SingularTable(true)
}