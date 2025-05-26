package adapters

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type INewsRepository interface {
	GetAllNewsWithAuthor(ctx context.Context) ([]entities.News, error)
}
