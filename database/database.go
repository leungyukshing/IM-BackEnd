package database

import (
	"github.com/backend/conf"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)
//
//var (
//	user   = beego.AppConfig.String("mysqluser")
//	pass   = beego.AppConfig.String("mysqlpass")
//	url    = beego.AppConfig.String("mysqlurl")
//	port   = beego.AppConfig.String("mysqlport")
//	dbname = beego.AppConfig.String("mysqldbname")
//)

var Server *gorm.DB

func InitMySQL() {
	//if env := os.Getenv("ENV"); env == "product" {
	//	dsn := user + ":" + pass + "@tcp(" + url + ":" + port + ")/" + dbname + "?charset=utf8" + "&parseTime=True&loc=Local"
	//	logs.Info("Product ENV dsn: %v", dsn)
	//	Server, err = gorm.Open("mysql", dsn)
	//} else {
	//	dsn := "root:******@tcp(127.0.0.1:3306)/gotest?charset=utf8&parseTime=True&loc=Local"
	//	logs.Info("Dev ENV dsn: %v", dsn)
	//	Server, err = gorm.Open("mysql", dsn)
	//}
	log := logger()
	log.Info("InitMySQL start")
	dsn := conf.Conf.MysqlDsn
	log.Infof("dsn: %v", dsn)
	Server, err := gorm.Open("mysql", dsn)

	if err == nil {
		log.Info("Open DB Success")
		// fmt.Println("open db success")
	} else {
		panic(err)
	}
	Server.SingularTable(true)
	log.Info("InitMySQL finish")
}

func logger() *logrus.Entry {
	fields := logrus.Fields{}
	return logrus.WithFields(fields)
}