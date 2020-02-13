package database

import (
	"github.com/backend/database/entities"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func ValidateLogin(email string, password string) (bool, entities.User, error) {
	var u entities.User
	db := Server.Where("email = ?", email).First(&u)
	if db.RecordNotFound() {
		return false, u, nil
	}
	if db.Error != nil {
		return false, u, db.Error
	}
	if u.Password != password {
		return false, u, nil
	}
	return true, u, nil
}

func IsUserExisted(email string) (bool, error) {
	log := logger()
	log.Infof("email: %v", email)
	log.Infof("Server: %v", Server)
	var u entities.User
	db := Server.Where("email = ?", email).First(&u)
	if db.RecordNotFound() {
		return false, nil
	}
	if db.Error != nil {
		return false, db.Error
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

func GetUserByIDs(userIDs []int64) ([]entities.User, error) {
	var u []entities.User
	db := Server.Where("id in (?)", userIDs).Find(&u)
	if db.Error != nil {
		return u, db.Error
	}
	return u, nil
}
