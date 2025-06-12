package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/booking"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type BookingHandler struct {
	createBookingUseCase booking.ICreateBookingUseCase
	tokenMaker           token.Maker
	userRepository       adapters.IUserRepository
	getBookingUseCase    booking.IGetBookingUseCase
}

func NewBookingHandler(createBookingUseCase booking.ICreateBookingUseCase, tokenMaker token.Maker, userRepository adapters.IUserRepository, getBookingUseCase booking.IGetBookingUseCase) *BookingHandler {
	return &BookingHandler{
		createBookingUseCase: createBookingUseCase,
		tokenMaker:           tokenMaker,
		userRepository:       userRepository,
		getBookingUseCase:    getBookingUseCase,
	}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
		return
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header format"})
		return
	}

	tokenStr := authHeader[len(bearerPrefix):]
	payload, err := h.tokenMaker.VerifyToken(tokenStr, token.TokenTypeAccessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("Unauthorized. %v", err.Error())})
		return
	}

	userId := payload.UserId
	// Lấy email từ UserId
	user, err := h.userRepository.GetUser(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to retrieve user email. %v", err.Error())})
		return
	}

	// Parse request body
	var request dto.CreateBookingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid booking data. Please check the input fields."})
		return
	}

	// Gọi use case để tạo booking
	bookingResponse, err := h.createBookingUseCase.Execute(c.Request.Context(), request, user.Email)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "One or more flights not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err.Error())})
		return
	}

	// Trả về response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking created successfully.",
		"data":    bookingResponse,
	})
}

func (h *BookingHandler) GetBooking(c *gin.Context) {
	bookingIDStr := c.Query("id")
	if bookingIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Booking ID is required."})
		return
	}

	bookingID, err := strconv.ParseInt(bookingIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid booking ID."})
		return
	}

	booking, departureTickets, returnTickets, err := h.getBookingUseCase.Execute(c.Request.Context(), bookingID)
	if err != nil {
		if errors.Is(err, adapters.ErrBookingNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Booking not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	response := mappers.ToGetBookingResponse(booking, departureTickets, returnTickets)

	c.JSON(http.StatusOK, gin.H{
		"message": "Booking details retrieved successfully.",
		"data":    response,
	})
}
