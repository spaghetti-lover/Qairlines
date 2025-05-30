package entities

import (
	"database/sql"
	"time"
)

type TicketStatus string

const (
	TicketStatusActive    TicketStatus = "Booked"
	TicketStatusCancelled TicketStatus = "Cancelled"
)

type FlightClass string

const (
	FlightClassEconomy    FlightClass = "economy"
	FlightClassBusiness   FlightClass = "business"
	FlightClassFirstClass FlightClass = "firstClass"
)

type Ticket struct {
	TicketID    int64         `json:"ticket_id"`    // ID của vé
	SeatID      sql.NullInt64 `json:"seat_id"`      // ID của chỗ ngồi
	FlightClass FlightClass   `json:"flight_class"` // Hạng ghế
	Price       int32         `json:"price"`        // Giá vé
	Status      TicketStatus  `json:"status"`       // Trạng thái vé
	BookingID   int64         `json:"booking_id"`   // ID của booking
	FlightID    int64         `json:"flight_id"`    // ID của chuyến bay
	CreatedAt   time.Time     `json:"created_at"`   // Thời gian tạo
	UpdatedAt   time.Time     `json:"updated_at"`   // Thời gian cập nhật
	Seat        Seat          `json:"seat"`         // Thông tin chỗ ngồi (liên kết với bảng Seats)
	Owner       TicketOwner   `json:"owner"`        // Thông tin chủ sở hữu vé (liên kết với bảng TicketOwnerSnapshot)
}
