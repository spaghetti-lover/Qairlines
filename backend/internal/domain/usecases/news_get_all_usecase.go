package usecases

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type INewsGetAllUseCase interface {
	Execute(ctx context.Context) ([]entities.News, error)
}

type NewsGetAllUseCase struct {
	newsRepository adapters.INewsRepository
}

func NewNewsGetAllUseCase(newsRepository adapters.INewsRepository) INewsGetAllUseCase {
	return &NewsGetAllUseCase{
		newsRepository: newsRepository,
	}
}

func (r *NewsGetAllUseCase) Execute(ctx context.Context) ([]entities.News, error) {
	news, err := r.newsRepository.GetAllNews(ctx)
	if err != nil {
		return []entities.News{}, err
	}
	return news, nil
}
