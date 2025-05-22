package postgresql

import (
	"context"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type NewsModelRepositoryPostgres struct {
	store db.Store
}

func NewNewsModelRepositoryPostgres(store *db.Store) adapters.INewsRepository {
	return &NewsModelRepositoryPostgres{store: *store}
}

func (r *NewsModelRepositoryPostgres) GetAllNews(ctx context.Context) ([]entities.News, error) {
	news, err := r.store.GetAllNews(ctx)
	if err != nil {
		return []entities.News{}, err
	}
	var newsList []entities.News
	for _, n := range news {
		newsList = append(newsList, entities.News{
			ID:          n.NewsID,
			Slug:        n.Slug,
			Image:       n.ImageUrl,
			Title:       n.Title,
			Description: n.Description,
			Author:      n.Author,
			Content:     n.Content,
		})
	}

	return newsList, nil
}

func (r *NewsModelRepositoryPostgres) GetNews(ctx context.Context, newsID int64) (entities.News, error) {
	news, err := r.store.GetNews(ctx, newsID)
	if err != nil {
		return entities.News{}, err
	}
	return entities.News{
		Slug:        news.Slug,
		Image:       news.ImageUrl,
		Title:       news.Title,
		Description: news.Description,
		Author:      news.Author,
		Content:     news.Content,
		CreateAt:    news.CreatedAt,
	}, nil
}

func (r *NewsModelRepositoryPostgres) CreateNews(ctx context.Context, arg entities.CreateNewsParams) (entities.News, error) {
	news, err := r.store.CreateNews(ctx, db.CreateNewsParams{
		Slug:        arg.Slug,
		ImageUrl:    arg.Image,
		Title:       arg.Title,
		Description: arg.Description,
		Author:      arg.Author,
		Content:     arg.Content,
	})
	if err != nil {
		return entities.News{}, err

	}
	return entities.News{
		Slug:        news.Slug,
		Image:       news.ImageUrl,
		Title:       news.Title,
		Description: news.Description,
		Author:      news.Author,
		Content:     news.Content,
		CreateAt:    news.CreatedAt,
	}, nil
}

func (r *NewsModelRepositoryPostgres) ListNews(ctx context.Context, arg entities.ListNewsParams) ([]entities.News, error) {
	news, err := r.store.ListNews(ctx, db.ListNewsParams{
		Limit:  arg.Limit,
		Offset: arg.Offset,
	})
	if err != nil {
		return []entities.News{}, err
	}
	var newsList []entities.News
	for _, n := range news {
		newsList = append(newsList, entities.News{
			ID:          n.NewsID,
			Slug:        n.Slug,
			Image:       n.ImageUrl,
			Title:       n.Title,
			Description: n.Description,
			Author:      n.Author,
			Content:     n.Content,
			CreateAt:    n.CreatedAt,
		})
	}
	return newsList, nil
}
