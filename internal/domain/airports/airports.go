package airports

import "github.com/google/uuid"

// Airport defines model for Airport.
type Airport struct {
	// City Город аэропорта.
	City string `json:"city"`

	// Code Код аэропорта.
	Code string `json:"code"`

	// Country Страна аэропорта.
	Country string `json:"country"`

	// Description Описание аэропорта.
	Description *string `json:"description,omitempty"`

	// Id Идентификатор аэропорта.
	Id uuid.UUID `json:"id"`

	// Name Название аэропорта.
	Name string `json:"name"`
}

// AirportsList defines model for AirportsList.
type AirportsList = []Airport
