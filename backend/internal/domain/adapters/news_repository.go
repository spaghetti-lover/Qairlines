package adapters

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)


type INewsRepository interface {
	GetNews(ctx context.Context, newsID int64) (entities.News, error)
	CreateNews(ctx context.Context, arg entities.CreateNewsParams) (entities.News, error)
	ListNews(ctx context.Context, arg entities.ListNewsParams) ([]entities.News, error)
	GetAllNews(ctx context.Context) ([]entities.News, error)
}