package postgresql

import (
	"context"
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

func (r *NewsModelRepositoryPostgres) GetAllNewsWithAuthor(ctx context.Context) ([]entities.News, error) {
	news, err := r.store.GetAllNewsWithAuthor(ctx)
	if err != nil {
		return []entities.News{}, err
	}
	var newsList []entities.News
	for _, n := range news {
		newsList = append(newsList, entities.News{
			ID:        n.ID,
			Title:     n.Title,
			Content:   n.Content.String,
			Image:     n.Image.String,
			Author:    n.FirstName.String + n.LastName.String,
			CreatedAt: n.CreatedAt,
			UpdatedAt: n.UpdatedAt,
		})
	}

	return newsList, nil
}

func (r *NewsModelRepositoryPostgres) DeleteNewsByID(ctx context.Context, newsID int64) error {
	rowsAffected, err := r.store.DeleteNews(ctx, newsID)
	if err != nil {
		return fmt.Errorf("failed to delete news post: %w", err)
	}

	if rowsAffected == 0 {
		return adapters.ErrNewsNotFound // Trả về lỗi nếu không có hàng nào bị xóa
	}

	return nil
}
