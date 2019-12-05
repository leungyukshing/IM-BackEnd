package database

import (
	"errors"
	"fmt"
	"github.com/backend/database/entities"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ValidateUser(user *entities.User) error {
	var u []entities.User
	Server.Where("username = ?", user.Username).First(&u)
	if len(u) == 0 {
		return errors.New("User Not Existed")
	}
	temp := u[0]
	fmt.Printf("u: %v, u.Username: %v, u.Password: %v", temp, temp.Username, temp.Password)
	if u[0].Username == "" {
		return errors.New("用户名或密码错误！")
	}
	return nil
}
