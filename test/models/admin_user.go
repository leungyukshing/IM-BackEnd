package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	user = beego.AppConfig.String("mysqluser")
	pass = beego.AppConfig.String("mysqlpass")
	url = beego.AppConfig.String("mysqlurl")
	port = beego.AppConfig.String("mysqlport")
	dbname = beego.AppConfig.String("mysqldbname")
)

type AdminUser struct {
	Id int
	UserName string
	Password string
	Email string
	Author int
}

func (users AdminUser)ValidateUser() error {
	var db_info = user+":"+pass+"@tcp("+url+":"+port+")/"+dbname+"?charset=utf8"+"&parseTime=True&loc=Local"
	//orm := getLink(user.UserName, user.Password)    //获得用于操作数据库的orm
	db, err := gorm.Open("mysql", db_info)
	if err == nil {
		fmt.Println("open db sucess")
	} else {
		fmt.Println("open db error ", err)
	}
	db.SingularTable(true)
	var u []AdminUser
	db.Where("user_name = ?", users.UserName).First(&u)

	temp := u[0]
	fmt.Printf("u: %v, u.Username: %v, u.Password: %v", temp, temp.UserName, temp.Password)
	if u[0].UserName == "" {
		return errors.New("用户名或密码错误！")
	}
	return nil
}
