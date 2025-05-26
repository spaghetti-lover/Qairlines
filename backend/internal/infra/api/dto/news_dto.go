package dto

import "time"

type NewsResponse struct {
	NewsID      int64     `json:"newsId"`
	Image       string    `json:"image"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateNewsRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Image       string `json:"news-image"`
	AuthorID    string `json:"authorId"`
}

type CreateNewsResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	AuthorID    string `json:"authorId"`
	Image       string `json:"image"`
	CreatedAt   string `json:"createdAt"`
}

type UpdateNewsRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	AuthorID    string `json:"authorId"`
}

type UpdateNewsResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	AuthorID    string `json:"authorId"`
	Image       string `json:"image"`
	UpdatedAt   string `json:"updatedAt"`
}
