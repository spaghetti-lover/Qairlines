package user

// import (
// 	"context"

// 	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
// 	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
// )

// type IUserGetUseCase interface {
// 	Execute(ctxtx context.Context, userID int64) (entities.User, error)
// }

// type UserGetUseCase struct {
// 	userRepository adapters.IUserRepository
// }

// func NewUserGetUseCase(userRepository adapters.IUserRepository) IUserGetUseCase {
// 	return &UserGetUseCase{
// 		userRepository: userRepository,
// 	}
// }
// func (r *UserGetUseCase) Execute(ctxtx context.Context, userID int64) (entities.User, error) {
// 	user, err := r.userRepository.GetUser(ctx, userID)
// 	if err != nil {
// 		return entities.User{}, err
// 	}
// 	return user, nil
// }
