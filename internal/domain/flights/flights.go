package flights

import (
	"time"
)

type Flight struct {
	// Код рейса.
	Code string `json:"code"`

	// Время вылета.
	From_timestamp time.Time `json:"from"`

	// Идентификатор аэропорта вылета.
	From_airport_id string `json:"from_airport_id"`

	// Идентификатор рейса.
	Id string `json:"id"`

	// Статус.
	Status string `json:"status"`

	// Время прилёта.
	To_timestamp time.Time `json:"to"`

	// Идентификатор аэропорта прилёта.
	To_airport_id string `json:"to_airport_id"`
}

// FlightsList defines model for FlightsList.
type FlightsList []Flight
