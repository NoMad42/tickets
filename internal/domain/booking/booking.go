package booking

// Статус бронирования.
type BookingStatus string

// Defines values for BookingStatus.
const (
	BookingStatusBooked BookingStatus = "booked"

	BookingStatusBooking BookingStatus = "booking"

	BookingStatusFree BookingStatus = "free"
)

// Booking defines model for Booking.
type Booking struct {
	// Идентификатор рейса.
	FlightId string `json:"flight_id" db:"flight_id"`

	// Идентификатор услуги для места.
	Id string `json:"id"`

	// Идентификатор места.
	SeatId string `json:"seat_id" db:"seat_id"`

	// Идентификаторы услуг для места
	SeatOptionsIds *[]string `json:"seat_options_ids,omitempty" db:"-"`

	// Статус бронирования.
	Status *BookingStatus `json:"status,omitempty" db:"-"`

	// Идентификатор транзакции.
	TransactionId *string `json:"transaction_id,omitempty" db:"transaction_id"`

	// Идентификатор пользователя.
	UserProfileId string `json:"user_profile_id" db:"user_profile_id"`
}

// BookingList defines model for BookingList.
type BookingList []Booking
