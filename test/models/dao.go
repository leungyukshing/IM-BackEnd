package models

//import "github.com/astaxie/beego"

/*import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	user = beego.AppConfig.String("mysqluser")
	pass = beego.AppConfig.String("mysqlpass")
	url = beego.AppConfig.String("mysqlurl")
	port = beego.AppConfig.String("mysqlport")
	dbname = beego.AppConfig.String("mysqldbname")
)
func getLink() gorm.Model{
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
	db.Where("username = ? AND password = ?", user, pass).Find("admin_user")
	defer db.Close()
}*/
/*
var db_info = user+":"+pass+"@tcp("+url+":"+port+")/"+dbname+"?charset=utf8"

func getLink() beedb.Model {
	db, err := sql.Open("mysql", db_info)
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}*/