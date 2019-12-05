package routers

import (
	"github.com/astaxie/beego"
	"github.com/backend/service"
)

func init() {
	beego.Router("/", &service.Handler{})
	beego.Router("/login", &service.Handler{}, "post:Login")
	beego.Router("/out", &service.Handler{}, "post:Out")
}
