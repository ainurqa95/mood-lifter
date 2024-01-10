package model

import (
	"time"
)

type User struct {
	Id        int
	UUID      string
	Info      UserInfo
	CreatedAt time.Time
}

type UserInfo struct {
	Name     string
	UserName string
	ChatID   int64
}
