package news

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

type IGetNewsUseCase interface {
	Execute(ctx context.Context, newsID int64) (*dto.GetNewsResponse, error)
}

type GetNewsUseCase struct {
	newsRepository  adapters.INewsRepository
	cacheRepository adapters.ICacheRepository
}

func NewGetNewsUseCase(newsRepository adapters.INewsRepository, cacheRepository adapters.ICacheRepository) IGetNewsUseCase {
	return &GetNewsUseCase{
		newsRepository:  newsRepository,
		cacheRepository: cacheRepository,
	}
}

func (u *GetNewsUseCase) Execute(ctx context.Context, newsID int64) (*dto.GetNewsResponse, error) {
	// Create dynamic key
	cacheKey := fmt.Sprintf("getNews:%d", newsID)
	// Get data from cache
	var cachedData dto.GetNewsResponse
	if err := u.cacheRepository.Get(cacheKey, &cachedData); err == nil {
		return &cachedData, nil
	}

	// Lấy bài viết từ repository
	news, err := u.newsRepository.GetNews(ctx, newsID)
	if err != nil {
		return nil, err
	}
	newsDTO := &dto.GetNewsResponse{
		ID:          news.ID,
		Title:       news.Title,
		Description: news.Description,
		Content:     news.Content,
		AuthorID:    news.AuthorID,
		Image:       news.Image,
		CreatedAt:   news.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   news.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	// Serialize DTO to JSON
	jsonData, err := json.Marshal(newsDTO)
	if err != nil {
		return nil, err
	}

	// Create cache data
	err = u.cacheRepository.Set(cacheKey, jsonData, 5*time.Minute)
	if err != nil {
		return nil, err
	}

	// Map entity sang DTO
	return newsDTO, nil
}
