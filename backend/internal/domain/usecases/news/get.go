package news

import (
	"context"
	"errors"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

type IGetNewsUseCase interface {
	Execute(ctx context.Context, newsID int64) (*dto.GetNewsResponse, error)
}

type GetNewsUseCase struct {
	newsRepository adapters.INewsRepository
}

func NewGetNewsUseCase(newsRepository adapters.INewsRepository) IGetNewsUseCase {
	return &GetNewsUseCase{
		newsRepository: newsRepository,
	}
}

func (u *GetNewsUseCase) Execute(ctx context.Context, newsID int64) (*dto.GetNewsResponse, error) {
	// Lấy bài viết từ repository
	news, err := u.newsRepository.GetNews(ctx, newsID)
	if err != nil {
		if errors.Is(err, adapters.ErrNewsNotFound) {
			return nil, ErrNewsNotFound
		}
		return nil, err
	}

	// Map entity sang DTO
	return &dto.GetNewsResponse{
		ID:          strconv.FormatInt(news.ID, 10),
		Title:       news.Title,
		Description: news.Description,
		Content:     news.Content,
		AuthorID:    strconv.FormatInt(news.AuthorID, 10),
		Image:       news.Image,
		CreatedAt:   news.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   news.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}
