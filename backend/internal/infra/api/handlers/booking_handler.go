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

func (h *BookingHandler) CreateBooking(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
		return
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header format"})
		return
	}

	tokenStr := authHeader[len(bearerPrefix):]
	payload, err := h.tokenMaker.VerifyToken(tokenStr, token.TokenTypeAccessToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("Unauthorized. %v", err.Error())})
		return
	}

	userId := payload.UserId
	// Lấy email từ UserId
	user, err := h.userRepository.GetUser(ctx.Request.Context(), userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to retrieve user email. %v", err.Error())})
		return
	}

	// Parse request body
	var request dto.CreateBookingRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid booking data. Please check the input fields."})
		return
	}

	// Gọi use case để tạo booking
	bookingResponse, err := h.createBookingUseCase.Execute(ctx.Request.Context(), request, user.Email)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "One or more flights not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err.Error())})
		return
	}

	// Trả về response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Booking created successfully.",
		"data":    bookingResponse,
	})
}

func (h *BookingHandler) GetBooking(ctx *gin.Context) {
	bookingIDStr := ctx.Query("id")
	if bookingIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Booking ID is required."})
		return
	}

	bookingID, err := strconv.ParseInt(bookingIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid booking ID."})
		return
	}

	booking, departureTickets, returnTickets, err := h.getBookingUseCase.Execute(ctx.Request.Context(), bookingID)
	if err != nil {
		if errors.Is(err, adapters.ErrBookingNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Booking not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	response := mappers.ToGetBookingResponse(booking, departureTickets, returnTickets)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Booking details retrieved successfully.",
		"data":    response,
	})
}
