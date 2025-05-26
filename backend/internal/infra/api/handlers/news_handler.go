package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/news"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type NewsHandler struct {
	getAllNewsWithAuthor news.IGetAllNewsWithAuthor
	deleteNewsUseCase    news.IDeleteNewsUseCase
	createNewsUseCase    news.ICreateNewsUseCase
}

func NewNewsHandler(getAllNewsWithAuthor news.IGetAllNewsWithAuthor, deleteNewsUseCase news.IDeleteNewsUseCase, createNewsUseCase news.ICreateNewsUseCase) *NewsHandler {
	return &NewsHandler{
		getAllNewsWithAuthor: getAllNewsWithAuthor,
		deleteNewsUseCase:    deleteNewsUseCase,
		createNewsUseCase:    createNewsUseCase,
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

func (h *NewsHandler) DeleteNews(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Authentication failed. Admin privileges required.",
		})
		return
	}

	// Lấy ID từ query parameter
	newsIDStr := r.URL.Query().Get("id")
	if newsIDStr == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "News ID is required.",
		})
		return
	}

	newsID, err := strconv.ParseInt(newsIDStr, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid News ID.",
		})
		return
	}

	err = h.deleteNewsUseCase.Execute(r.Context(), newsID)
	if err != nil {
		if err == adapters.ErrNewsNotFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "News post not found.",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "An unexpected error occurred. Please try again later.",
		})
		return
	}

	// Trả về phản hồi thành công
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "News post deleted successfully.",
	})
}

func (h *NewsHandler) CreateNews(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Authentication failed. Admin privileges required.",
		})
		return
	}

	// Parse request body
	var req dto.CreateNewsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid news data. Please check the input fields.",
		})
		return
	}

	// Gọi use case để tạo bài viết
	new, err := h.createNewsUseCase.Execute(r.Context(), req)
	if err != nil {
		if errors.Is(err, news.ErrInvalidNewsData) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Invalid news data. Please check the input fields.",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "An unexpected error occurred. Please try again later.",
		})
		return
	}

	// Trả về phản hồi thành công
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "News post created successfully.",
		"data":    new,
	})
}
