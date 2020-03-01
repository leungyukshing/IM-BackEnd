package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	im_entities "github.com/backend/im-protobuf/improto"
	"github.com/backend/push"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

type TestHandler struct {
	beego.Controller
}

type RegisterHandler struct {
	beego.Controller
}

type LoginHandler struct {
	beego.Controller
}

type ContactHandler struct {
	beego.Controller
}

type ChatHandler struct {
	beego.Controller
}

type WebSocketHandler struct {
	beego.Controller
}

// @Title Test
// @router / [get]
func (handler *TestHandler) Test() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	log.Info("Test start")
	handler.Ctx.Output.Body([]byte("connection success"))
}

// @Title SetUpWebsocket
// @router / [get]
func (handler *WebSocketHandler) SetUpWebsocket()  {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	log.Info("SetUpWebsocket start")
	push.SetUpWebsocket(ctx)
	push.PushMessage(ctx)
	handler.Ctx.Output.Body([]byte("set up websocket success"))
}

// @Title Login
// @router / [post]
func (handler *LoginHandler) Login() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	loginRequest := im_entities.LoginRequest{}
	err := proto.Unmarshal(ctx.Input.RequestBody, &loginRequest)
	if err != nil {
		log.WithError(err).Error("LoginRequest Unmarshal Failed!")
		return
	}
	loginResponse, err := handleLogin(ctx, loginRequest)
	data, err := proto.Marshal(&loginResponse)
	if err != nil {
		log.WithError(err).Error("LoginResponse Marshal Failed!")
		return
	}
	handler.Ctx.Output.Body(data)
}

// @Title Register
// @router / [post]
func (handler *RegisterHandler) Register() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	registerRequest := im_entities.RegisterRequest{}
	err := proto.Unmarshal(ctx.Input.RequestBody, &registerRequest)
	if err != nil {
		log.WithError(err).Error("RegisterRequest Unmarshal Failed!")
		return
	}
	registerResponse, err := handleRegister(ctx, registerRequest)
	data, err := proto.Marshal(&registerResponse)
	if err != nil {
		log.WithError(err).Error("RegisterResponse Marshal Failed!")
		return
	}
	handler.Ctx.Output.Body(data)
}

// @Title GetChatList
// @router / [post]
func (handler *ChatHandler) GetChatList() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	getChatListRequest := im_entities.GetChatListRequest{}
	err := proto.Unmarshal(ctx.Input.RequestBody, &getChatListRequest)
	if err != nil {
		log.WithError(err).Error("GetChatListRequest Unmarshal Failed!")
		return
	}
	getChatListResponse, err := handleGetChatList(ctx, getChatListRequest)
	data, err := proto.Marshal(&getChatListResponse)
	if err != nil {
		log.WithError(err).Error("GetChatListResponse Marshal Failed!")
		return
	}
	handler.Ctx.Output.Body(data)
}

// @Title GetContactList
// @router / [post]
func (handler *ContactHandler) GetContactList() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	getContactListRequest := im_entities.GetContactListRequest{}
	err := proto.Unmarshal(ctx.Input.RequestBody, &getContactListRequest)
	if err != nil {
		log.WithError(err).Error("GetContactListRequest Unmarshal Failed!")
		return
	}
	getContactListResponse, err := handleGetContactList(ctx, getContactListRequest)
	data, err := proto.Marshal(&getContactListResponse)
	if err != nil {
		log.WithError(err).Error("GetContactListResponse Marshal Failed!")
		return
	}
	handler.Ctx.Output.Body(data)
}

// @Title AddContact
// @router / [post]
func (handler *ContactHandler) AddContact() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	addContactRequest := im_entities.AddContactRequest{}
	err := proto.Unmarshal(ctx.Input.RequestBody, &addContactRequest)
	if err != nil {
		log.WithError(err).Error("AddContactRequest Unmarshal Failed!")
		return
	}
	addContactResponse, err := handleAddContact(ctx,addContactRequest)
	data, err := proto.Marshal(&addContactResponse)
	if err != nil {
		log.WithError(err).Error("AddContactResponse Marshal Failed!")
		return
	}
	handler.Ctx.Output.Body(data)
}

// @Title CreateChat
// @router / [post]
func (handler *ChatHandler) CreateChat() {
	ctx := *(handler.Ctx)
	log := logger(ctx)
	createChatRequest := im_entities.CreateChatRequest{}
	err := proto.Unmarshal(ctx.Input.RequestBody, &createChatRequest)
	if err != nil {
		log.WithError(err).Error("CreateChatRequest Unmarshal Failed!")
		return
	}
	createChatResponse, err := handleCreateChat(ctx, createChatRequest)
	data, err := proto.Marshal(&createChatResponse)
	if err != nil {
		log.WithError(err).Error("CreateChatResponse Marshal Failed!")
		return
	}
	handler.Ctx.Output.Body(data)
}

func logger(ctx context.Context) *logrus.Entry {
	fields := logrus.Fields{}
	return logrus.WithFields(fields)
}