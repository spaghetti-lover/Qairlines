package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

// AirplaneModelCreateRequestToInput converts a request DTO to a use case input.
func AirplaneModelCreateRequestToInput(req dto.AirplaneModelCreateRequest) entities.CreateAirplaneModelParams {
	return entities.CreateAirplaneModelParams{
		Name:         req.Name,
		Manufacturer: req.Manufacturer,
		TotalSeats:   req.TotalSeats,
	}
}

// AirplaneModelCreateInputToRequest converts a use case input to a request DTO.
func AirplaneModelCreateOutputToResponse(output entities.AirplaneModel) dto.AirplaneModelCreateResponse {
	return dto.AirplaneModelCreateResponse{
		AirplaneModelID: output.AirplaneModelID,
		Name:            output.Name,
		Manufacturer:    output.Manufacturer,
		TotalSeats:      output.TotalSeats,
	}
}
