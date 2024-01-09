package model

import (
	"database/sql"
	"time"
)

type User struct {
	UUID      string `db:uid`
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	Name     string `db:"name"`
	UserName string `db:"username"`
	ChatID   int64  `db:"chat_id"`
}
