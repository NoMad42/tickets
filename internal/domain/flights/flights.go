package flights

import (
	"time"
)

type Flight struct {
	// Код рейса.
	Code string `json:"code"`

	// Время вылета.
	From time.Time `json:"from" db:"from_timestamp"`

	// Идентификатор аэропорта вылета.
	FromAirportId string `json:"from_airport_id" db:"from_airport_id"`

	// Идентификатор рейса.
	Id string `json:"id"`

	// Статус.
	Status string `json:"status"`

	// Время прилёта.
	To time.Time `json:"to" db:"to_timestamp"`

	// Идентификатор аэропорта прилёта.
	ToAirportId string `json:"to_airport_id" db:"to_airport_id"`
}

// FlightsList defines model for FlightsList.
type FlightsList []Flight
