package service

import (
	"github.com/astaxie/beego"
	"github.com/backend/database/entities"
)

type Handler struct {
	beego.Controller
}

func (this *Handler) Get() {
	sess := this.GetSession("name") //判断此次会话的session是否已经存在
	if sess == nil {
		this.Redirect("/login", 301) //跳转到登录逻辑
	} else {
		user := sess.(entities.User)
		this.Data["user"] = user.Username //用于向前端页面传送数据
		this.Data["pass"] = user.Password
		this.TplName = "succeed.html" //渲染succeed.html页面
	}
}
