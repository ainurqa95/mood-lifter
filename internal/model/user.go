package model

import (
	"time"
)

type User struct {
	UUID      string
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserInfo struct {
	Name     string
	UserName string
	ChatID   int64
}
