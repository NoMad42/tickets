package airports

import (
	"context"

	"homework/internal/domain/airports"
)

type AirportsService interface {
	GetAirportsList(context.Context) (airports.AirportsList, error)
}

type AirportsStorage interface {
	GetAirportsList(context.Context) (airports.AirportsList, error)
}

type service struct {
	airportsStorage AirportsStorage
}

func (s service) GetAirportsList(ctx context.Context) (airports.AirportsList, error) {
	return s.airportsStorage.GetAirportsList(ctx)
}

func NewAirportsService(
	airportsStorage AirportsStorage,
) AirportsService {
	return &service{
		airportsStorage: airportsStorage,
	}
}
