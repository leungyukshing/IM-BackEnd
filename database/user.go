package database

import (
	"github.com/backend/database/entities"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func ValidateLogin(email string, password string) (bool, entities.User, error) {
	var u entities.User
	db := Server.Where("email = ?", email).First(&u)
	if db.Error != nil {
		return false, u, db.Error
	}
	if u.Password != password {
		return false, u, nil
	}
	return true, u,nil
}

func IsUserExisted(email string) (bool, error) {
	var u []entities.User
	db := Server.Where("email = ?", email).First(&u)
	if db.Error != nil {
		return false, db.Error
	}
	if len(u) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func AddUser(username string, password string, email string) error {
	user := entities.User{
		Username:    username,
		Password:    password,
		Email:       email,
		CreatedTime: time.Now(),
		EncryptKey:  "",
	}
	db := Server.Create(&user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}