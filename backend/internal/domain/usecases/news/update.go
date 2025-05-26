package news

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

var ErrNewsNotFound = errors.New("news post not found")
var ErrInvalidNewsData = errors.New("invalid news data")

type IUpdateNewsUseCase interface {
	Execute(ctx context.Context, newsID int64, req dto.UpdateNewsRequest) (*dto.UpdateNewsResponse, error)
}

type UpdateNewsUseCase struct {
	newsRepository adapters.INewsRepository
}

func NewUpdateNewsUseCase(newsRepository adapters.INewsRepository) IUpdateNewsUseCase {
	return &UpdateNewsUseCase{
		newsRepository: newsRepository,
	}
}

func (u *UpdateNewsUseCase) Execute(ctx context.Context, newsID int64, req dto.UpdateNewsRequest) (*dto.UpdateNewsResponse, error) {
	// Validate input
	if req.Title == "" || req.Description == "" || req.Content == "" || req.AuthorID == "" {
		return nil, ErrInvalidNewsData
	}

	// Lấy bài viết hiện tại
	existingNews, err := u.newsRepository.GetNews(ctx, newsID)
	if err != nil {
		if errors.Is(err, adapters.ErrNewsNotFound) {
			return nil, ErrNewsNotFound
		}
		return nil, err
	}

	// Cập nhật bài viết
	existingNews.Title = req.Title
	existingNews.Description = req.Description
	existingNews.Content = req.Content
	authorID, err := strconv.ParseInt(req.AuthorID, 10, 64)
	if err != nil {
		return nil, ErrInvalidNewsData
	}
	existingNews.AuthorID = authorID
	existingNews.UpdatedAt = time.Now()

	updatedNews, err := u.newsRepository.UpdateNews(ctx, &existingNews)
	if err != nil {
		return nil, err
	}

	// Map entity sang DTO
	return &dto.UpdateNewsResponse{
		ID:          strconv.FormatInt(updatedNews.ID, 10),
		Title:       updatedNews.Title,
		Description: updatedNews.Description,
		Content:     updatedNews.Content,
		AuthorID:    strconv.FormatInt(updatedNews.AuthorID, 10),
		Image:       updatedNews.Image,
		UpdatedAt:   updatedNews.UpdatedAt.Format(time.RFC3339),
	}, nil
}
