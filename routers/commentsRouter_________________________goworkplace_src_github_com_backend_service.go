package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/backend/service:ChatHandler"] = append(beego.GlobalControllerRouter["github.com/backend/service:ChatHandler"],
		beego.ControllerComments{
			Method:           "GetChatList",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/backend/service:LoginHandler"] = append(beego.GlobalControllerRouter["github.com/backend/service:LoginHandler"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/backend/service:RegisterHandler"] = append(beego.GlobalControllerRouter["github.com/backend/service:RegisterHandler"],
		beego.ControllerComments{
			Method:           "Register",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/backend/service:TestHandler"] = append(beego.GlobalControllerRouter["github.com/backend/service:TestHandler"],
		beego.ControllerComments{
			Method:           "Test",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

}
