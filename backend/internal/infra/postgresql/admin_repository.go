package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type AdminRepositoryPostgres struct {
	store      db.Store
	tokenMaker token.Maker
}

func NewAdminRepositoryPostgres(store *db.Store, tokenMaker token.Maker) *AdminRepositoryPostgres {
	return &AdminRepositoryPostgres{
		store:      *store,
		tokenMaker: tokenMaker,
	}
}

func (r *AdminRepositoryPostgres) CreateAdminTx(ctx context.Context, arg entities.CreateUserParams) (entities.User, error) {
	user, err := r.store.CreateAdminTx(ctx, db.CreateUserParams{
		FirstName:      &arg.FirstName,
		LastName:       &arg.LastName,
		HashedPassword: arg.Password,
		Email:          arg.Email,
		Role:           db.UserRoleAdmin,
	})

	if err != nil {
		return entities.User{}, err
	}
	return entities.User{
		UserID:    user.UserID,
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     user.Email,
		HashedPwd: user.HashedPassword,
		Role:      entities.UserRole(entities.RoleAdmin),
	}, nil
}

func (r *AdminRepositoryPostgres) GetAllAdmins(ctx context.Context) ([]entities.Admin, error) {
	adminIDs, err := r.store.GetAllAdmin(ctx)
	if err != nil {
		return nil, err
	}

	admins := make([]entities.Admin, len(adminIDs))
	for i, adminID := range adminIDs {
		user, err := r.store.GetUser(ctx, adminID)
		if err != nil {
			return nil, err
		}

		admins[i] = entities.Admin{
			UserID:    strconv.FormatInt(user.UserID, 10),
			FirstName: *user.FirstName,
			LastName:  *user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
	}

	return admins, nil
}

func (r *AdminRepositoryPostgres) GetAdminByID(ctx context.Context, adminID string) (entities.Admin, error) {
	id, err := strconv.ParseInt(adminID, 10, 64)
	if err != nil {
		return entities.Admin{}, err
	}

	// Kiểm tra xem người dùng có phải admin không
	isAdmin, err := r.store.IsAdmin(ctx, id)
	if err != nil {
		return entities.Admin{}, err
	}
	if !isAdmin {
		return entities.Admin{}, errors.New("user is not an admin")
	}

	// Lấy thông tin user
	user, err := r.store.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Admin{}, errors.New("admin not found")
		}
		return entities.Admin{}, err
	}

	return entities.Admin{
		UserID:    strconv.FormatInt(user.UserID, 10),
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *AdminRepositoryPostgres) GetAdminByUserID(ctx context.Context, userID int64) (entities.Admin, error) {
	// Kiểm tra xem người dùng có phải admin không
	isAdmin, err := r.store.IsAdmin(ctx, userID)
	if err != nil {
		return entities.Admin{}, err
	}
	if !isAdmin {
		return entities.Admin{}, errors.New("user is not an admin")
	}

	// Lấy thông tin user
	user, err := r.store.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Admin{}, errors.New("admin not found")
		}
		return entities.Admin{}, err
	}

	return entities.Admin{
		UserID:    strconv.FormatInt(user.UserID, 10),
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *AdminRepositoryPostgres) UpdateAdmin(ctx context.Context, admin entities.Admin) (entities.Admin, error) {
	userID, err := strconv.ParseInt(admin.UserID, 10, 64)
	if err != nil {
		return entities.Admin{}, err
	}

	isAdmin, err := r.store.IsAdmin(ctx, userID)
	if err != nil {
		return entities.Admin{}, err
	}
	if !isAdmin {
		return entities.Admin{}, errors.New("user is not an admin")
	}

	err = r.store.UpdateUser(ctx, db.UpdateUserParams{
		UserID:    userID,
		FirstName: &admin.FirstName,
		LastName:  &admin.LastName,
	})
	if err != nil {
		return entities.Admin{}, err
	}

	updatedUser, err := r.store.GetUser(ctx, userID)
	if err != nil {
		return entities.Admin{}, err
	}

	return entities.Admin{
		UserID:    strconv.FormatInt(updatedUser.UserID, 10),
		FirstName: *updatedUser.FirstName,
		LastName:  *updatedUser.LastName,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt,
	}, nil
}

func (r *AdminRepositoryPostgres) DeleteAdmin(ctx context.Context, userID int64) error {
	// Sử dụng transaction để đảm bảo tính toàn vẹn dữ liệu
	result, err := r.store.DeleteAdminTx(ctx, db.DeleteAdminTxParams{
		UserID: userID,
	})

	if err != nil {
		log.Printf("Error deleting admin (user_id=%d): %v", userID, err)

		// Xử lý các loại lỗi phổ biến
		if strings.Contains(err.Error(), "not found") {
			return errors.New("admin not found")
		} else if strings.Contains(err.Error(), "is not an admin") {
			return errors.New("user is not an admin")
		}

		return fmt.Errorf("failed to delete admin: %w", err)
	}

	if !result.Success {
		return errors.New("delete admin transaction failed")
	}

	log.Printf("Admin (user_id=%d) deleted successfully", userID)
	return nil
}
