package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/news"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
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

func (h *NewsHandler) GetAllNews(c *gin.Context) {
	news, err := h.getAllNewsWithAuthor.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get news", "error": err.Error()})
		return
	}

	response := mappers.NewsListToResponse(news)
	c.JSON(http.StatusOK, response)
}

func (h *NewsHandler) DeleteNews(c *gin.Context) {
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	newsIDStr := c.Query("id")
	if newsIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "News ID is required."})
		return
	}

	newsID, err := strconv.ParseInt(newsIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid News ID."})
		return
	}

	err = h.deleteNewsUseCase.Execute(c.Request.Context(), newsID)
	if err != nil {
		if err == adapters.ErrNewsNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "News post not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News post deleted successfully."})
}

func (h *NewsHandler) CreateNews(c *gin.Context) {
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	var req dto.CreateNewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid news data. Please check the input fields."})
		return
	}

	new, err := h.createNewsUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, news.ErrInvalidNewsData) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid news data. Please check the input fields."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "News post created successfully.", "data": new})
}

// func (h *NewsHandler) UpdateNews(c *gin.Context) {
// 	isAdmin := c.GetHeader("admin")
// 	if isAdmin != "true" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
// 		return
// 	}

// 	newsIDStr := c.Query("id")
// 	if newsIDStr == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "News ID is required."})
// 		return
// 	}

// 	newsID, err := strconv.ParseInt(newsIDStr, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid News ID."})
// 		return
// 	}

// 	err = c.Request.ParseMultipartForm(10 << 20) // Giới hạn kích thước file upload (10MB)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid form data."})
// 		return
// 	}

// 	req := dto.UpdateNewsRequest{
// 		Title:       c.PostForm("title"),
// 		Description: c.PostForm("description"),
// 		Content:     c.PostForm("content"),
// 		AuthorID:    c.PostForm("authorId"),
// 	}

// 	file, _, err := c.FormFile("news-image")
// 	if err == nil {
// 		defer file.Close()
// 		req.Image = "https://example.com/path/to/updated-image.jpg" // Thay bằng URL thực tế
// 	}

// 	updatedNews, err := h.updateNewsUseCase.Execute(c.Request.Context(), newsID, req)
// 	if err != nil {
// 		if errors.Is(err, news.ErrNewsNotFound) {
// 			c.JSON(http.StatusNotFound, gin.H{"message": "News post not found."})
// 			return
// 		}
// 		if errors.Is(err, news.ErrInvalidNewsData) {
// 			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid news data. Please check the input fields."})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "News post updated successfully.", "data": updatedNews})
// }

func (h *NewsHandler) GetNews(c *gin.Context) {
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	newsIDStr := c.Query("id")
	if newsIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "News ID is required."})
		return
	}

	newsID, err := strconv.ParseInt(newsIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid News ID."})
		return
	}

	new, err := h.getNewsUseCase.Execute(c.Request.Context(), newsID)
	if err != nil {
		if errors.Is(err, news.ErrNewsNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "News post not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News post retrieved successfully.", "data": new})
}
