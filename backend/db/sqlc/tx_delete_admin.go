package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

// DeleteAdminTxParams chứa thông tin cần thiết để xoá admin
type DeleteAdminTxParams struct {
	UserID int64 `json:"user_id"`
}

// DeleteAdminTxResult chứa kết quả của transaction xoá admin
type DeleteAdminTxResult struct {
	Success bool `json:"success"`
}

// DeleteAdminTx thực hiện transaction xoá hoàn toàn một admin và user tương ứng
func (store *SQLStore) DeleteAdminTx(ctx context.Context, arg DeleteAdminTxParams) (DeleteAdminTxResult, error) {
	var result DeleteAdminTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// 1. Kiểm tra xem user có tồn tại không
		user, err := q.GetUser(ctx, arg.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("user with ID %d not found", arg.UserID)
			}
			return err
		}

		// 2. Kiểm tra xem user có phải admin không
		isAdmin, err := q.IsAdmin(ctx, user.UserID)
		if err != nil {
			return err
		}
		if !isAdmin {
			return fmt.Errorf("user with ID %d is not an admin", arg.UserID)
		}

		// 3. Xoá các liên kết trong bảng Bookings (nếu có)
		err = q.RemoveUserFromBookings(ctx, pgtype.Text{String: user.Email, Valid: true})
		if err != nil {
			return err
		}

		// 4. Xoá các liên kết trong bảng Blog Posts (nếu user là tác giả)
		err = q.RemoveAuthorFromBlogPosts(ctx, pgtype.Int8{Int64: user.UserID, Valid: true})
		if err != nil {
			return err
		}

		// 5. Xoá khỏi bảng Admin
		err = q.DeleteAdmin(ctx, user.UserID)
		if err != nil {
			return err
		}

		// 6. Cuối cùng, xoá user
		err = q.DeleteUser(ctx, user.UserID)
		if err != nil {
			return err
		}

		result.Success = true
		return nil
	})

	return result, err
}
