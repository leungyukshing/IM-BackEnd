package service

import (
	"github.com/astaxie/beego/context"
	"github.com/backend/database"
	"github.com/backend/database/dbtest"
	im_entities "github.com/backend/im-protobuf/improto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ValidateEmail(t *testing.T) {
	assert.Equal(t, true, validateEmail("test@qq.com"))
	assert.Equal(t, true, validateEmail("test.hh@gmail.com"))
	assert.Equal(t, true, validateEmail("test@163.com"))
	assert.Equal(t, true, validateEmail("testgo@outlook.com"))
	assert.Equal(t, false, validateEmail("test.com"))
	assert.Equal(t, false, validateEmail("test@@gg.com"))
}

func Test_handleRegister(t *testing.T) {
	dbtest.InitTestingMySQL()
	defer dbtest.ClearTables()
	ctx := context.Context{}

	// Correct
	username := "test1"
	password := "password1"
	email := "test1@gmail.com"
	req := im_entities.RegisterRequest{
		Username:             &username,
		Password:             &password,
		Email:                &email,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	resp, err := handleRegister(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, "200", resp.GetCode())
	assert.Equal(t, "Register Success", resp.GetMessage())

	ok, err := database.IsUserExisted(email)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)

	// Existed User
	dbtest.GenerateUsers()
	username = "test"
	password = "password"
	email = "test@gmail.com"
	req.Username = &username
	req.Password = &password
	req.Email = &email
	resp, err = handleRegister(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, "200", resp.GetCode())
	assert.Equal(t, "Email Existed", resp.GetMessage())
}
