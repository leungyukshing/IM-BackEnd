package service

import (
	"github.com/astaxie/beego/context"
	"github.com/backend/database"
	"github.com/backend/database/entities"
	im_entities "github.com/backend/im-protobuf/improto"
	"strconv"
)

func handleGetChatList(ctx context.Context, getChatListRequest im_entities.GetChatListRequest) (im_entities.GetChatListResponse, error) {
	log := logger(ctx)
	log.Info("handleGetChatList start")
	getChatListResponse := im_entities.GetChatListResponse{}
	userID := getChatListRequest.GetUserid()

	if userID == "" {
		code := "200"
		message := "UserID Empty"
		getChatListResponse.Code = &code
		getChatListResponse.Message = &message
		return getChatListResponse, nil
	}

	chatList, err := database.GetChatListByUserID(userID)
	if err != nil {
		log.WithError(err).Error("database.GetChatListByUserID Failed")
		code := "200"
		message := "DB_ERR"
		getChatListResponse.Code = &code
		getChatListResponse.Message = &message
	} else {
		list := toPBChatList(chatList)
		code := "200"
		message := "GetChatList Success"
		getChatListResponse.Code = &code
		getChatListResponse.Message = &message
		getChatListResponse.Chat = list
	}
	return getChatListResponse, nil
}


func toPBChatList(chatList []entities.Chat) []*im_entities.Chat {
	var result []*im_entities.Chat
	for _, chat := range chatList {
		chatID := strconv.FormatInt(chat.ID, 10)
		userID := strconv.FormatInt(chat.UserID, 10)
		lastUpdateTime := chat.LastUpdateTime.String()
		c := im_entities.Chat{
			Chatid:               &(chatID),
			Userid:               &(userID),
			Chatname:             &(chat.Name),
			LastUpdateTime:       &(lastUpdateTime),
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
		m := &c
		result = append(result, m)
	}
	return result
}