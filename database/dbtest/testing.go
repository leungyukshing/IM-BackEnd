package dbtest

import (
	"github.com/backend/database"
	"github.com/backend/database/entities"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

// All testing tables
var tables = []interface{}{
	&entities.User{},
}

func InitTestingMySQL() {
	var err error
	var dsn string
	if database.Server == nil {
		if env := os.Getenv("ENV"); env == "ci" {
			dsn = "root:ci@tcp(mysql:3306)/ci?charset=utf8&parseTime=True&loc=Local"
		} else {
			dsn = "root:******@tcp(127.0.0.1:3306)/gotest?charset=utf8&parseTime=True&loc=Local"
		}
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
		if !database.Server.HasTable(table) {
			continue
		}
		if result := database.Server.Delete(table); result.Error != nil {
			// retry if table is locked
			i := 0
			d := time.Second
			for ; i < 3; i++ {
				_ = <-time.After(d)
				result := database.Server.Delete(table)
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

func GenerateUsers() {
	user := entities.User{
		Username:    "test",
		Password:    "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", // sha256(password)
		Email:       "test@gmail.com",
		CreatedTime: time.Now(),
		EncryptKey:  "",
	}
	db := database.Server.Create(&user)
	if db.Error != nil {
		panic(db.Error)
	}
}