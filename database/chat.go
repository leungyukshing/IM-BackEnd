package database

import "github.com/backend/database/entities"

func GetChatListByUserID(userID string) ([]entities.Chat, error) {
	var chats []entities.Chat
	db := Server.Where("user_id = ?", userID).Find(&chats)
	if db.Error != nil {
		return chats, db.Error
	}
	return chats, nil
}
