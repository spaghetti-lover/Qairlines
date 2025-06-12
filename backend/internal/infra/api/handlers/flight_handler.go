package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/flight"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type FlightHandler struct {
	createFlightUseCase        flight.ICreateFlightUseCase
	getFlightUseCase           flight.IGetFlightUseCase
	updateFlightTimesUseCase   flight.IUpdateFlightTimesUseCase
	getAllFlightsUseCase       flight.IGetAllFlightsUseCase
	deleteFlightUseCase        flight.IDeleteFlightUseCase
	searchFlightsUseCase       flight.ISearchFlightsUseCase
	getSuggestedFlightsUseCase flight.IGetSuggestedFlightsUseCase
}

func NewFlightHandler(createFlightUseCase flight.ICreateFlightUseCase, getFlightUseCase flight.IGetFlightUseCase, updateFlightTimesUseCase flight.IUpdateFlightTimesUseCase, getAllFlightsUseCase flight.IGetAllFlightsUseCase, deleteFlightUseCase flight.IDeleteFlightUseCase, searchFlightsUseCase flight.ISearchFlightsUseCase, getSuggestedFlightsUseCase flight.IGetSuggestedFlightsUseCase) *FlightHandler {
	return &FlightHandler{createFlightUseCase: createFlightUseCase, getFlightUseCase: getFlightUseCase, updateFlightTimesUseCase: updateFlightTimesUseCase,
		getAllFlightsUseCase: getAllFlightsUseCase, deleteFlightUseCase: deleteFlightUseCase,
		searchFlightsUseCase:       searchFlightsUseCase,
		getSuggestedFlightsUseCase: getSuggestedFlightsUseCase,
	}
}

func (h *FlightHandler) CreateFlight(c *gin.Context) {
	var req dto.CreateFlightRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid request body, %v", err)})
		return
	}

	flightEntity := mappers.CreateFlightRequestToEntity(req)
	createdFlight, err := h.createFlightUseCase.Execute(c.Request.Context(), flightEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to create flight, %v", err)})
		return
	}

	response := mappers.CreateFlightEntityToResponse(createdFlight)
	c.JSON(http.StatusCreated, response)
}

func (h *FlightHandler) GetFlight(c *gin.Context) {
	flightIDStr := c.Query("id")
	if flightIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Flight ID is required."})
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Flight ID."})
		return
	}

	flightDetails, err := h.getFlightUseCase.Execute(c.Request.Context(), flightID)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Flight not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Flight details retrieved successfully.",
		"data":    flightDetails,
	})
}

func (h *FlightHandler) UpdateFlightTimes(c *gin.Context) {
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	flightIDStr := c.Query("id")
	if flightIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Flight ID is required."})
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Flight ID."})
		return
	}

	var request dto.UpdateFlightTimesRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid flight data. Please check the input fields."})
		return
	}

	response, err := h.updateFlightTimesUseCase.Execute(c.Request.Context(), flightID, request)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Flight not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Flight updated successfully.",
		"flight":  response,
	})
}

func (h *FlightHandler) GetAllFlights(c *gin.Context) {
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	flights, tickets, err := h.getAllFlightsUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	response := mappers.MapFlightsAndTicketsToResponse(flights, tickets)
	c.JSON(http.StatusOK, gin.H{
		"message": "Flights retrieved successfully.",
		"data":    response,
	})
}

func (h *FlightHandler) DeleteFlight(c *gin.Context) {
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	flightIDStr := c.Query("id")
	if flightIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Flight ID is required."})
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Flight ID."})
		return
	}

	err = h.deleteFlightUseCase.Execute(c.Request.Context(), flightID)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Flight not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully."})
}

func (h *FlightHandler) SearchFlights(c *gin.Context) {
	departureCity := c.Query("departureCity")
	arrivalCity := c.Query("arrivalCity")
	flightDate := c.Query("flightDate")

	if departureCity == "" || arrivalCity == "" || flightDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters. Please check departureCity, arrivalCity, and flightDate."})
		return
	}

	flights, err := h.searchFlightsUseCase.Execute(c.Request.Context(), departureCity, arrivalCity, flightDate)
	if err != nil {
		if errors.Is(err, adapters.ErrNoFlightsFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "No flights found for the given criteria."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Flights retrieved successfully.",
		"data":    flights,
	})
}

func (h *FlightHandler) GetSuggestedFlights(c *gin.Context) {
	flights, err := h.getSuggestedFlightsUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("An unexpected error occurred. %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Suggested flights retrieved successfully.",
		"data":    flights,
	})
}
