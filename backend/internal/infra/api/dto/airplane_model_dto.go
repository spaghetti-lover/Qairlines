package dto

import "time"

type AirplaneModelCreateRequest struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	TotalSeats   int64  `json:"total_seats"`
}

type AirplaneModelCreateResponse struct {
	AirplaneModelID int64     `json:"airplane_model_id"`
	Name            string    `json:"name"`
	Manufacturer    string    `json:"manufacturer"`
	TotalSeats      int64     `json:"total_seats"`
	CreatedAt       time.Time `json:"created_at"`
}
