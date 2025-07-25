// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type BookingStatus string

const (
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
	BookingStatusPending   BookingStatus = "pending"
)

func (e *BookingStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BookingStatus(s)
	case string:
		*e = BookingStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for BookingStatus: %T", src)
	}
	return nil
}

type NullBookingStatus struct {
	BookingStatus BookingStatus `json:"booking_status"`
	Valid         bool          `json:"valid"` // Valid is true if BookingStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBookingStatus) Scan(value interface{}) error {
	if value == nil {
		ns.BookingStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BookingStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBookingStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BookingStatus), nil
}

type FlightClass string

const (
	FlightClassEconomy    FlightClass = "economy"
	FlightClassBusiness   FlightClass = "business"
	FlightClassFirstClass FlightClass = "firstClass"
)

func (e *FlightClass) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FlightClass(s)
	case string:
		*e = FlightClass(s)
	default:
		return fmt.Errorf("unsupported scan type for FlightClass: %T", src)
	}
	return nil
}

type NullFlightClass struct {
	FlightClass FlightClass `json:"flight_class"`
	Valid       bool        `json:"valid"` // Valid is true if FlightClass is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFlightClass) Scan(value interface{}) error {
	if value == nil {
		ns.FlightClass, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FlightClass.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFlightClass) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FlightClass), nil
}

type FlightStatus string

const (
	FlightStatusOnTime    FlightStatus = "On Time"
	FlightStatusDelayed   FlightStatus = "Delayed"
	FlightStatusCancelled FlightStatus = "Cancelled"
	FlightStatusBoarding  FlightStatus = "Boarding"
	FlightStatusTakeoff   FlightStatus = "Takeoff"
	FlightStatusLanding   FlightStatus = "Landing"
	FlightStatusLanded    FlightStatus = "Landed"
)

func (e *FlightStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FlightStatus(s)
	case string:
		*e = FlightStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for FlightStatus: %T", src)
	}
	return nil
}

type NullFlightStatus struct {
	FlightStatus FlightStatus `json:"flight_status"`
	Valid        bool         `json:"valid"` // Valid is true if FlightStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFlightStatus) Scan(value interface{}) error {
	if value == nil {
		ns.FlightStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FlightStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFlightStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FlightStatus), nil
}

type GenderType string

const (
	GenderTypeMale   GenderType = "Male"
	GenderTypeFemale GenderType = "Female"
	GenderTypeOther  GenderType = "Other"
)

func (e *GenderType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = GenderType(s)
	case string:
		*e = GenderType(s)
	default:
		return fmt.Errorf("unsupported scan type for GenderType: %T", src)
	}
	return nil
}

type NullGenderType struct {
	GenderType GenderType `json:"gender_type"`
	Valid      bool       `json:"valid"` // Valid is true if GenderType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGenderType) Scan(value interface{}) error {
	if value == nil {
		ns.GenderType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.GenderType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGenderType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.GenderType), nil
}

type TicketStatus string

const (
	TicketStatusActive    TicketStatus = "Active"
	TicketStatusCancelled TicketStatus = "Cancelled"
)

func (e *TicketStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TicketStatus(s)
	case string:
		*e = TicketStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for TicketStatus: %T", src)
	}
	return nil
}

type NullTicketStatus struct {
	TicketStatus TicketStatus `json:"ticket_status"`
	Valid        bool         `json:"valid"` // Valid is true if TicketStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTicketStatus) Scan(value interface{}) error {
	if value == nil {
		ns.TicketStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TicketStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTicketStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TicketStatus), nil
}

type TripType string

const (
	TripTypeOneWay    TripType = "oneWay"
	TripTypeRoundTrip TripType = "roundTrip"
)

func (e *TripType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TripType(s)
	case string:
		*e = TripType(s)
	default:
		return fmt.Errorf("unsupported scan type for TripType: %T", src)
	}
	return nil
}

type NullTripType struct {
	TripType TripType `json:"trip_type"`
	Valid    bool     `json:"valid"` // Valid is true if TripType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTripType) Scan(value interface{}) error {
	if value == nil {
		ns.TripType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TripType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTripType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TripType), nil
}

type UserRole string

const (
	UserRoleCustomer UserRole = "customer"
	UserRoleAdmin    UserRole = "admin"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"user_role"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Admin struct {
	UserID int64 `json:"user_id"`
}

type Booking struct {
	BookingID         int64         `json:"booking_id"`
	UserEmail         pgtype.Text   `json:"user_email"`
	TripType          TripType      `json:"trip_type"`
	DepartureFlightID pgtype.Int8   `json:"departure_flight_id"`
	ReturnFlightID    pgtype.Int8   `json:"return_flight_id"`
	Status            BookingStatus `json:"status"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

type Customer struct {
	UserID               int64       `json:"user_id"`
	PhoneNumber          pgtype.Text `json:"phone_number"`
	Gender               GenderType  `json:"gender"`
	DateOfBirth          time.Time   `json:"date_of_birth"`
	PassportNumber       pgtype.Text `json:"passport_number"`
	IdentificationNumber pgtype.Text `json:"identification_number"`
	Address              pgtype.Text `json:"address"`
	LoyaltyPoints        pgtype.Int4 `json:"loyalty_points"`
}

type Flight struct {
	FlightID         int64        `json:"flight_id"`
	FlightNumber     string       `json:"flight_number"`
	Airline          pgtype.Text  `json:"airline"`
	AircraftType     pgtype.Text  `json:"aircraft_type"`
	DepartureCity    pgtype.Text  `json:"departure_city"`
	ArrivalCity      pgtype.Text  `json:"arrival_city"`
	DepartureAirport pgtype.Text  `json:"departure_airport"`
	ArrivalAirport   pgtype.Text  `json:"arrival_airport"`
	DepartureTime    time.Time    `json:"departure_time"`
	ArrivalTime      time.Time    `json:"arrival_time"`
	BasePrice        int32        `json:"base_price"`
	TotalSeatsRow    int32        `json:"total_seats_row"`
	TotalSeatsColumn int32        `json:"total_seats_column"`
	Status           FlightStatus `json:"status"`
}

type News struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
	Content     pgtype.Text `json:"content"`
	Image       pgtype.Text `json:"image"`
	AuthorID    pgtype.Int8 `json:"author_id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Seat struct {
	SeatID      int64       `json:"seat_id"`
	FlightID    pgtype.Int8 `json:"flight_id"`
	SeatCode    string      `json:"seat_code"`
	IsAvailable bool        `json:"is_available"`
	Class       FlightClass `json:"class"`
}

type Ticket struct {
	TicketID    int64        `json:"ticket_id"`
	SeatID      int64        `json:"seat_id"`
	FlightClass FlightClass  `json:"flight_class"`
	Price       int32        `json:"price"`
	Status      TicketStatus `json:"status"`
	BookingID   pgtype.Int8  `json:"booking_id"`
	FlightID    int64        `json:"flight_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type Ticketownersnapshot struct {
	TicketID             int64       `json:"ticket_id"`
	FirstName            pgtype.Text `json:"first_name"`
	LastName             pgtype.Text `json:"last_name"`
	PhoneNumber          pgtype.Text `json:"phone_number"`
	Gender               GenderType  `json:"gender"`
	DateOfBirth          time.Time   `json:"date_of_birth"`
	PassportNumber       pgtype.Text `json:"passport_number"`
	IdentificationNumber pgtype.Text `json:"identification_number"`
	Address              pgtype.Text `json:"address"`
}

type User struct {
	UserID         int64              `json:"user_id"`
	Email          string             `json:"email"`
	HashedPassword string             `json:"hashed_password"`
	FirstName      pgtype.Text        `json:"first_name"`
	LastName       pgtype.Text        `json:"last_name"`
	Role           UserRole           `json:"role"`
	IsActive       bool               `json:"is_active"`
	DeletedAt      pgtype.Timestamptz `json:"deleted_at"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
