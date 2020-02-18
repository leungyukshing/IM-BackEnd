# IM BackEnd Data Transporation Design Doc

**Date:** 2019.10.13

## Goal

This document try to clarify which data type we use to communicate with front end. That is, what data type we handle in backend.

## ProtoBuf

ProtoBuf, PB in short, is a data type used to transport data between front end and back end. It's more convenient to use PB than json.

We only need to edit `.proto` file. And use the script to generate `.go` automatically. Those generated files can be used in backend code.

---

## PB Format

### User

```protobuf
message User {
    optional string userid = 1;
    optional string username = 2;
    optional string email = 3;
}
```

### Chat

```protobuf
message Chat {
    optional string chatid = 1;
    optional string userid = 2; // ownerID
    optional string chatname = 3;
    optional string last_update_time = 4;
}
```

### Register

```protobuf
message RegisterRequest {
    optional string username = 1;
    optional string password = 2;
    optional string email = 3;
}

message RegisterResponse {
    optional string code = 1;
    optional string message = 2;
}
```

code:

+ 200: Request Succeed
+ 500: Server Database Error

message:

+ Some Field Empty: "Empty Field"

+ Duplicate Email: "Email Existed"
+ Email Format Invalid: "Invalid Email"
+ DB Error: "DB_ERR"
+ Success: "Register Success"

### Login

```protobuf
message LoginRequest {
    optional string email = 1;
    optional string password = 2;
}

message LoginResponse {
	optional string code = 1;
    optional string message = 2;
    optional User user = 3;
}
```

code:

+ 200: Request Succeed
+ 500: Server Database Error

message:

+ Some Field Empty: "Email or Password Empty"
+ DB Error: "DB_ERR"
+ Verfity Fail: "Email or Password Incorrect"
+ Success: "Login Success"

### GetChatList

```protobuf
message GetChatListRequest {
    optional string userid = 1;
}

message GetChatListResponse {
    optional string code = 1;
    optional string message = 2;
    repeated Chat chat = 3;
}
```

code:

+ 200: Request Succeed
+ 500: Server Database Error

message:

+ Field Empty: "UserID Empty"
+ DB Error: "DB_ERR"
+ Success: "GetChatList Success"

### GetContactList

```proto
message GetContactListRequest {
    optional string userid = 1;
}

message GetContactListResponse {
    optional string code = 1;
    optional string message = 2;
    repeated User user = 3;
}
```

code:

+ 200: Request Succeed
+ 500: Server Database Error

message:

+ Field Empty: "UserID Empty"
+ DB Error: "DB_ERR"
+ Success: "GetContactList Success"

### SendMessage

```protobuf
message SendMessageRequest {
    optional string senderid = 1;
    optional string chatid = 2;
    optional string message =3;
}

message SendMessageResponse {
    optional string code = 1;
}
```

code:

+ 200: Request Succeed
+ 500: Server Database Error

message:

### CreateChat

```protobuf
message CreateChatRequest {
    optional string userid = 1;
    optional string receiverid = 2;
}

message CreateChatResponse {
    optional string code = 1;
    optional string message = 2;
    optional Chat chat = 3;
}
```

code:

+ 200: Request Succeed
+ 500: Server Database Error

message:

+ Field Empty: "Some Field is Empty"
+ DB Error: "DB_ERR"
+ Chat already Existed: "Chat Existed"
+ Success: "CreateChat Success"

### AddContact

```protobuf
message AddContactRequest {
    optional string userid = 1;
    optional string receiverid = 2;
}

message AddContactResponse {
    optional string code = 1;
    optional string message = 2;
}
```





## Reference

1. [PB Usage in Go](https://www.jianshu.com/p/c1723e5f6a46)