package service

import (
	"github.com/astaxie/beego/context"
	"github.com/backend/database/dbtest"
	im_entities "github.com/backend/im-protobuf/improto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_handleGetChatList(t *testing.T) {
	dbtest.InitTestingMySQL()
	defer dbtest.ClearTables()

	dbtest.GenerateChats()

	// Correct
	ctx := context.Context{}
	userID := "100"
	req := im_entities.GetChatListRequest{
		Userid:               &userID,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	resp, err := handleGetChatList(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, "200", resp.GetCode())
	assert.Equal(t, "GetChatList Success", resp.GetMessage())
	assert.Equal(t, 2, len(resp.GetChat()))
}
