package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/customer"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/user"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type CustomerHandler struct {
	customerCreateUseCase customer.ICreateCustomerUseCase
	customerUpdateUseCase customer.ICustomerUpdateUseCase
	userUpdateUseCase     user.IUserUpdateUseCase
	getAllCustomerUseCase customer.IGetAllCustomersUseCase
}

func NewCustomerHandler(customerCreateUseCase customer.ICreateCustomerUseCase, customerUpdateUseCase customer.ICustomerUpdateUseCase, userUpdateUseCase user.IUserUpdateUseCase, getAllCustomerUseCase customer.IGetAllCustomersUseCase) *CustomerHandler {
	return &CustomerHandler{
		customerCreateUseCase: customerCreateUseCase,
		customerUpdateUseCase: customerUpdateUseCase,
		userUpdateUseCase:     userUpdateUseCase,
		getAllCustomerUseCase: getAllCustomerUseCase,
	}
}

func (h *CustomerHandler) CreateCustomerTx(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var createCustomerRequest dto.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&createCustomerRequest); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid customer data. Please check the input fields.", err)
		return
	}

	// Validate required fields
	if createCustomerRequest.FirstName == "" || createCustomerRequest.LastName == "" || createCustomerRequest.Email == "" || createCustomerRequest.Password == "" {
		utils.WriteError(w, http.StatusBadRequest, "All fields are required.", nil)
		return
	}

	// Map DTO to entity
	customerParams := entities.CreateUserParams{
		FirstName: createCustomerRequest.FirstName,
		LastName:  createCustomerRequest.LastName,
		Email:     createCustomerRequest.Email,
		Password:  createCustomerRequest.Password,
	}

	// Execute use case to create customer
	createdCustomer, err := h.customerCreateUseCase.Execute(r.Context(), customerParams)

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "Email được sử dụng hoặc mật khẩu không hợp lệ, %v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	// Map entity to response DTO
	response := mappers.CreateCustomerGetOutputToResponse(createdCustomer)
	if response.User.FirstName == "" || response.User.LastName == "" || response.User.Email == "" {
		http.Error(w, `{"message": "Failed to create customer."}`, http.StatusInternalServerError)
		return
	}
	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	if idParam == "" {
		utils.WriteError(w, http.StatusBadRequest, "customer ID is required", nil)
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid customer ID", err)
		return
	}
	var customerUpdateRequest dto.CustomerUpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&customerUpdateRequest); err != nil {
		http.Error(w, `{"message": "Invalid customer data. Please check the input fields."}`, http.StatusBadRequest)
		return
	}
	customer := entities.Customer{
		UserID:               id,
		PhoneNumber:          customerUpdateRequest.PhoneNumber,
		Gender:               entities.CustomerGender(customerUpdateRequest.Gender),
		Address:              customerUpdateRequest.Address,
		DateOfBirth:          time.Unix(customerUpdateRequest.DateOfBirth.Seconds, 0).UTC(),
		PassportNumber:       customerUpdateRequest.PassportNumber,
		IdentificationNumber: customerUpdateRequest.IdentificationNumber,
	}

	user := entities.User{
		UserID:    id,
		FirstName: customerUpdateRequest.FirstName,
		LastName:  customerUpdateRequest.LastName,
	}

	updatedCustomer, updatedUser, err := h.customerUpdateUseCase.Execute(r.Context(), id, customer, user)

	if err != nil {
		http.Error(w, `{"message": "Customer not found."}`, http.StatusInternalServerError)
		return
	}

	response := mappers.CustomerUpdateResponse(updatedCustomer, updatedUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		http.Error(w, "Authentication failed. Admin privileges required.", http.StatusUnauthorized)
		return
	}

	// Gọi use case để lấy danh sách khách hàng
	customers, err := h.getAllCustomerUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Customers retrieved successfully.",
		"data":    customers,
	})
}
