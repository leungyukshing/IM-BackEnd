package dbtest

import (
	"github.com/astaxie/beego"
	"github.com/backend/database"
	"github.com/backend/database/entities"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	user   = beego.AppConfig.String("mysqluserTest")
	pass   = beego.AppConfig.String("mysqlpassTest")
	url    = beego.AppConfig.String("mysqlurlTest")
	port   = beego.AppConfig.String("mysqlportTest")
	dbname = beego.AppConfig.String("mysqldbnameTest")
)

// All testing tables
var tables = []interface{}{
	&entities.User{},
}

func InitTestingMySQL() {
	if database.Server == nil {
		dsn := user + ":" + pass + "@tcp(" + url + ":" + port + ")/" + dbname + "?charset=utf8" + "&parseTime=True&loc=Local"
		var err error
		database.Server, err = gorm.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		database.Server.SingularTable(true)
		createTables()
	}
}

func createTables() {
	if result := database.Server.AutoMigrate(tables...); result.Error != nil {
		panic(result.Error)
	}
}

func ClearTables() {
	for _, table := range tables {
		if !database.Server.HasTable(tables) {
			continue
		}
		if result := database.Server.Where(true).Delete(table); result.Error != nil {
			// retry if table is locked
			i := 0
			d := time.Second
			for ; i < 3; i++ {
				_ = <-time.After(d)
				result := database.Server.Where(true).Delete(table)
				if result.Error == nil {
					break
				}
				d *= 2
			}
			if i >= 3 && result.Error != nil {
				panic(result.Error)
			}
		}
	}
}
