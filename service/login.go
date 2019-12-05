package service

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/context"
	im_entities "github.com/backend/im-protobuf/improto"
	"github.com/backend/models"
	"github.com/golang/protobuf/proto"
)

func (c *Handler) Login() {
	loginRequest := im_entities.LoginRequest{}
	proto.Unmarshal(c.Ctx.Input.RequestBody, &loginRequest)
	loginResponse := handleLogin(*c.Ctx, loginRequest)
	data, err := proto.Marshal(&loginResponse)
	if err != nil {
		// log error
	}
	c.Ctx.Output.Body(data)
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

func (c *Handler) LoginCheck() {
	//proto.unmarshal(c.Ctx.Input.RequestBody, req)
	var user database.AdminUser
	inputs := c.Input()
	user.Username = inputs.Get("username")
	user.Password = inputs.Get("password")
	fmt.Printf("input: %v", inputs)
	fmt.Printf("user: %v", user)
	err := user.ValidateUser()
	if err == nil {
		c.SetSession("name", user)
		c.Redirect("/", 301)
	} else {
		fmt.Println(err)
		c.Data["info"] = err
		c.TplName = "error.html"
	}
}

func (c *Handler) Out() {
	c.DelSession("name")
	c.Redirect("/", 301)
}
