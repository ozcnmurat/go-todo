package models

import "time"

type TodoList struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Completion float32    `json:"completion"`
	UserID     uint       `json:"user_id"`
	Items      []TodoItem `json:"items"`
}

type TodoItem struct {
	ID          uint       `json:"id"`
	TodoListID  uint       `json:"todo_list_id"`
	Content     string     `json:"content" binding:"required"`
	IsCompleted bool       `json:"is_completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
