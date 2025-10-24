package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"desc"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}
