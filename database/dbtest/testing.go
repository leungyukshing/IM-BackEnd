package dbtest

import (
	"fmt"
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
			dsn = "root:ci@tcp(127.0.0.1:3306)/ci?charset=utf8&parseTime=True&loc=Local"
		} else {
			dsn = "root:******@tcp(127.0.0.1:3306)/gotest?charset=utf8&parseTime=True&loc=Local"
		}
		fmt.Printf("dsn: %v", dsn)
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
