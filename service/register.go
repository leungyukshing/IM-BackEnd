package service

import (
	"github.com/astaxie/beego/context"
	"github.com/backend/database"
	im_entities "github.com/backend/im-protobuf/improto"
	"regexp"
)

func handleRegister(ctx context.Context, registerRequest im_entities.RegisterRequest) (im_entities.RegisterResponse, error) {
	log := logger(ctx)
	log.Info("handleRegister start")

	registerResponse := im_entities.RegisterResponse{}
	email := registerRequest.GetEmail()
	password := registerRequest.GetPassword()
	username := registerRequest.GetUsername()

	if email == "" || password == "" || username == "" {
		log.Info("Empty Field")
		code := "200"
		message := "Empty Field"
		registerResponse.Code = &code
		registerResponse.Message = &message
		return registerResponse, nil
	}
	isExisted, err := database.IsUserExisted(email)

	ok := validateEmail(email)
	if !ok {
		log.Info("Invalid Email")
		code := "200"
		message := "Invalid Email"
		registerResponse.Code = &code
		registerResponse.Message = &message
		return registerResponse, nil
	}
	if err != nil {
		log.WithError(err).Error("DB_ERR IsUserExisted Failed.")
		code := "500"
		message := "DB_ERR"
		registerResponse.Code = &code
		registerResponse.Message = &message
		return registerResponse, nil
	}
	if isExisted {
		log.Info("Email Existed")
		code := "200"
		message := "Email Existed"
		registerResponse.Code = &code
		registerResponse.Message = &message
	} else {
		sum := encodePassword(password)
		err = database.AddUser(username, sum, email)
		if err != nil {
			log.WithError(err).Error("AddUser Failed")
			code := "500"
			message := "DB_ERR"
			registerResponse.Code = &code
			registerResponse.Message = &message
			return registerResponse, err
		} else {
			code := "200"
			message := "Register Success"
			registerResponse.Code = &code
			registerResponse.Message = &message
		}
	}
	log.Info("handleRegister finish")
	return registerResponse, nil
}

func validateEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
