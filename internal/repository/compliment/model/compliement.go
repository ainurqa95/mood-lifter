package model

import "time"

type Compliment struct {
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}
