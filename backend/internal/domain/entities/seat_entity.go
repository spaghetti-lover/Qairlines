package entities

type Seat struct {
	SeatID      int64        `json:"seat_id"`
	FlightID    int64        `json:"flight_id"`
	SeatCode    string       `json:"seat_code"`
	IsAvailable bool         `json:"is_available"`
	FlightClass FlightStatus `json:"flight_class"`
}
