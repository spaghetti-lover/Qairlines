package adapters

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var ErrNewsNotFound = errors.New("news post not found")

type INewsRepository interface {
	GetAllNewsWithAuthor(ctx context.Context) ([]entities.News, error)
	DeleteNewsByID(ctx context.Context, newsID int64) error
}
