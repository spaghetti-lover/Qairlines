package news

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

type ICreateNewsUseCase interface {
	Execute(ctx context.Context, req dto.CreateNewsRequest) (*dto.CreateNewsResponse, error)
}

type CreateNewsUseCase struct {
	newsRepository adapters.INewsRepository
}

func NewCreateNewsUseCase(newsRepository adapters.INewsRepository) ICreateNewsUseCase {
	return &CreateNewsUseCase{
		newsRepository: newsRepository,
	}
}

func (u *CreateNewsUseCase) Execute(ctx context.Context, req dto.CreateNewsRequest) (*dto.CreateNewsResponse, error) {
	// Validate input
	if req.Title == "" || req.Description == "" || req.Content == "" || req.AuthorID == "" {
		return nil, ErrInvalidNewsData
	}

	authorID, err := strconv.ParseInt(req.AuthorID, 10, 64)
	if err != nil {
		return nil, errors.New("invalid author ID: " + err.Error())
	}

	// Tạo entity News
	news := &entities.News{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		Image:       req.Image,
		AuthorID:    authorID,
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
		ID:          strconv.FormatInt(createdNews.ID, 10),
		Title:       createdNews.Title,
		Description: createdNews.Description,
		Content:     createdNews.Content,
		AuthorID:    strconv.FormatInt(createdNews.AuthorID, 10),
		Image:       createdNews.Image,
		CreatedAt:   createdNews.CreatedAt.Format(time.RFC3339),
	}, nil
}
