package database

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	user   = beego.AppConfig.String("mysqluser")
	pass   = beego.AppConfig.String("mysqlpass")
	url    = beego.AppConfig.String("mysqlurl")
	port   = beego.AppConfig.String("mysqlport")
	dbname = beego.AppConfig.String("mysqldbname")
)

var Server *gorm.DB

func InitMySQL() {
	var err error
	if env := os.Getenv("ENV"); env == "product" {
		dsn := user + ":" + pass + "@tcp(" + url + ":" + port + ")/" + dbname + "?charset=utf8" + "&parseTime=True&loc=Local"
		logs.Info("Product ENV dsn: %v", dsn)
		Server, err = gorm.Open("mysql", dsn)
	} else {
		dsn := "root:******@tcp(127.0.0.1:3306)/gotest?charset=utf8&parseTime=True&loc=Local"
		logs.Info("Dev ENV dsn: %v", dsn)
		Server, err = gorm.Open("mysql", dsn)
	}

	if err == nil {
		fmt.Println("open db success")
	} else {
		panic(err)
	}
	Server.SingularTable(true)
}
