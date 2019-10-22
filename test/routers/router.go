package routers

import (
	"github.com/astaxie/beego"
	"Test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{},"get:Login;post:LoginCheck")
	beego.Router("/out", &controllers.MainController{},"post:Out")

}