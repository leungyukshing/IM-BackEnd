package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_encodePassword(t *testing.T) {
	assert.Equal(t, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", encodePassword("test"))
	assert.Equal(t, "936a185caaa266bb9cbe981e9e05cb78cd732b0b3280eb944412bb6f8f8f07af", encodePassword("helloworld"))
}