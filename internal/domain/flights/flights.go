package flights

import (
	"time"

	"github.com/google/uuid"
)

// Flight defines model for Flight.
type Flight struct {
	// Code Код рейса.
	Code string `json:"code"`

	// From Время вылета.
	From time.Time `json:"from" db:"from_timestamp"`

	// FromAirportId Идентификатор аэропорта вылета.
	FromAirportId uuid.UUID `json:"from_airport_id" db:"from_airport_id"`

	// Id Идентификатор рейса.
	Id uuid.UUID `json:"id"`

	// Status Статус.
	Status string `json:"status"`

	// To Время прилёта.
	To time.Time `json:"to"`

	// ToAirportId Идентификатор аэропорта прилёта.
	ToAirportId uuid.UUID `json:"to_airport_id" db:"to_airport_id"`
}

// FlightsList defines model for FlightsList.
type FlightsList []Flight
