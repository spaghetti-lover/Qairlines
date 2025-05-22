package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type SendMailHandler struct {
	mailUseCase usecases.IMailUseCase
}

func NewSendMailHandler(mailUseCase usecases.IMailUseCase) *SendMailHandler {
	return &SendMailHandler{
		mailUseCase: mailUseCase,
	}
}

func (h *SendMailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req entities.Mail

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "failed to decode request body", err)
		return
	}

	// Basic validation
	if req.To == "" || req.Subject == "" || req.Body == "" {
		utils.WriteError(w, http.StatusBadRequest, "to, subject, and body are required fields", nil)
		return
	}

	emailMessage := dto.MailMessage{
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}

	if err := h.mailUseCase.Execute(r.Context(), emailMessage.To, emailMessage.Subject, emailMessage.Body); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to send email", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(emailMessage); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to encode response", err)
		return
	}
}
