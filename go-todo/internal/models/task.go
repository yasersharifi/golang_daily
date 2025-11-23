package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TaskInput struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TaskOutput struct {
	ID int `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
