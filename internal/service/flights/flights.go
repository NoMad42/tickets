package flights

import (
	"context"

	"homework/internal/domain/flights"
)

type FlightsService interface {
	GetFlightsList(context.Context) (flights.FlightsList, error)
}

type FlightsStorage interface {
	GetFlightsList(context.Context) (flights.FlightsList, error)
}

type service struct {
	flightsStorage FlightsStorage
}

func (s service) GetFlightsList(ctx context.Context) (flights.FlightsList, error) {
	return s.flightsStorage.GetFlightsList(ctx)
}

func NewFlightsService(
	flightsStorage FlightsStorage,
) FlightsService {
	return &service{
		flightsStorage: flightsStorage,
	}
}
