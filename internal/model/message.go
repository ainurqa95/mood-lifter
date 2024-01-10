package model

import "time"

type Message struct {
	Id        int
	CreatedAt time.Time
	Text      string
	ChatId    int64
}
