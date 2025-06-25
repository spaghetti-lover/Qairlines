package user

// import (
// 	"context"
// 	"database/sql"
// 	"errors"
// 	"fmt"

// 	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
// 	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
// 	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/customer"
// )

// type IUserCreateUseCase interface {
// 	Execute(ctxtx context.Context, arg entities.CreateUserParams) (entities.User, error)
// }

// type UserCreateUseCase struct {
// 	userRepository        adapters.IUserRepository
// 	customerCreateUseCase customer.ICreateCustomerUseCase
// }

// func NewUserCreateUseCase(userRepository adapters.IUserRepository, customerCreateUseCase customer.ICreateCustomerUseCase) IUserCreateUseCase {
// 	return &UserCreateUseCase{
// 		userRepository:        userRepository,
// 		customerCreateUseCase: customerCreateUseCase,
// 	}
// }

// func (u *UserCreateUseCase) Execute(ctxtx context.Context, arg entities.CreateUserParams) (entities.User, error) {
// 	// Kiểm tra email đã tồn tại
// 	existingUser, err := u.userRepository.GetUserByEmail(ctx, arg.Email)
// 	if err != nil && !errors.Is(err, sql.ErrNoRows) {
// 		return entities.User{}, err
// 	}
// 	if existingUser != nil {
// 		return entities.User{}, fmt.Errorf("email already in use")
// 	}

// 	// Tạo user
// 	user, err := u.userRepository.CreateUser(ctx, arg)
// 	if err != nil {
// 		return entities.User{}, fmt.Errorf("failed to create user: %w", err)
// 	}

// 	if err != nil {
// 		return entities.User{}, fmt.Errorf("failed to create customer: %w", err)
// 	}

// 	return user, nil
// }
