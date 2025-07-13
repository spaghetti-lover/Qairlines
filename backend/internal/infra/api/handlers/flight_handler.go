package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/flight"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type FlightHandler struct {
	createFlightUseCase      flight.ICreateFlightUseCase
	getFlightUseCase         flight.IGetFlightUseCase
	updateFlightTimesUseCase flight.IUpdateFlightTimesUseCase
	getAllFlightsUseCase     flight.IGetAllFlightsUseCase
	deleteFlightUseCase      flight.IDeleteFlightUseCase
	searchFlightsUseCase     flight.ISearchFlightsUseCase
	listFlightsUseCase       flight.IListFlightsUseCase
}

func NewFlightHandler(createFlightUseCase flight.ICreateFlightUseCase, getFlightUseCase flight.IGetFlightUseCase, updateFlightTimesUseCase flight.IUpdateFlightTimesUseCase, getAllFlightsUseCase flight.IGetAllFlightsUseCase, deleteFlightUseCase flight.IDeleteFlightUseCase, searchFlightsUseCase flight.ISearchFlightsUseCase, listFlightsUseCase flight.IListFlightsUseCase) *FlightHandler {
	return &FlightHandler{createFlightUseCase: createFlightUseCase, getFlightUseCase: getFlightUseCase, updateFlightTimesUseCase: updateFlightTimesUseCase,
		getAllFlightsUseCase: getAllFlightsUseCase, deleteFlightUseCase: deleteFlightUseCase,
		searchFlightsUseCase: searchFlightsUseCase,
		listFlightsUseCase:   listFlightsUseCase,
	}
}

func (h *FlightHandler) CreateFlight(ctx *gin.Context) {
	var req dto.CreateFlightRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid request body, %v", err)})
		return
	}

	flightEntity := mappers.CreateFlightRequestToEntity(req)
	createdFlight, err := h.createFlightUseCase.Execute(ctx.Request.Context(), flightEntity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to create flight, %v", err)})
		return
	}

	response := mappers.CreateFlightEntityToResponse(createdFlight)
	ctx.JSON(http.StatusCreated, response)
}

func (h *FlightHandler) GetFlight(ctx *gin.Context) {
	flightIDStr := ctx.Param("id")
	if flightIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Flight ID is required."})
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Flight ID."})
		return
	}

	flightDetails, err := h.getFlightUseCase.Execute(ctx.Request.Context(), flightID)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Flight not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Flight details retrieved successfully.",
		"data":    flightDetails,
	})
}

func (h *FlightHandler) UpdateFlightTimes(ctx *gin.Context) {
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	flightIDStr := ctx.Query("id")
	if flightIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Flight ID is required."})
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Flight ID."})
		return
	}

	var request dto.UpdateFlightTimesRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid flight data. Please check the input fields."})
		return
	}

	response, err := h.updateFlightTimesUseCase.Execute(ctx.Request.Context(), flightID, request)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Flight not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Flight updated successfully.",
		"flight":  response,
	})
}

func (h *FlightHandler) GetAllFlights(ctx *gin.Context) {
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	flights, tickets, err := h.getAllFlightsUseCase.Execute(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	response := mappers.MapFlightsAndTicketsToResponse(flights, tickets)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Flights retrieved successfully.",
		"data":    response,
	})
}

func (h *FlightHandler) DeleteFlight(ctx *gin.Context) {
	isAdmin := ctx.GetHeader("admin")
	if isAdmin != "true" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	flightIDStr := ctx.Query("id")
	if flightIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Flight ID is required."})
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Flight ID."})
		return
	}

	err = h.deleteFlightUseCase.Execute(ctx.Request.Context(), flightID)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Flight not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully."})
}

func (h *FlightHandler) SearchFlights(ctx *gin.Context) {
	departureCity := ctx.Query("departureCity")
	arrivalCity := ctx.Query("arrivalCity")
	flightDate := ctx.Query("flightDate")

	if departureCity == "" || arrivalCity == "" || flightDate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters. Please check departureCity, arrivalCity, and flightDate."})
		return
	}

	flights, err := h.searchFlightsUseCase.Execute(ctx.Request.Context(), departureCity, arrivalCity, flightDate)
	if err != nil {
		if errors.Is(err, adapters.ErrNoFlightsFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "No flights found for the given criteria."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Flights retrieved successfully.",
		"data":    flights,
	})
}

func (h *FlightHandler) ListFlights(ctx *gin.Context) {
	var params entities.ListFlightsParams
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

	flights, err := h.listFlightsUseCase.Execute(ctx.Request.Context(), params.Page, params.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Suggested flights retrieved successfully.",
		"data":    flights,
	})
}
