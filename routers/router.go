package routers

import (
	"github.com/astaxie/beego"
	"github.com/backend/service"
)

func init() {
	beego.Router("/login", &service.Handler{}, "post:Login")
}
