package usecases

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
)

// IMailUseCase defines the interface for the mail use case.
// Benifit: Handler only depend on the use case interface, not the implementation.
// This allows for easier testing and decoupling of components.
type IMailUseCase interface {
	Execute(ctx context.Context, to string, subject string, body string) error
}

// MailUseCase implements the IMailUseCase interface.
type MailUseCase struct {
	emailRepository adapters.IEmailRepository
}

// NewMailUseCase creates a new instance of MailUseCase.
func NewMailUseCase(emailRepository adapters.IEmailRepository) IMailUseCase {
	return &MailUseCase{
		emailRepository: emailRepository,
	}
}

// Execute sends an email using the provided email repository.
func (r *MailUseCase) Execute(ctx context.Context, to string, subject string, body string) error {
	return r.emailRepository.Send(ctx, to, subject, body)
}
