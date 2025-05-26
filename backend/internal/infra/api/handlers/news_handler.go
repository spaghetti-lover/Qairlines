package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/news"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type NewsHandler struct {
	getAllNewsWithAuthor news.IGetAllNewsWithAuthor
}

func NewNewsHandler(getAllNewsWithAuthor news.IGetAllNewsWithAuthor) *NewsHandler {
	return &NewsHandler{
		getAllNewsWithAuthor: getAllNewsWithAuthor,
	}
}

func (h *NewsHandler) GetAllNews(w http.ResponseWriter, r *http.Request) {
	news, err := h.getAllNewsWithAuthor.Execute(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to get news", err)
		return
	}

	response := mappers.NewsListToResponse(news)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to encode response", err)
		return
	}
}
