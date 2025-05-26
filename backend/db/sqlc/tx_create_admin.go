package db

import (
	"context"
)

func (store *SQLStore) CreateAdminTx(ctx context.Context, arg CreateUserParams) (User, error) {
	var user User

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		user, err = q.CreateUser(ctx, CreateUserParams{
			FirstName:      arg.FirstName,
			LastName:       arg.LastName,
			HashedPassword: arg.HashedPassword,
			Email:          arg.Email,
			Role:           UserRoleAdmin,
		})
		if err != nil {
			return err
		}
		_, err = q.CreateAdmin(ctx, user.UserID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return User{}, err
	}

	return user, nil
}
