package v1

import (
	"homework/specs"

	airportsService "homework/internal/service/airports"
)

// Быстрая проверка актуальности текущего интерфейса сервера.
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
	airportsService airportsService.AirportsService
}

func NewAPIServer(
	airportsService airportsService.AirportsService,
) specs.ServerInterface {
	return &apiServer{
		airportsService: airportsService,
	}
}
