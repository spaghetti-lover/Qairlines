package entities

import "time"

type News struct {
	ID          int64     `json:"id"`
	Slug        string    `json:"slug"`
	Image       string    `json:"image"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Content     string    `json:"content"`
	CreateAt    time.Time `json:"createAt"`
}

type CreateNewsParams struct {
	Slug        string `json:"slug"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Content     string `json:"content"`
}

type ListNewsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}
