package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

type SendMailHandler struct {
	mailUseCase usecases.IMailUseCase
}

func NewSendMailHandler(mailUseCase usecases.IMailUseCase) *SendMailHandler {
	return &SendMailHandler{
		mailUseCase: mailUseCase,
	}
}

func (h *SendMailHandler) SendMail(ctx *gin.Context) {
	var req entities.Mail

	// Decode request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to decode request body", "error": err.Error()})
		return
	}

	// Basic validation
	if req.To == "" || req.Subject == "" || req.Body == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "to, subject, and body are required fields"})
		return
	}

	emailMessage := dto.MailMessage{
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}

	// Execute use case
	if err := h.mailUseCase.Execute(ctx.Request.Context(), emailMessage.To, emailMessage.Subject, emailMessage.Body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to send email", "error": err.Error()})
		return
	}

	// Return response
	ctx.JSON(http.StatusAccepted, emailMessage)
}
