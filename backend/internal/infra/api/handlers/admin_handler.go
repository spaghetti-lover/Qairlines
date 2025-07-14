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
	ListAdminsUseCase      admin.IListAdminsUseCase
	getCurrentAdminUseCase admin.IGetCurrentAdminUseCase
	updateAdminUseCase     admin.IUpdateAdminUseCase
	deleteAdminUseCase     admin.IDeleteAdminUseCase
}

func NewAdminHandler(
	adminCreateUseCase admin.ICreateAdminUseCase,
	getCurrentAdminUseCase admin.IGetCurrentAdminUseCase,
	ListAdminsUseCase admin.IListAdminsUseCase,
	updateAdminUseCase admin.IUpdateAdminUseCase,
	deleteAdminUseCase admin.IDeleteAdminUseCase,
) *AdminHandler {
	return &AdminHandler{
		adminCreateUseCase:     adminCreateUseCase,
		getCurrentAdminUseCase: getCurrentAdminUseCase,
		ListAdminsUseCase:      ListAdminsUseCase,
		updateAdminUseCase:     updateAdminUseCase,
		deleteAdminUseCase:     deleteAdminUseCase,
	}
}

func (h *AdminHandler) CreateAdminTx(ctx *gin.Context) {
	// Kiểm tra header "admin"
	if ctx.GetHeader("admin") != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	// Decode request body
	var createAdminRequest dto.CreateAdminRequest
	if err := ctx.ShouldBindJSON(&createAdminRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin data. Please check the input fields."})
		return
	}

	// Validate required fields
	if createAdminRequest.FirstName == "" || createAdminRequest.LastName == "" || createAdminRequest.Email == "" || createAdminRequest.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "All fields are required."})
		return
	}

	hashedPassword, err := utils.HashPassword(createAdminRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to hash password, %v", err.Error())})
		return
	}

	// Execute use case to create admin
	createdAdmin, err := h.adminCreateUseCase.Execute(ctx.Request.Context(), entities.CreateUserParams{
		FirstName: createAdminRequest.FirstName,
		LastName:  createAdminRequest.LastName,
		Email:     createAdminRequest.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Map entity to response DTO
	response := mappers.CreateAdminGetOutputToResponse(createdAdmin)

	// Write response
	ctx.JSON(http.StatusCreated, response)
}

func (h *AdminHandler) ListAdmins(ctx *gin.Context) {
	// Kiểm tra quyền admin
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	var params dto.ListAdminsParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can not bind query param"})
		return
	}

	if params.Page <= 0 {
		params.Page = 1
	}

	if params.Limit <= 0 {
		params.Limit = 10
	}

	if params.Limit > 100 {
		params.Limit = 100
	}

	// Lấy danh sách admin
	admins, err := h.ListAdminsUseCase.Execute(ctx.Request.Context(), params.Page, params.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	// Tạo response
	response := dto.ListAdminsResponse{
		Message: "Admins retrieved successfully.",
		Data:    mappers.AdminsEntitiesToResponse(admins),
	}

	// Trả về response
	ctx.JSON(http.StatusOK, response)
}

func (h *AdminHandler) GetCurrentAdmin(ctx *gin.Context) {
	// Kiểm tra quyền admin
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required.", "error": "Admin privileges required."})
		return
	}

	// Lấy payload từ context với key đúng
	authPayload, ok := ctx.Request.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Invalid token."})
		return
	}

	// Gọi use case để lấy thông tin admin hiện tại
	currentAdmin, err := h.getCurrentAdminUseCase.Execute(ctx.Request.Context(), authPayload.UserId)
	if err != nil {
		if errors.Is(err, adapters.ErrAdminNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Admin not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	// Tạo response
	response := mappers.CurrentAdminEntityToResponse(currentAdmin)

	// Trả về response
	ctx.JSON(http.StatusOK, response)
}

func (h *AdminHandler) UpdateAdmin(ctx *gin.Context) {
	// Kiểm tra quyền admin
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	// Lấy payload từ context với key đúng
	authPayload, ok := ctx.Request.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Invalid token."})
		return
	}

	userID := authPayload.UserId

	// Parse request body
	var updateRequest dto.AdminUpdateRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin data. Please check the input fields."})
		return
	}

	// Validate request
	if updateRequest.FirstName == "" || updateRequest.LastName == "" || updateRequest.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin data. Please check the input fields."})
		return
	}

	// Call use case
	updateInput := mappers.AdminUpdateRequestToInput(updateRequest, userID)
	updatedAdmin, err := h.updateAdminUseCase.Execute(ctx.Request.Context(), updateInput)
	if err != nil {
		if errors.Is(err, adapters.ErrAdminNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Admin not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	// Create response
	response := mappers.AdminUpdateEntityToResponse(updatedAdmin)

	// Return response
	ctx.JSON(http.StatusOK, response)
}

func (h *AdminHandler) DeleteAdmin(ctx *gin.Context) {
	// Lấy payload từ context với key đúng
	authPayload, ok := ctx.Request.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Invalid token."})
		return
	}

	// Gọi use case để xóa admin
	err := h.deleteAdminUseCase.Execute(ctx.Request.Context(), authPayload.UserId)
	if err != nil {
		if errors.Is(err, adapters.ErrAdminNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Admin not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	// Trả về thành công
	ctx.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully."})
}
