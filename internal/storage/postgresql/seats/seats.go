package seats

import (
	"context"
	"log"

	"homework/internal/domain/seats"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SeatsStorage interface {
	GetSeatsList(context.Context) (seats.SeatsList, error)
	GetSeatById(ctx context.Context, id string) (seats.Seat, error)
	GetSeatOptionsList(context.Context) (seats.SeatOptionsList, error)
}

type storage struct {
	dbp *pgxpool.Pool
}

func (s storage) GetSeatsList(ctx context.Context) (seats.SeatsList, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from seats limit 100")
	defer rows.Close()

	a, err := pgx.CollectRows(rows, pgx.RowToStructByName[seats.Seat])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return a, err
}

func (s storage) GetSeatById(ctx context.Context, id string) (seats.Seat, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from seats where id = $1 limit 1", id)
	defer rows.Close()

	seat, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[seats.Seat])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return seat, err
}

func (s storage) GetSeatOptionsList(ctx context.Context) (seats.SeatOptionsList, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from seat_options limit 100")
	defer rows.Close()

	a, err := pgx.CollectRows(rows, pgx.RowToStructByName[seats.SeatOption])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return a, err
}

func NewSeatsStorage(
	dbp *pgxpool.Pool,
) SeatsStorage {
	return &storage{
		dbp: dbp,
	}
}
