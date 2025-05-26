package news

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
)

type IDeleteNewsUseCase interface {
	Execute(ctx context.Context, newsID int64) error
}

type DeleteNewsUseCase struct {
	newsRepository adapters.INewsRepository
}

func NewDeleteNewsUseCase(newsRepository adapters.INewsRepository) IDeleteNewsUseCase {
	return &DeleteNewsUseCase{
		newsRepository: newsRepository,
	}
}

func (u *DeleteNewsUseCase) Execute(ctx context.Context, newsID int64) error {
	return u.newsRepository.DeleteNewsByID(ctx, newsID)
}
