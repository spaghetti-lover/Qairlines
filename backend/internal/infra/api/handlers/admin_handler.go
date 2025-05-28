package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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

func (h *AdminHandler) CreateAdminTx(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra header "admin"
	if r.Header.Get("admin") != "true" {
		http.Error(w, `{"message": "Authentication failed. Admin privileges required."}`, http.StatusUnauthorized)
		return
	}

	// Decode request body
	var createAdminRequest dto.CreateAdminRequest
	if err := json.NewDecoder(r.Body).Decode(&createAdminRequest); err != nil {
		http.Error(w, `{"message": "Invalid admin data. Please check the input fields."}`, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if createAdminRequest.FirstName == "" || createAdminRequest.LastName == "" || createAdminRequest.Email == "" || createAdminRequest.Password == "" {
		http.Error(w, `{"message": "All fields are required."}`, http.StatusBadRequest)
		return
	}
	hashedPassword, err := utils.HashPassword(createAdminRequest.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "Failed to hash password, %v"}`, err.Error()), http.StatusInternalServerError)
		return
	}
	// Execute use case to create admin
	createdAdmin, err := h.adminCreateUseCase.Execute(r.Context(), entities.CreateUserParams{
		FirstName: createAdminRequest.FirstName,
		LastName:  createAdminRequest.LastName,
		Email:     createAdminRequest.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "%v"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Map entity to response DTO
	response := mappers.CreateAdminGetOutputToResponse(createdAdmin)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *AdminHandler) GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		http.Error(w, `{"message": "Authentication failed. Admin privileges required."}`, http.StatusUnauthorized)
		return
	}

	// Lấy danh sách admin
	admins, err := h.getAllAdminsUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}

	// Tạo response
	response := dto.GetAllAdminsResponse{
		Message: "Admins retrieved successfully.",
		Data:    mappers.AdminsEntitiesToResponse(admins),
	}

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandler) GetCurrentAdmin(w http.ResponseWriter, r *http.Request) {
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		http.Error(w, `{"message": "Authentication failed. Admin privileges required."}`, http.StatusUnauthorized)
		return
	}

	// Lấy payload từ context với key đúng
	authPayload, ok := r.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		http.Error(w, `{"message": "Authentication failed. Invalid token."}`, http.StatusUnauthorized)
		return
	}

	currentAdmin, err := h.getCurrentAdminUseCase.Execute(r.Context(), authPayload.UserId)
	if err != nil {
		if errors.Is(err, admin.ErrAdminNotFound) {
			http.Error(w, `{"message": "Admin not found."}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}

	// Tạo response
	response := mappers.CurrentAdminEntityToResponse(currentAdmin)

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandler) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		http.Error(w, `{"message": "Authentication failed. Admin privileges required."}`, http.StatusUnauthorized)
		return
	}

	// Lấy payload từ context với key đúng
	authPayload, ok := r.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		http.Error(w, `{"message": "Authentication failed. Invalid token."}`, http.StatusUnauthorized)
		return
	}

	userID := authPayload.UserId

	// Parse request body
	var updateRequest dto.AdminUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, `{"message": "Invalid admin data. Please check the input fields."}`, http.StatusBadRequest)
		return
	}

	// Validate request
	if updateRequest.FirstName == "" || updateRequest.LastName == "" || updateRequest.Email == "" {
		http.Error(w, `{"message": "Invalid admin data. Please check the input fields."}`, http.StatusBadRequest)
		return
	}

	// Call use case
	updateInput := mappers.AdminUpdateRequestToInput(updateRequest, userID)
	updatedAdmin, err := h.updateAdminUseCase.Execute(r.Context(), updateInput)
	if err != nil {
		if errors.Is(err, admin.ErrAdminNotFound) {
			http.Error(w, `{"message": "Admin not found."}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}

	// Create response
	response := mappers.AdminUpdateEntityToResponse(updatedAdmin)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		http.Error(w, `{"message": "Authentication failed. Admin privileges required."}`, http.StatusUnauthorized)
		return
	}

	// Lấy payload từ context với key đúng
	authPayload, ok := r.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		http.Error(w, `{"message": "Authentication failed. Invalid token."}`, http.StatusUnauthorized)
		return
	}

	// Gọi use case để xóa admin
	err := h.deleteAdminUseCase.Execute(r.Context(), authPayload.UserId)
	if err != nil {
		if errors.Is(err, admin.ErrAdminNotFound) {
			http.Error(w, `{"message": "Admin not found."}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}

	// Trả về thành công (204 No Content hoặc 200 OK với message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Admin deleted successfully."}`))
}
