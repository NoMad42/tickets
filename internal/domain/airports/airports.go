package airports

// Airport defines model for Airport.
type Airport struct {
	// Город аэропорта.
	City string `json:"city"`

	// Код аэропорта.
	Code string `json:"code"`

	// Страна аэропорта.
	Country string `json:"country"`

	// Описание аэропорта.
	Description *string `json:"description,omitempty"`

	// Идентификатор аэропорта.
	Id string `json:"id"`

	// Название аэропорта.
	Name string `json:"name"`
}

// AirportsList defines model for AirportsList.
type AirportsList []Airport
