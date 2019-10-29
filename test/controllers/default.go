package controllers

import (
	"github.com/backend/test/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


func (this *MainController) Get() {
	sess := this.GetSession("name")    //判断此次会话的session是否已经存在
	if sess == nil{
		this.Redirect("/login",301)    //跳转到登录逻辑
	} else {
		user := sess.(models.AdminUser)
		this.Data["user"] = user.Username    //用于向前端页面传送数据
		this.Data["pass"] = user.Password
		this.TplName = "succeed.html"    //渲染succeed.html页面
	}
}