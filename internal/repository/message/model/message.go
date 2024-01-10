package model

import "time"

type Message struct {
	Id        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Text      string    `db:"text"`
	ChatId    int64     `db:"chat_id"`
}
