package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func EmailRequestToInput(req dto.MailMessage) entities.Mail {
	return entities.Mail{
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}
}
