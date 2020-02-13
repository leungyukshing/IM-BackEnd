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

type Chat struct {
	ID int64
	UserID int64
	Name string
	CreatedTIme time.Time
	LastUpdateTime time.Time
}