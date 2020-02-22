package service

import (
	"github.com/astaxie/beego/context"
	"github.com/backend/database"
	"github.com/backend/database/entities"
	im_entities "github.com/backend/im-protobuf/improto"
)

func handleGetContactList(ctx context.Context, getContactListRequest im_entities.GetContactListRequest) (im_entities.GetContactListResponse, error) {
	log := logger(ctx)
	log.Info("handleGetContactList start")
	getContactListResponse := im_entities.GetContactListResponse{}
	userID := getContactListRequest.GetUserid()

	if userID == "" {
		code := "200"
		message := "UserID Empty"
		getContactListResponse.Code = &code
		getContactListResponse.Message = &message
		return getContactListResponse, nil
	}

	userIDs, err := database.GetContactIDsByUserID(userID)
	if err != nil {
		log.WithError(err).Error("database.GetContactListByUserID Failed")
		code := "200"
		message := "DB_ERR"
		getContactListResponse.Code = &code
		getContactListResponse.Message = &message
	} else {
		users, err := database.GetUserByIDs(userIDs)
		if err != nil {
			log.WithError(err).Error("database.GetUserByIDs Failed")
			return getContactListResponse, err
		} else {
			code := "200"
			message := "GetContactList Success"
			getContactListResponse.Code = &code
			getContactListResponse.Message = &message
			list := toPBUsers(users)
			getContactListResponse.User = list
		}
	}
	log.Info("handleGetContactList finish")
	return getContactListResponse, nil
}

func toPBUsers(users []entities.User) []*im_entities.User {
	var result []*im_entities.User
	for _, u := range users {
		t := toPBUser(u)
		b := &t
		result = append(result, b)
	}
	return result
}

func handleAddContact(ctx context.Context, addContactRequest im_entities.AddContactRequest) (im_entities.AddContactResponse, error) {
	log := logger(ctx)
	log.Info("handleAddContact start")
	addContactResponse := im_entities.AddContactResponse{}
	// userID := addContactRequest.GetUserid()
	// receiverID := addContactRequest.GetReceiverid()
	log.Info("handleAddContact finish")
	return addContactResponse, nil
}
