package routers

import (
	"github.com/astaxie/beego"
	"github.com/backend/service"
)

func init() {
	// test api
	beego.Router("/test", &service.Handler{}, "get:Test")
	beego.Router("/login", &service.Handler{}, "post:Login")
	beego.Router("/register", &service.Handler{}, "post:Register")
	//beego.Router("/getChatList", &service.Handler{}, "get:GetChatList")
	//beego.Router("/getContactsLayout", &service.Handler{},"GetContactsLayout")
}
