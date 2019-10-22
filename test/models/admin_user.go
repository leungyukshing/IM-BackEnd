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
	var u *AdminUser
	//fmt.Printf("%v",*u,users.UserName,users.Password)
	db.Where("username = ? AND password = ?", users.UserName, users.Password).Find(u)
	//orm.Where("username=? and password=?", user.UserName, user.Password).Find(&u)
	//fmt.Printf("%v%s%s",*u,*u.UserName,*u.Password)
	if u.UserName == "" {
		return errors.New("用户名或密码错误！")
	}
	return nil
}
