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
	&entities.Chat{}, &entities.Contact{}, &entities.User{},
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

func GenerateChats() {
	chat1 := entities.Chat{
		UserID:         int64(100),
		Name:           "testChat1",
		CreatedTIme:    time.Now(),
		LastUpdateTime: time.Now(),
	}
	chat2 := entities.Chat{
		UserID:         int64(100),
		Name:           "testChat2",
		CreatedTIme:    time.Now(),
		LastUpdateTime: time.Now(),
	}
	chat3 := entities.Chat{
		UserID:         int64(101),
		Name:           "testChat3",
		CreatedTIme:    time.Now(),
		LastUpdateTime: time.Now(),
	}
	db := database.Server.Create(&chat1)
	if db.Error != nil {
		panic(db.Error)
	}
	database.Server.Create(&chat2)
	database.Server.Create(&chat3)
}

func GenerateContacts() {
	contact1 := entities.Contact{
		UserID1:     int64(100),
		UserID2:     int64(105),
		CreatedTime: time.Now(),
	}
	contact2 := entities.Contact{
		UserID1:     int64(100),
		UserID2:     int64(102),
		CreatedTime: time.Now(),
	}
	contact3 := entities.Contact{
		UserID1:     int64(104),
		UserID2:     int64(100),
		CreatedTime: time.Now(),
	}
	contact4 := entities.Contact{
		UserID1:     int64(105),
		UserID2:     int64(108),
		CreatedTime: time.Now(),
	}
	user1 := entities.User{
		ID:          int64(102),
		Username:    "user1",
		Password:    "",
		Email:       "",
		CreatedTime: time.Time{},
		EncryptKey:  "",
	}
	user2 := entities.User{
		ID:          int64(104),
		Username:    "user2",
		Password:    "",
		Email:       "",
		CreatedTime: time.Time{},
		EncryptKey:  "",
	}
	user3 := entities.User{
		ID:          int64(105),
		Username:    "",
		Password:    "",
		Email:       "",
		CreatedTime: time.Time{},
		EncryptKey:  "",
	}

	db := database.Server.Create(&contact1)
	if db.Error != nil {
		panic(db.Error)
	}
	database.Server.Create(&contact2)
	database.Server.Create(&contact3)
	database.Server.Create(&contact4)

	database.Server.Create(&user1)
	database.Server.Create(&user2)
	database.Server.Create(&user3)
}