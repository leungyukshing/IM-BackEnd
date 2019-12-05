package service

import (
	"errors"
	"github.com/astaxie/beego/context"
	im_entities "github.com/backend/im-protobuf/improto"
	"github.com/golang/protobuf/proto"
)

func (this *Handler) Login() {
	loginRequest := im_entities.LoginRequest{}
	proto.Unmarshal(this.Ctx.Input.RequestBody, &loginRequest)
	loginResponse := handleLogin(*this.Ctx, loginRequest)
	data, err := proto.Marshal(&loginResponse)
	if err != nil {
		// log error
	}
	this.Ctx.Output.Body(data)
}

func handleLogin(ctx context.Context, loginRequest im_entities.LoginRequest) im_entities.LoginResponse {
	loginResponse := im_entities.LoginResponse{}
	email := loginRequest.GetEmail()
	password := loginRequest.GetPassword()
	user, err := verifyLoginUser(ctx, email, password)
	loginResponse.User = &user
	if err != nil {
		// log err
		code := "400"
		message := "Login Failed!"
		loginResponse.Code = &code
		loginResponse.Message = &message
	} else {
		code := "200"
		message := "Login Succeed!"
		loginResponse.Code = &code
		loginResponse.Message = &message
	}
	return loginResponse
}

func verifyLoginUser(ctx context.Context, email string, password string) (im_entities.User, error) {
	// query in DB
	email = "test@mail.com"

	return im_entities.User{
		Email: &email,
	}, errors.New("test login")
}
