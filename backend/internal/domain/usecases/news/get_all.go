package news

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IGetAllNewsWithAuthor interface {
	Execute(ctx context.Context) ([]entities.News, error)
}

type NewsGetAllNewsWithAuthorUseCase struct {
	newsRepository adapters.INewsRepository
}

func NewNewsGetAllWithAuthorUseCase(newsRepository adapters.INewsRepository) IGetAllNewsWithAuthor {
	return &NewsGetAllNewsWithAuthorUseCase{
		newsRepository: newsRepository,
	}
}

func (r *NewsGetAllNewsWithAuthorUseCase) Execute(ctx context.Context) ([]entities.News, error) {
	news, err := r.newsRepository.GetAllNewsWithAuthor(ctx)
	if err != nil {
		return []entities.News{}, err
	}
	return news, nil
}
