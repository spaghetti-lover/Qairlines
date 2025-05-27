package dto

type GetBookingResponse struct {
	BookingID          string   `json:"bookingId"`
	Email              string   `json:"email"`
	TripType           string   `json:"tripType"`
	DepartureFlightID  string   `json:"departureFlightId"`
	ReturnFlightID     string   `json:"returnFlightId"`
	DepartureIDTickets []string `json:"departureIdTickets"`
	ReturnIDTickets    []string `json:"returnIdTickets"`
	CreatedAt          string   `json:"createdAt"`
	UpdatedAt          string   `json:"updatedAt"`
}
