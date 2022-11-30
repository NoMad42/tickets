package v1

import (
	"encoding/json"
	"homework/specs"
	"log"
	"net/http"

	airportsService "homework/internal/service/airports"
	bookingService "homework/internal/service/booking"
	flightsService "homework/internal/service/flights"
	transactionsService "homework/internal/service/transactions"
)

// Быстрая проверка актуальности текущего интерфейса сервера.
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
	airportsService     airportsService.AirportsService
	bookingService      bookingService.BookingService
	flightsService      flightsService.FlightsService
	transactionsService transactionsService.TransactionsService
}

func NewAPIServer(
	airportsService airportsService.AirportsService,
	bookingService bookingService.BookingService,
	flightsService flightsService.FlightsService,
	transactionsService transactionsService.TransactionsService,
) specs.ServerInterface {
	return &apiServer{
		airportsService:     airportsService,
		bookingService:      bookingService,
		flightsService:      flightsService,
		transactionsService: transactionsService,
	}
}

func (a *apiServer) writeSuccessResponse(result any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Fatal(err)
	}
}
