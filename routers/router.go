package routers

import (
	"github.com/astaxie/beego"
	"github.com/backend/service"
)

func init() {
	// test api
	//beego.Router("/test", &service.Handler{}, "get:Test")
	//beego.Router("/login", &service.Handler{}, "post:Login")
	//beego.Router("/register", &service.Handler{}, "post:Register")
	//beego.Router("/getChatList", &service.Handler{}, "get:GetChatList")
	//beego.Router("/getContactsList", &service.Handler{},"GetContactsLayout")
	ns := beego.NewNamespace("/apiv1",
		beego.NSNamespace("/test",
			beego.NSInclude(
				&service.TestHandler{}),
		),
		beego.NSNamespace("/register",
			beego.NSInclude(
				&service.RegisterHandler{}),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&service.LoginHandler{}),
		),
		beego.NSNamespace("getChatList",
			beego.NSInclude(
				&service.ChatHandler{}),
		),
		beego.NSNamespace("getContactList",
			beego.NSInclude(
				&service.ContactHandler{}),
		),
	)
	beego.AddNamespace(ns)
}
