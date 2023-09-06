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
	FlightId uuid.UUID `json:"flight_id" db:"flight_id"`

	// Id Идентификатор бронирования.
	Id uuid.UUID `json:"id"`

	// SeatId Идентификатор места.
	SeatId uuid.UUID `json:"seat_id" db:"seats_id"`

	// SeatOptionsIds Идентификаторы услуг для места
	SeatOptionsIds *[]uuid.UUID `json:"seat_options_ids,omitempty" db:"-"`

	// Status Статус бронирования.
	Status BookingStatus `json:"status"`

	// TransactionId Идентификатор транзакции.
	TransactionId *uuid.UUID `json:"transaction_id,omitempty" db:"transaction_id"`

	// UserProfileId Идентификатор пользователя.
	UserProfileId uuid.UUID `json:"user_profile_id" db:"user_profiles_id"`
}
