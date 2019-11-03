# IM BackEnd Data Transporation Design Doc

**Date:** 2019.10.13

## Goal

This document try to clarify which data type we use to communicate with front end. That is, what data type we handle in backend.

## ProtoBuf

ProtoBuf, PB in short, is a data type used to transport data between front end and back end. It's more convenient to use PB than json.

We only need to edit `.proto` file. And use the script to generate `.go` automatically. Those generated files can be used in backend code.

---

## PB Format

### Register

```protobuf
message RegisterRequest {
    optional string username = 1;
    optional string password = 2;
    optional string email = 3;
}

message RegisternResponse {
    optional string code = 1;
    optional string message = 2;
}
```

code:

+ 200: Succeed
+ xxx: Failed

message:

+ Duplicate Email: "Email existed"

### Login

```protobuf
message LoginRequest {
    optional string email = 1;
    optional string password = 2;
}

message LoginResponse {
	optional string code = 1;
    optional string message = 2;
    optional string userid = 3;
    optional string username = 4;    
}
```

code:

+ 200: Succeed

message:

+ 

### GetContactList

```proto
message GetContactListRequest {
	
}

message GetContactListResponse {

}
```



## Reference

1. [PB Usage in Go](https://www.jianshu.com/p/c1723e5f6a46)