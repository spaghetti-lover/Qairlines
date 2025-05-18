package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

// AirplaneModelCreateRequestToInput converts a request DTO to a use case input.
func AirplaneModelCreateRequestToInput(req dto.AirplaneModelCreateRequest) usecases.AirplaneModelCreateInput {
	return usecases.AirplaneModelCreateInput{
		Name:         req.Name,
		Manufacturer: req.Manufacturer,
		TotalSeats:   req.TotalSeats,
	}
}

// AirplaneModelCreateInputToRequest converts a use case input to a request DTO.
func AirplaneModelCreateOutputToResponse(output usecases.AirplaneModelCreateOutput) dto.AirplaneModelCreateResponse {
	return dto.AirplaneModelCreateResponse{
		AirplaneModelID: output.AirplaneModelID,
		Name:            output.Name,
		Manufacturer:    output.Manufacturer,
		TotalSeats:      output.TotalSeats,
	}
}
