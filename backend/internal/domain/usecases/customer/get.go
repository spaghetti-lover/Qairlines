package customer

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type IGetCustomerDetailsUseCase interface {
	Execute(ctx context.Context, uid int64) (*dto.CustomerDetailsResponse, error)
}
type GetCustomerDetailsUseCase struct {
	customerRepository adapters.ICustomerRepository
	tokenMaker         token.Maker
}

func NewGetCustomerDetailsUseCase(customerRepository adapters.ICustomerRepository, tokenMaker token.Maker) IGetCustomerDetailsUseCase {
	return &GetCustomerDetailsUseCase{
		customerRepository: customerRepository,
		tokenMaker:         tokenMaker,
	}
}

func (u *GetCustomerDetailsUseCase) Execute(ctx context.Context, uid int64) (*dto.CustomerDetailsResponse, error) {
	// Lấy thông tin khách hàng từ repository
	customer, err := u.customerRepository.GetCustomerByUID(ctx, uid)
	if err != nil {
		if errors.Is(err, adapters.ErrCustomerNotFound) {
			return nil, adapters.ErrCustomerNotFound
		}
		return nil, err
	}

	// Lấy lịch sử đặt vé từ repository
	bookingHistory, err := u.customerRepository.GetBookingHistoryByUID(ctx, uid)
	if err != nil {
		return nil, err
	}

	// Sử dụng mapper để chuyển đổi entity sang DTO
	return mappers.ToCustomerDetailsResponse(customer, bookingHistory), nil
}
