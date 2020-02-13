package service

import (
	"github.com/astaxie/beego/context"
	"github.com/backend/database/dbtest"
	im_entities "github.com/backend/im-protobuf/improto"
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test_handleGetContactList(t *testing.T) {
	dbtest.InitTestingMySQL()
	defer dbtest.ClearTables()

	dbtest.GenerateContacts()

	ctx := context.Context{}
	userID := "100"
	req := im_entities.GetContactListRequest{
		Userid:               &userID,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	resp, err := handleGetContactList(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, "200", resp.GetCode())
	assert.Equal(t, "GetContactList Success", resp.GetMessage())
	assert.Equal(t, 3, len(resp.GetUser()))
}