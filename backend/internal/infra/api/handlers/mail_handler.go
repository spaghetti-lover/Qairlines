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

func (h *SendMailHandler) SendMail(c *gin.Context) {
	var req entities.Mail

	// Decode request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to decode request body", "error": err.Error()})
		return
	}

	// Basic validation
	if req.To == "" || req.Subject == "" || req.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "to, subject, and body are required fields"})
		return
	}

	emailMessage := dto.MailMessage{
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}

	// Execute use case
	if err := h.mailUseCase.Execute(c.Request.Context(), emailMessage.To, emailMessage.Subject, emailMessage.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to send email", "error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusAccepted, emailMessage)
}
