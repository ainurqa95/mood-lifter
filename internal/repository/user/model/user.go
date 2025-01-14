package model

import (
	"time"
)

type User struct {
	Id        int       `db:"id"`
	UUID      string    `db:"uid"`
	Name      *string   `db:"name"`
	UserName  *string   `db:"username"`
	ChatID    int64     `db:"chat_id"`
	CreatedAt time.Time `db:"created_at"`
}

type UserInfo struct {
	Name     *string `db:"name"`
	UserName *string `db:"username"`
	ChatID   int64   `db:"chat_id"`
}
