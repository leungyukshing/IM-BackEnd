package routers

import (
	"github.com/astaxie/beego"
	"github.com/backend/service"
)

func init() {
	// test api
	beego.Router("/test", &service.Handler{}, "get:Test")
	beego.Router("/login", &service.Handler{}, "post:Login")
}
