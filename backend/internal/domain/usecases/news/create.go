package news

import (
	"context"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

type ICreateNewsUseCase interface {
	Execute(ctx context.Context, req dto.CreateNewsToDBRequest) (*dto.CreateNewsResponse, error)
}

type CreateNewsUseCase struct {
	newsRepository adapters.INewsRepository
}

func NewCreateNewsUseCase(newsRepository adapters.INewsRepository) ICreateNewsUseCase {
	return &CreateNewsUseCase{
		newsRepository: newsRepository,
	}
}

func (u *CreateNewsUseCase) Execute(ctx context.Context, req dto.CreateNewsToDBRequest) (*dto.CreateNewsResponse, error) {
	// Validate input
	if req.Title == "" || req.Description == "" || req.Content == "" {
		return nil, ErrInvalidNewsData
	}

	// Tạo entity News
	news := &entities.News{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		AuthorID:    req.AuthorID,
		Image:       req.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Lưu vào database
	createdNews, err := u.newsRepository.CreateNews(ctx, news)
	if err != nil {
		return nil, err
	}

	// Map entity sang DTO
	return &dto.CreateNewsResponse{
		ID:          createdNews.ID,
		Title:       createdNews.Title,
		Description: createdNews.Description,
		Content:     createdNews.Content,
		AuthorID:    createdNews.AuthorID,
		Image:       createdNews.Image,
		CreatedAt:   createdNews.CreatedAt.Format(time.RFC3339),
	}, nil
}
