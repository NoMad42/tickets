package seats

import (
	"homework/internal/domain/booking"

	"github.com/google/uuid"
)

// Defines values for SeatPosition.
const (
	Aisle  SeatPosition = "aisle"
	Middle SeatPosition = "middle"
	Window SeatPosition = "window"
)

// Defines values for SeatType.
const (
	Buisness SeatType = "buisness"
	Economy  SeatType = "economy"
)

// Seat defines model for Seat.
type Seat struct {
	// BookingStatus Статус бронирования.
	BookingStatus booking.BookingStatus `json:"booking_status" db:"booking_status"`

	// Code Код места.
	Code string `json:"code"`

	// FlightId Идентификатор рейса.
	FlightId uuid.UUID `json:"flight_id" db:"flight_id"`

	// Id Идентификатор места.
	Id uuid.UUID `json:"id"`

	// Position Место места.
	Position SeatPosition `json:"position"`

	// Price Цена места.
	Price float64 `json:"price"`

	// Type Тип места.
	Type SeatType `json:"type" db:"seat_type"`
}

// SeatPosition Место места.
type SeatPosition string

// SeatType Тип места.
type SeatType string

// SeatOption defines model for SeatOption.
type SeatOption struct {
	// Description Описание услуги для места.
	Description *string `json:"description,omitempty"`

	// Id Идентификатор услуги для места.
	Id uuid.UUID `json:"id"`

	// Name Название услуги для места.
	Name string `json:"name"`

	// Price Цена места.
	Price float64 `json:"price"`
}

// SeatOptionsList defines model for SeatOptionsList.
type SeatOptionsList = []SeatOption

// SeatsList defines model for SeatsList.
type SeatsList = []Seat
