package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/config"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/news"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type NewsHandler struct {
	listNewsUseCase   news.IListNewsUseCase
	deleteNewsUseCase news.IDeleteNewsUseCase
	createNewsUseCase news.ICreateNewsUseCase
	updateNewsUseCase news.IUpdateNewsUseCase
	getNewsUseCase    news.IGetNewsUseCase
}

func NewNewsHandler(listNewsUseCase news.IListNewsUseCase, deleteNewsUseCase news.IDeleteNewsUseCase, createNewsUseCase news.ICreateNewsUseCase, updateNewsUseCase news.IUpdateNewsUseCase, getNewsUseCase news.IGetNewsUseCase) *NewsHandler {
	return &NewsHandler{
		listNewsUseCase:   listNewsUseCase,
		deleteNewsUseCase: deleteNewsUseCase,
		createNewsUseCase: createNewsUseCase,
		updateNewsUseCase: updateNewsUseCase,
		getNewsUseCase:    getNewsUseCase,
	}
}

func (h *NewsHandler) ListNews(ctx *gin.Context) {
	var params entities.ListNewsParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can not bind query param"})
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}
	news, err := h.listNewsUseCase.Execute(ctx.Request.Context(), params.Page, params.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get news", "error": err.Error()})
		return
	}

	response := mappers.NewsListToResponse(news)
	ctx.JSON(http.StatusOK, response)
}

func (h *NewsHandler) DeleteNews(ctx *gin.Context) {
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	newsIDStr := ctx.Query("id")
	if newsIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "News ID is required."})
		return
	}

	newsID, err := strconv.ParseInt(newsIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid News ID."})
		return
	}

	err = h.deleteNewsUseCase.Execute(ctx.Request.Context(), newsID)
	if err != nil {
		if err == adapters.ErrNewsNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "News post not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "News post deleted successfully."})
}

func (h *NewsHandler) CreateNews(ctx *gin.Context) {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	var publicURL = fmt.Sprintf("http://localhost%s/images/", config.ServerAddressPort)
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	var req dto.CreateNewsRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid news data. Please check the input fields." + err.Error()})
		return
	}

	image, err := ctx.FormFile("news-image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Image file is required."})
		return
	}

	// Yêu cầu giới hạn file nhỏ hơn 5MB
	// 1 << 20 = 1 * 2^20 = 1 * 1048576 = 1MB
	// 5 << 20 = 5 * 2^20 = 5 * 1048576 = 5MB
	if image.Size > 5<<20 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File too large (5 MB)"})
		return
	}

	// Tạo thư mục uploads nếu chưa tồn tại
	uploadsDir := "./uploads"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		err = os.MkdirAll(uploadsDir, os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create uploads directory."})
			return
		}
	}

	// Lưu file vào thư mục uploads
	filename, err := utils.ValidateAndSaveFile(image, "./uploads")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}
	dst := filepath.Join(uploadsDir, filename)
	if err := ctx.SaveUploadedFile(image, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image file."})
		return
	}

	publicImageURL := publicURL + filename

	createNews := dto.CreateNewsToDBRequest{Title: req.Title, Description: req.Description, Content: req.Content, Image: publicImageURL, AuthorID: req.AuthorID}
	news, err := h.createNewsUseCase.Execute(ctx, createNews)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create news." + err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "News post created successfully.",
		"data":    news,
	})
}

// func (h *NewsHandler) UpdateNews(ctx *gin.Context) {
// 	isAdmin := ctx.GetHeader("admin")
// 	if isAdmin != "true" {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
// 		return
// 	}

// 	newsIDStr := ctx.Query("id")
// 	if newsIDStr == "" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "News ID is required."})
// 		return
// 	}

// 	newsID, err := strconv.ParseInt(newsIDStr, 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid News ID."})
// 		return
// 	}

// 	err = ctx.Request.ParseMultipartForm(10 << 20) // Giới hạn kích thước file upload (10MB)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid form data."})
// 		return
// 	}

// 	req := dto.UpdateNewsRequest{
// 		Title:       c.PostForm("title"),
// 		Description: c.PostForm("description"),
// 		Content:     c.PostForm(5 <<"content"),
// 		AuthorID:    c.PostForm("authorId"),
// 	}

// 	file, _, err := ctx.FormFile("news-image")
// 	if err == nil {
// 		defer file.Close()
// 		req.Image = "https://example.com/path/to/updated-image.jpg" // Thay bằng URL thực tế
// 	}

// 	updatedNews, err := h.updateNewsUseCase.Execute(ctx.Request.Context(), newsID, req)
// 	if err != nil {
// 		if errors.Is(err, news.ErrNewsNotFound) {
// 			ctx.JSON(http.StatusNotFound, gin.H{"message": "News post not found."})
// 			return
// 		}
// 		if errors.Is(err, news.ErrInvalidNewsData) {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid news data. Please check the input fields."})
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "News post updated successfully.", "data": updatedNews})
// }

func (h *NewsHandler) GetNews(ctx *gin.Context) {
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	newsIDStr := ctx.Param("id")
	if newsIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "News ID is required."})
		return
	}

	newsID, err := strconv.ParseInt(newsIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid News ID."})
		return
	}

	new, err := h.getNewsUseCase.Execute(ctx.Request.Context(), newsID)
	if err != nil {
		if errors.Is(err, news.ErrNewsNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "News post not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "News post retrieved successfully.", "data": new})
}
