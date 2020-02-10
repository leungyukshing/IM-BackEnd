package entities

import "time"

type User struct {
	ID          int64
	Username    string
	Password    string
	Email       string
	CreatedTime time.Time
	EncryptKey  string
}
