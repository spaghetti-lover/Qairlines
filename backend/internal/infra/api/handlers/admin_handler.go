package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/admin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/middleware"
	"github.com/spaghetti-lover/qairlines/pkg/token"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type AdminHandler struct {
	adminCreateUseCase     admin.ICreateAdminUseCase
	getAllAdminsUseCase    admin.IGetAllAdminsUseCase
	getCurrentAdminUseCase admin.IGetCurrentAdminUseCase
	updateAdminUseCase     admin.IUpdateAdminUseCase
	deleteAdminUseCase     admin.IDeleteAdminUseCase
}

func NewAdminHandler(
	adminCreateUseCase admin.ICreateAdminUseCase,
	getCurrentAdminUseCase admin.IGetCurrentAdminUseCase,
	getAllAdminsUseCase admin.IGetAllAdminsUseCase,
	updateAdminUseCase admin.IUpdateAdminUseCase,
	deleteAdminUseCase admin.IDeleteAdminUseCase,
) *AdminHandler {
	return &AdminHandler{
		adminCreateUseCase:     adminCreateUseCase,
		getCurrentAdminUseCase: getCurrentAdminUseCase,
		getAllAdminsUseCase:    getAllAdminsUseCase,
		updateAdminUseCase:     updateAdminUseCase,
		deleteAdminUseCase:     deleteAdminUseCase,
	}
}

func (h *AdminHandler) CreateAdminTx(c *gin.Context) {
	// Kiểm tra header "admin"
	if c.GetHeader("admin") != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	// Decode request body
	var createAdminRequest dto.CreateAdminRequest
	if err := c.ShouldBindJSON(&createAdminRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin data. Please check the input fields."})
		return
	}

	// Validate required fields
	if createAdminRequest.FirstName == "" || createAdminRequest.LastName == "" || createAdminRequest.Email == "" || createAdminRequest.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "All fields are required."})
		return
	}

	hashedPassword, err := utils.HashPassword(createAdminRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to hash password, %v", err.Error())})
		return
	}

	// Execute use case to create admin
	createdAdmin, err := h.adminCreateUseCase.Execute(c.Request.Context(), entities.CreateUserParams{
		FirstName: createAdminRequest.FirstName,
		LastName:  createAdminRequest.LastName,
		Email:     createAdminRequest.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Map entity to response DTO
	response := mappers.CreateAdminGetOutputToResponse(createdAdmin)

	// Write response
	c.JSON(http.StatusCreated, response)
}

func (h *AdminHandler) GetAllAdmins(c *gin.Context) {
	// Kiểm tra quyền admin
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	// Lấy danh sách admin
	admins, err := h.getAllAdminsUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	// Tạo response
	response := dto.GetAllAdminsResponse{
		Message: "Admins retrieved successfully.",
		Data:    mappers.AdminsEntitiesToResponse(admins),
	}

	// Trả về response
	c.JSON(http.StatusOK, response)
}

func (h *AdminHandler) GetCurrentAdmin(c *gin.Context) {
	// Kiểm tra quyền admin
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	// Lấy payload từ context với key đúng
	authPayload, ok := c.Request.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Invalid token."})
		return
	}

	// Gọi use case để lấy thông tin admin hiện tại
	currentAdmin, err := h.getCurrentAdminUseCase.Execute(c.Request.Context(), authPayload.UserId)
	if err != nil {
		if errors.Is(err, adapters.ErrAdminNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Admin not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	// Tạo response
	response := mappers.CurrentAdminEntityToResponse(currentAdmin)

	// Trả về response
	c.JSON(http.StatusOK, response)
}

func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	// Kiểm tra quyền admin
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	// Lấy payload từ context với key đúng
	authPayload, ok := c.Request.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Invalid token."})
		return
	}

	userID := authPayload.UserId

	// Parse request body
	var updateRequest dto.AdminUpdateRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin data. Please check the input fields."})
		return
	}

	// Validate request
	if updateRequest.FirstName == "" || updateRequest.LastName == "" || updateRequest.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin data. Please check the input fields."})
		return
	}

	// Call use case
	updateInput := mappers.AdminUpdateRequestToInput(updateRequest, userID)
	updatedAdmin, err := h.updateAdminUseCase.Execute(c.Request.Context(), updateInput)
	if err != nil {
		if errors.Is(err, adapters.ErrAdminNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Admin not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	// Create response
	response := mappers.AdminUpdateEntityToResponse(updatedAdmin)

	// Return response
	c.JSON(http.StatusOK, response)
}

func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	// Lấy payload từ context với key đúng
	authPayload, ok := c.Request.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Invalid token."})
		return
	}

	// Gọi use case để xóa admin
	err := h.deleteAdminUseCase.Execute(c.Request.Context(), authPayload.UserId)
	if err != nil {
		if errors.Is(err, adapters.ErrAdminNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Admin not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	// Trả về thành công
	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully."})
}
