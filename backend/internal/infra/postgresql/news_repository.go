package postgresql

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *NewsModelRepositoryPostgres) ListNews(ctx context.Context, offset int, limit int) ([]entities.News, error) {
	news, err := r.store.ListNews(ctx, db.ListNewsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return []entities.News{}, err
	}
	var newsList []entities.News
	for _, n := range news {
		newsList = append(newsList, entities.News{
			ID:        n.ID,
			Title:     n.Title,
			Content:   *n.Content,
			Image:     *n.Image,
			AuthorID:  n.ID,
			CreatedAt: n.CreatedAt,
			UpdatedAt: n.UpdatedAt,
		})
	}

	return newsList, nil
}

func (r *NewsModelRepositoryPostgres) DeleteNewsByID(ctx context.Context, newsID int64) error {
	rowsAffected, err := r.store.DeleteNews(ctx, newsID)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // Trả về lỗi nếu không có hàng nào bị xóa
	}

	return nil
}

func (r *NewsModelRepositoryPostgres) CreateNews(ctx context.Context, news *entities.News) (*entities.News, error) {
	newsModel := db.CreateNewsParams{
		Title:       news.Title,
		Description: &news.Description,
		Content:     &news.Content,
		Image:       &news.Image,
		AuthorID:    &news.AuthorID,
	}

	createdNews, err := r.store.CreateNews(ctx, newsModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create news post: %w", err)
	}

	return &entities.News{
		ID:          createdNews.ID,
		Title:       createdNews.Title,
		Description: *createdNews.Description,
		Content:     *createdNews.Content,
		Image:       *createdNews.Image,
		AuthorID:    *createdNews.AuthorID,
		CreatedAt:   createdNews.CreatedAt,
		UpdatedAt:   createdNews.UpdatedAt,
	}, nil
}

func (r *NewsModelRepositoryPostgres) UpdateNews(ctx context.Context, news *entities.News) (*entities.News, error) {
	newsModel := db.UpdateNewsParams{
		ID:          news.ID,
		Title:       news.Title,
		Description: &news.Description,
		Content:     &news.Content,
		Image:       &news.Image,
		AuthorID:    &news.AuthorID,
	}

	updatedNews, err := r.store.UpdateNews(ctx, newsModel)
	if err != nil {
		return nil, fmt.Errorf("failed to update news post: %w", err)
	}

	return &entities.News{
		ID:          updatedNews.ID,
		Title:       updatedNews.Title,
		Description: *updatedNews.Description,
		Content:     *updatedNews.Content,
		Image:       *updatedNews.Image,
		AuthorID:    *updatedNews.AuthorID,
		CreatedAt:   updatedNews.CreatedAt,
		UpdatedAt:   updatedNews.UpdatedAt,
	}, nil
}

func (r *NewsModelRepositoryPostgres) GetNews(ctx context.Context, newsID int64) (entities.News, error) {
	news, err := r.store.GetNews(ctx, newsID)
	if err != nil {
		return entities.News{}, err
	}

	return entities.News{
		ID:          news.ID,
		Title:       news.Title,
		Description: *news.Description,
		Content:     *news.Content,
		Image:       *news.Image,
		AuthorID:    *news.AuthorID,
		CreatedAt:   news.CreatedAt,
		UpdatedAt:   news.UpdatedAt,
	}, nil
}
