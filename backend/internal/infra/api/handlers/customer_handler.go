package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/customer"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/user"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/middleware"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type CustomerHandler struct {
	customerCreateUseCase     customer.ICreateCustomerUseCase
	customerUpdateUseCase     customer.ICustomerUpdateUseCase
	userUpdateUseCase         user.IUserUpdateUseCase
	getAllCustomerUseCase     customer.IGetAllCustomersUseCase
	deleteCustomerUseCase     customer.IDeleteCustomerUseCase
	getCustomerDetailsUseCase customer.IGetCustomerDetailsUseCase
}

func NewCustomerHandler(customerCreateUseCase customer.ICreateCustomerUseCase, customerUpdateUseCase customer.ICustomerUpdateUseCase, userUpdateUseCase user.IUserUpdateUseCase, getAllCustomerUseCase customer.IGetAllCustomersUseCase, deleteCustomerUseCase customer.IDeleteCustomerUseCase, getCustomerDetailsUseCase customer.IGetCustomerDetailsUseCase) *CustomerHandler {
	return &CustomerHandler{
		customerCreateUseCase:     customerCreateUseCase,
		customerUpdateUseCase:     customerUpdateUseCase,
		userUpdateUseCase:         userUpdateUseCase,
		getAllCustomerUseCase:     getAllCustomerUseCase,
		deleteCustomerUseCase:     deleteCustomerUseCase,
		getCustomerDetailsUseCase: getCustomerDetailsUseCase,
	}
}

func (h *CustomerHandler) CreateCustomerTx(ctx *gin.Context) {
	var createCustomerRequest dto.CreateCustomerRequest
	if err := ctx.ShouldBindJSON(&createCustomerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid customer data. Please check the input fields.", "error": err.Error()})
		return
	}

	if createCustomerRequest.FirstName == "" || createCustomerRequest.LastName == "" || createCustomerRequest.Email == "" || createCustomerRequest.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "All fields are required."})
		return
	}

	customerParams := entities.CreateUserParams{
		FirstName: createCustomerRequest.FirstName,
		LastName:  createCustomerRequest.LastName,
		Email:     createCustomerRequest.Email,
		Password:  createCustomerRequest.Password,
	}

	createdCustomer, err := h.customerCreateUseCase.Execute(ctx.Request.Context(), customerParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Email được sử dụng hoặc mật khẩu không hợp lệ, %v", err.Error())})
		return
	}

	response := mappers.CreateCustomerGetOutputToResponse(createdCustomer)
	if response.User.FirstName == "" || response.User.LastName == "" || response.User.Email == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create customer."})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *CustomerHandler) UpdateCustomer(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	var customerUpdateRequest dto.CustomerUpdateRequest
	if err := ctx.ShouldBindJSON(&customerUpdateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid customer data. Please check the input fields."})
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID."})
		return
	}

	customer := entities.Customer{
		UserID:               userID,
		PhoneNumber:          customerUpdateRequest.PhoneNumber,
		Gender:               entities.CustomerGender(customerUpdateRequest.Gender),
		Address:              customerUpdateRequest.Address,
		DateOfBirth:          time.Unix(customerUpdateRequest.DateOfBirth.Seconds, 0).UTC(),
		PassportNumber:       customerUpdateRequest.PassportNumber,
		IdentificationNumber: customerUpdateRequest.IdentificationNumber,
	}

	user := entities.User{
		UserID:    userID,
		FirstName: customerUpdateRequest.FirstName,
		LastName:  customerUpdateRequest.LastName,
	}

	updatedCustomer, updatedUser, err := h.customerUpdateUseCase.Execute(ctx.Request.Context(), userID, customer, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Update customer failed: %v", err.Error())})
		return
	}

	response := mappers.CustomerUpdateResponse(updatedCustomer, updatedUser)
	ctx.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) GetAllCustomers(ctx *gin.Context) {
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	customers, err := h.getAllCustomerUseCase.Execute(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Customers retrieved successfully.",
		"data":    customers,
	})
}

func (h *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	customerIDStr := ctx.Query("id")
	if customerIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	err = h.deleteCustomerUseCase.Execute(ctx.Request.Context(), customerID)
	if err != nil {
		if errors.Is(err, adapters.ErrCustomerNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Customer not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully."})
}

func (h *CustomerHandler) GetCustomerDetails(ctx *gin.Context) {
	payload, ok := ctx.Request.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || payload == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	customerDetails, err := h.getCustomerDetailsUseCase.Execute(ctx.Request.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, adapters.ErrCustomerNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Customer not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": customerDetails})
}
