package v1

import (
	"homework/specs"

	airportsService "homework/internal/service/airports"
	flightsService "homework/internal/service/flights"
)

// Быстрая проверка актуальности текущего интерфейса сервера.
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
	airportsService airportsService.AirportsService
	flightsService  flightsService.FlightsService
}

func NewAPIServer(
	airportsService airportsService.AirportsService,
	flightsService flightsService.FlightsService,
) specs.ServerInterface {
	return &apiServer{
		airportsService: airportsService,
		flightsService:  flightsService,
	}
}
