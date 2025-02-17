package agent

import "time"

type Base struct {
	ID uint32 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserMessage struct {
	Base
	UserId uint32 `json:"user_id"`
	Content string `json:"content"`
}
