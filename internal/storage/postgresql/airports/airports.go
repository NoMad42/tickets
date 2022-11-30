package airports

import (
	"context"
	"log"

	"homework/internal/domain/airports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AirportsStorage interface {
	GetAirportsList(context.Context) (airports.AirportsList, error)
}

type storage struct {
	dbp *pgxpool.Pool
}

func (s storage) GetAirportsList(ctx context.Context) (airports.AirportsList, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from airports limit 100")
	defer rows.Close()

	a, err := pgx.CollectRows(rows, pgx.RowToStructByName[airports.Airport])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return a, err
}

func NewAirportsStorage(
	dbp *pgxpool.Pool,
) AirportsStorage {
	return &storage{
		dbp: dbp,
	}
}
