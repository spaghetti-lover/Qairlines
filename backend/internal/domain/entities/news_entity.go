package entities

import "time"

type News struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Image       string    `json:"image"`
	AuthorID    int64     `json:"author_id"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateNewsParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	AuthorID    int64  `json:"author_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ListNewsParams struct {
	Page  int `form:"page" binding:"required"`
	Limit int `form:"limit" binding:"required"`
}
