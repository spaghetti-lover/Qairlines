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
