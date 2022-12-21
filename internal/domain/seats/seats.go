package seats

import "homework/internal/domain/booking"

// Defines values for SeatPosition.
const (
	SeatPositionAisle SeatPosition = "aisle"

	SeatPositionMiddle SeatPosition = "middle"

	SeatPositionWindow SeatPosition = "window"
)

// Defines values for SeatType.
const (
	SeatTypeBuisness SeatType = "buisness"

	SeatTypeEconomy SeatType = "economy"
)

// Seat defines model for Seat.
type Seat struct {
	// Статус бронирования.
	BookingStatus booking.BookingStatus `json:"booking_status" db:"booking_status"`

	// Код места.
	Code string `json:"code"`

	// Идентификатор рейса.
	FlightId *string `json:"flight_id,omitempty" db:"flight_id"`

	// Идентификатор места.
	Id string `json:"id"`

	// Место места.
	Position SeatPosition `json:"position"`

	// Цена места.
	Price float64 `json:"price"`

	// Тип места.
	Type SeatType `json:"type" db:"seat_type"`
}

// Место места.
type SeatPosition string

// Тип места.
type SeatType string

// SeatOption defines model for SeatOption.
type SeatOption struct {
	// Описание услуги для места.
	Description *string `json:"description,omitempty"`

	// Идентификатор услуги для места.
	Id string `json:"id"`

	// Название услуги для места.
	Name string `json:"name"`

	// Цена места.
	Price float64 `json:"price"`
}

// SeatOptionsList defines model for SeatOptionsList.
type SeatOptionsList []SeatOption

// SeatsList defines model for SeatsList.
type SeatsList []Seat
