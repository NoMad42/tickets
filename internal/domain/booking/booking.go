package booking

import "github.com/google/uuid"

// BookingStatus Статус бронирования.
type BookingStatus string

// Defines values for BookingStatus.
const (
	BookingStatusBooked  BookingStatus = "booked"
	BookingStatusBooking BookingStatus = "booking"
	BookingStatusFree    BookingStatus = "free"
)

// Booking Бронирование
type Booking struct {
	// FlightId Идентификатор рейса.
	FlightId uuid.UUID `json:"flight_id"`

	// Id Идентификатор бронирования.
	Id uuid.UUID `json:"id"`

	// SeatId Идентификатор места.
	SeatId uuid.UUID `json:"seat_id"`

	// SeatOptionsIds Идентификаторы услуг для места
	SeatOptionsIds *[]uuid.UUID `json:"seat_options_ids,omitempty"`

	// Status Статус бронирования.
	Status BookingStatus `json:"status"`

	// TransactionId Идентификатор транзакции.
	TransactionId *uuid.UUID `json:"transaction_id,omitempty"`

	// UserProfileId Идентификатор пользователя.
	UserProfileId uuid.UUID `json:"user_profile_id"`
}
