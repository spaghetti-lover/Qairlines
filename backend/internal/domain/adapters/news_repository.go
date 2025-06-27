package adapters

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var ErrNewsNotFound = errors.New("news post not found")

type INewsRepository interface {
	ListNews(ctx context.Context, page int, limit int) ([]entities.News, error)
	DeleteNewsByID(ctx context.Context, newsID int64) error
	CreateNews(ctx context.Context, news *entities.News) (*entities.News, error)
	UpdateNews(ctx context.Context, news *entities.News) (*entities.News, error)
	GetNews(ctx context.Context, id int64) (entities.News, error)
}
