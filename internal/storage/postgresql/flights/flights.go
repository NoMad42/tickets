package flights

import (
	"context"
	"log"

	"homework/internal/domain/flights"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FlightsStorage interface {
	GetFlightsList(context.Context) (flights.FlightsList, error)
}

type storage struct {
	dbp *pgxpool.Pool
}

func (s storage) GetFlightsList(ctx context.Context) (flights.FlightsList, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from flights limit 100")
	defer rows.Close()

	a, err := pgx.CollectRows(rows, pgx.RowToStructByName[flights.Flight])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return a, err
}

func NewFlightsStorage(
	dbp *pgxpool.Pool,
) FlightsStorage {
	return &storage{
		dbp: dbp,
	}
}
