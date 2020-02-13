package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/astaxie/beego/context"
	"github.com/backend/database"
	"github.com/backend/database/entities"
	im_entities "github.com/backend/im-protobuf/improto"
	"strconv"
)

func handleLogin(ctx context.Context, loginRequest im_entities.LoginRequest) (im_entities.LoginResponse, error) {
	log := logger(ctx)
	log.Info("handleLogin start")
	loginResponse := im_entities.LoginResponse{}
	email := loginRequest.GetEmail()
	password := loginRequest.GetPassword()

	if email == "" || password == "" {
		code := "200"
		message := "Email or Password Empty"
		loginResponse.Code = &code
		loginResponse.Message = &message
		return loginResponse, nil
	}
	sum := encodePassword(password)
	ok, user, err := verifyLoginUser(ctx, email, sum)
	if err != nil {
		code := "200"
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
		message := "Login Success"
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
	userPB := toPBUser(user)
	return true, userPB, nil
}

func toPBUser(user entities.User) im_entities.User {
	userID := strconv.FormatInt(user.ID, 10)
	userPB := im_entities.User{
		Userid:   &(userID),
		Email:    &(user.Email),
		Username: &(user.Username),
	}
	return userPB
}
