package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/news"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type NewsHandler struct {
	getAllNewsUseCase news.INewsGetAllUseCase
}

func NewNewsHandler(getAllNewsUseCase news.INewsGetAllUseCase) *NewsHandler {
	return &NewsHandler{
		getAllNewsUseCase: getAllNewsUseCase,
	}
}

func (h *NewsHandler) GetAllNews(w http.ResponseWriter, r *http.Request) {
	news, err := h.getAllNewsUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := mappers.NewsListToResponse(news)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
