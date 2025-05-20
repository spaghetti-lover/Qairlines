package handlers

import (
	"encoding/json"
	"net/http"

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

func (h *SendMailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req entities.Mail

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.To == "" || req.Subject == "" || req.Body == "" {
		http.Error(w, "to, subject, and body are required fields", http.StatusBadRequest)
		return
	}

	emailMessage := dto.MailMessage{
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}

	if err := h.mailUseCase.Execute(r.Context(), emailMessage.To, emailMessage.Subject, emailMessage.Body); err != nil {
		http.Error(w, "failed to send email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(emailMessage); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
