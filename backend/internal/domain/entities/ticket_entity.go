package entities

import "time"

type TicketStatus string

const (
	TicketStatusBooked    TicketStatus = "booked"
	TicketStatusCancelled TicketStatus = "cancelled"
	TicketStatusUsed      TicketStatus = "used"
)

type FlightClass string

const (
	FlightClassEconomy    FlightClass = "Economy"
	FlightClassBusiness   FlightClass = "Business"
	FlightClassFirstClass FlightClass = "First Class"
)

type Ticket struct {
	TicketID    int64        `json:"ticket_id"`    // ID của vé
	SeatID      int64        `json:"seat_id"`      // ID của chỗ ngồi
	FlightClass FlightClass  `json:"flight_class"` // Hạng ghế
	Price       int          `json:"price"`        // Giá vé
	Status      TicketStatus `json:"status"`       // Trạng thái vé
	BookingID   int64        `json:"booking_id"`   // ID của booking
	FlightID    int64        `json:"flight_id"`    // ID của chuyến bay
	CreatedAt   time.Time    `json:"created_at"`   // Thời gian tạo
	UpdatedAt   time.Time    `json:"updated_at"`   // Thời gian cập nhật
	Seat        Seat         `json:"seat"`         // Thông tin chỗ ngồi (liên kết với bảng Seats)
	Owner       TicketOwner  `json:"owner"`        // Thông tin chủ sở hữu vé (liên kết với bảng TicketOwnerSnapshot)
}
