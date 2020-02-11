package service

import (
	"github.com/astaxie/beego/context"
	"github.com/backend/database/dbtest"
	im_entities "github.com/backend/im-protobuf/improto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_encodePassword(t *testing.T) {
	assert.Equal(t, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", encodePassword("test"))
	assert.Equal(t, "936a185caaa266bb9cbe981e9e05cb78cd732b0b3280eb944412bb6f8f8f07af", encodePassword("helloworld"))
}

func Test_handleLogin(t *testing.T) {
	dbtest.InitTestingMySQL()
	defer dbtest.ClearTables()

	dbtest.GenerateUsers()

	// Correct
	ctx := context.Context{}
	email := "test@gmail.com"
	password := "password"
	req := im_entities.LoginRequest{
		Email:                &email,
		Password:             &password,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	resp, err := handleLogin(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, "200", resp.GetCode())
	assert.Equal(t, "Login Success", resp.GetMessage())
	assert.Equal(t, "test", resp.GetUser().GetUsername())

	// Password Incorrect
	password = "pass"
	req.Password = &password
	resp, err = handleLogin(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, "200", resp.GetCode())
	assert.Equal(t, "Email or Password Incorrect", resp.GetMessage())
	assert.Nil(t, resp.GetUser())

	// Email Incorrect
	email = "aaa@gmail.com"
	req.Email = &email
	resp, err = handleLogin(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, "200", resp.GetCode())
	assert.Equal(t, "Email or Password Incorrect", resp.GetMessage())
	assert.Nil(t, resp.GetUser())
}