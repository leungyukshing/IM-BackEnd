package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ValidateEmail(t *testing.T) {
	assert.Equal(t,true, validateEmail("test@qq.com"))
	assert.Equal(t, true, validateEmail("test.hh@gmail.com"))
	assert.Equal(t, true, validateEmail("test@163.com"))
	assert.Equal(t, true, validateEmail("testgo@outlook.com"))
	assert.Equal(t, false, validateEmail("test.com"))
	assert.Equal(t, false, validateEmail("test@@gg.com"))
}