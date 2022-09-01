package model

import "time"

type TodoList struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Done        bool      `json:"done"`
	CurrentTime time.Time `json:"time" binding:"required"`
	ListID      int       `json:"ListID" binding:"required"`
}
