package todo

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	IsDone    bool      `json:"isDone"`
}
