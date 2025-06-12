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

func (h *CustomerHandler) CreateCustomerTx(c *gin.Context) {
	var createCustomerRequest dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&createCustomerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid customer data. Please check the input fields.", "error": err.Error()})
		return
	}

	if createCustomerRequest.FirstName == "" || createCustomerRequest.LastName == "" || createCustomerRequest.Email == "" || createCustomerRequest.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "All fields are required."})
		return
	}

	customerParams := entities.CreateUserParams{
		FirstName: createCustomerRequest.FirstName,
		LastName:  createCustomerRequest.LastName,
		Email:     createCustomerRequest.Email,
		Password:  createCustomerRequest.Password,
	}

	createdCustomer, err := h.customerCreateUseCase.Execute(c.Request.Context(), customerParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Email được sử dụng hoặc mật khẩu không hợp lệ, %v", err.Error())})
		return
	}

	response := mappers.CreateCustomerGetOutputToResponse(createdCustomer)
	if response.User.FirstName == "" || response.User.LastName == "" || response.User.Email == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create customer."})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	var customerUpdateRequest dto.CustomerUpdateRequest
	if err := c.ShouldBindJSON(&customerUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid customer data. Please check the input fields."})
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID."})
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

	updatedCustomer, updatedUser, err := h.customerUpdateUseCase.Execute(c.Request.Context(), userID, customer, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Update customer failed: %v", err.Error())})
		return
	}

	response := mappers.CustomerUpdateResponse(updatedCustomer, updatedUser)
	c.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	customers, err := h.getAllCustomerUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customers retrieved successfully.",
		"data":    customers,
	})
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	customerIDStr := c.Query("id")
	if customerIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	err = h.deleteCustomerUseCase.Execute(c.Request.Context(), customerID)
	if err != nil {
		if errors.Is(err, adapters.ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Customer not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully."})
}

func (h *CustomerHandler) GetCustomerDetails(c *gin.Context) {
	payload, ok := c.Request.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || payload == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	customerDetails, err := h.getCustomerDetailsUseCase.Execute(c.Request.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, adapters.ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Customer not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customerDetails})
}
