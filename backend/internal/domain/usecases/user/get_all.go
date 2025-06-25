package user

// import (
// 	"context"

// 	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
// 	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
// )

// // UserGetAllUseCase is a use case for getting all users.
// type IUserGetAllUseCase interface {
// 	Execute(ctxtx context.Context) ([]entities.User, error)
// }

// // UserGetAllUseCase is a use case for getting all users.
// type UserGetAllUseCase struct {
// 	userRepository adapters.IUserRepository
// }

// func NewUserGetAllUseCase(userRepository adapters.IUserRepository) IUserGetAllUseCase {
// 	return &UserGetAllUseCase{
// 		userRepository: userRepository,
// 	}
// }

// func (r *UserGetAllUseCase) Execute(ctxtx context.Context) ([]entities.User, error) {
// 	users, err := r.userRepository.GetAllUser(ctx)
// 	if err != nil {
// 		return []entities.User{}, err
// 	}
// 	return users, nil
// }
