package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	im_entities "github.com/backend/im-protobuf/improto"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

type TestHandler struct {
	beego.Controller
}

type RegisterHandler struct {
	beego.Controller
}

type LoginHandler struct {
	beego.Controller
}

type ContactHandler struct {
	beego.Controller
}

type ChatHandler struct {
	beego.Controller
}


// @Title Test
// @router / [get]
func (handler *TestHandler) Test() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	log.Info("Test start")
	handler.Ctx.Output.Body([]byte("connection success"))
}

// @Title Login
// @router / [post]
func (handler *LoginHandler) Login() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	loginRequest := im_entities.LoginRequest{}
	err := proto.Unmarshal(ctx.Input.RequestBody, &loginRequest)
	if err != nil {
		log.Error("LoginRequest Unmarshal Failed!")
		return
	}
	loginResponse, err := handleLogin(ctx, loginRequest)
	data, err := proto.Marshal(&loginResponse)
	if err != nil {
		// log error
		log.Error("LoginResponse Marshal Failed!")
		return
	}
	handler.Ctx.Output.Body(data)
}

// @Title Register
// @route / [post]
func (handler *RegisterHandler) Register() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	registerRequest := im_entities.RegisterRequest{}
	log.Infof("RequestBody: %v", handler.Ctx.Input.RequestBody)
	err := proto.Unmarshal(ctx.Input.RequestBody, &registerRequest)
	if err != nil {
		log.Error("RegisterRequest Unmarshal Failed!")
		return
	}
	registerResponse, err := handleRegister(ctx, registerRequest)
	data, err := proto.Marshal(&registerResponse)
	if err != nil {
		log.Error("RegisterResponse Marshal Failed!")
		return
	}
	handler.Ctx.Output.Body(data)
}

func logger(ctx context.Context) *logrus.Entry {
	fields := logrus.Fields{}
	return logrus.WithFields(fields)
}