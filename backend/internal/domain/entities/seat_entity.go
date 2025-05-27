package entities

type Seat struct {
	SeatID      int64       `json:"seat_id"`      // ID của chỗ ngồi
	FlightID    int64       `json:"flight_id"`    // ID của chuyến bay
	SeatCode    string      `json:"seat_code"`    // Mã chỗ ngồi (ví dụ: 12A)
	IsAvailable bool        `json:"is_available"` // Trạng thái chỗ ngồi (còn trống hay không)
	Class       FlightClass `json:"class"`        // Hạng ghế
}
