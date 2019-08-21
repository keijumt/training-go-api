package domain

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"Title"`
	IsCompleted bool   `json:"IsCompleted"`
}

type Tasks []Task
