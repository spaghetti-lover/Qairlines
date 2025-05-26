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
	updateNewsUseCase    news.IUpdateNewsUseCase
	getNewsUseCase       news.IGetNewsUseCase
}

func NewNewsHandler(getAllNewsWithAuthor news.IGetAllNewsWithAuthor, deleteNewsUseCase news.IDeleteNewsUseCase, createNewsUseCase news.ICreateNewsUseCase, updateNewsUseCase news.IUpdateNewsUseCase, getNewsUseCase news.IGetNewsUseCase) *NewsHandler {
	return &NewsHandler{
		getAllNewsWithAuthor: getAllNewsWithAuthor,
		deleteNewsUseCase:    deleteNewsUseCase,
		createNewsUseCase:    createNewsUseCase,
		updateNewsUseCase:    updateNewsUseCase,
		getNewsUseCase:       getNewsUseCase,
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

func (h *NewsHandler) UpdateNews(w http.ResponseWriter, r *http.Request) {
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

	// Parse form data
	err = r.ParseMultipartForm(10 << 20) // Giới hạn kích thước file upload (10MB)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid form data.",
		})
		return
	}

	// Lấy dữ liệu từ form
	req := dto.UpdateNewsRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Content:     r.FormValue("content"),
		AuthorID:    r.FormValue("authorId"),
	}

	// Lấy file hình ảnh (nếu có)
	file, _, err := r.FormFile("news-image")
	if err == nil {
		defer file.Close()
		// Xử lý upload file (giả sử bạn lưu file và trả về URL)
		req.Image = "https://example.com/path/to/updated-image.jpg" // Thay bằng URL thực tế
	}

	// Gọi use case để cập nhật bài viết
	updatedNews, err := h.updateNewsUseCase.Execute(r.Context(), newsID, req)
	if err != nil {
		if errors.Is(err, news.ErrNewsNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "News post not found.",
			})
			return
		}
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "News post updated successfully.",
		"data":    updatedNews,
	})
}

func (h *NewsHandler) GetNews(w http.ResponseWriter, r *http.Request) {
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

	// Gọi use case để lấy bài viết
	new, err := h.getNewsUseCase.Execute(r.Context(), newsID)
	if err != nil {
		if errors.Is(err, news.ErrNewsNotFound) {
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
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "News post retrieved successfully.",
		"data":    new,
	})
}
