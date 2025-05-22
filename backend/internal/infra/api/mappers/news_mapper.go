package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func NewsToResponse(news entities.News) dto.NewsResponse {
	return dto.NewsResponse{
		NewsID:      news.ID,
		Image:       news.Image,
		Title:       news.Title,
		Description: news.Description,
		Author:      news.Author,
		Content:     news.Content,
		CreatedAt:   news.CreateAt,
	}
}

func NewsListToResponse(newsList []entities.News) []dto.NewsResponse {
	responses := make([]dto.NewsResponse, len(newsList))
	for i, news := range newsList {
		responses[i] = NewsToResponse(news)
	}
	return responses
}
