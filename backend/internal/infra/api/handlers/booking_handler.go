package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/booking"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type BookingHandler struct {
	createBookingUseCase booking.ICreateBookingUseCase
	tokenMaker           token.Maker
	userRepository       adapters.IUserRepository
}

func NewBookingHandler(createBookingUseCase booking.ICreateBookingUseCase, tokenMaker token.Maker, userRepository adapters.IUserRepository) *BookingHandler {
	return &BookingHandler{
		createBookingUseCase: createBookingUseCase,
		tokenMaker:           tokenMaker,
		userRepository:       userRepository,
	}
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		fmt.Println("Authorization header is missing")
		http.Error(w, `{"message": "Authorization header is missing"}`, http.StatusUnauthorized)
		return
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		fmt.Println("Invalid authorization header format")
		http.Error(w, `{"message": "Invalid authorization header format"}`, http.StatusUnauthorized)
		return
	}

	tokenStr := authHeader[len(bearerPrefix):]
	payload, err := h.tokenMaker.VerifyToken(tokenStr, token.TokenTypeAccessToken)
	if err != nil {
		fmt.Printf("Token verification failed: %v\n", err)
		http.Error(w, fmt.Sprintf(`{"message": "Unauthorized. %v"}`, err.Error()), http.StatusUnauthorized)
		return
	}

	userId := payload.UserId
	// Lấy email từ UserId
	user, err := h.userRepository.GetUser(r.Context(), userId)

	if err != nil {
		http.Error(w, `{"message": "Failed to retrieve user email."}`, http.StatusInternalServerError)
		return
	}

	// Parse request body
	var request dto.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"message": "Invalid booking data. Please check the input fields."}`, http.StatusBadRequest)
		return
	}

	// Gọi use case để tạo booking
	bookingResponse, err := h.createBookingUseCase.Execute(r.Context(), request, user.Email)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			http.Error(w, `{"message": "One or more flights not found."}`, http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf(`{"message": "An unexpected error occurred. %v"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Booking created successfully.",
		"data":    bookingResponse,
	})
}
