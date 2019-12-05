package routers

import (
	"github.com/astaxie/beego"
	"github.com/backend/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "post:Login")
	beego.Router("/out", &controllers.MainController{}, "post:Out")
}
