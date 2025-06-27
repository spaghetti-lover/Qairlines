package news

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IListNewsUseCase interface {
	Execute(ctx *gin.Context, page int, limit int) ([]entities.News, error)
}

type ListNewsUseCase struct {
	newsRepository adapters.INewsRepository
}

func NewListNewsUseCase(newsRepository adapters.INewsRepository) IListNewsUseCase {
	return &ListNewsUseCase{
		newsRepository: newsRepository,
	}
}

func (r *ListNewsUseCase) Execute(ctx *gin.Context, page int, limit int) ([]entities.News, error) {
	start := (page - 1) * limit
	news, err := r.newsRepository.ListNews(ctx, start, limit)
	if err != nil {
		return nil, err
	}
	return news, nil
}
