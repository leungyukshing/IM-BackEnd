package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/astaxie/beego/context"
	"github.com/backend/database"
	"github.com/backend/database/entities"
	im_entities "github.com/backend/im-protobuf/improto"
	"github.com/golang/protobuf/proto"
	"strconv"
)

func (handler *Handler) Login() {
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

func handleLogin(ctx context.Context, loginRequest im_entities.LoginRequest) (im_entities.LoginResponse, error) {
	log := logger(ctx)
	log.Info("handleLogin start")
	loginResponse := im_entities.LoginResponse{}
	email := loginRequest.GetEmail()
	password := loginRequest.GetPassword()
	sum := encodePassword(password)
	ok, user, err := verifyLoginUser(ctx, email, sum)
	if err != nil {
		code := "500"
		message := "DB_ERR"
		loginResponse.Code = &code
		loginResponse.Message = &message
		return loginResponse, err
	}
	if !ok {
		code := "200"
		message := "Email or Password Incorrect"
		loginResponse.Code = &code
		loginResponse.Message = &message
	} else {
		loginResponse.User = &user
		code := "200"
		message := "Login Succeed!"
		loginResponse.Code = &code
		loginResponse.Message = &message
	}
	log.Info("handleLogin finish")
	return loginResponse, nil
}

func encodePassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	sum := hex.EncodeToString(h.Sum(nil))
	return sum
}

func verifyLoginUser(ctx context.Context, email string, password string) (bool, im_entities.User, error) {
	log := logger(ctx)

	// query in DB
	ok, user, err := database.ValidateLogin(email, password)
	if err != nil {
		log.WithError(err).Error("DB_ERR ValidateLogin Failed")
		userPB := im_entities.User{}
		return false, userPB, err
	}
	if !ok {
		userPB := im_entities.User{}
		return false, userPB, nil
	}
	userPB := toUserPB(user)
	return true, userPB, nil
}

func toUserPB(user entities.User) im_entities.User {
	userID := strconv.FormatInt(user.ID, 10)
	userPB := im_entities.User{
		Userid:   &(userID),
		Email:    &(user.Email),
		Username: &(user.Username),
	}
	return userPB
}
