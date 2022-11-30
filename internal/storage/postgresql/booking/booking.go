package booking

import (
	"context"
	"log"

	"homework/internal/domain/booking"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingStorage interface {
	GetBookingList(context.Context) (booking.BookingList, error)
}

type storage struct {
	dbp *pgxpool.Pool
}

func (s storage) GetBookingList(ctx context.Context) (booking.BookingList, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from booking limit 100")
	defer rows.Close()

	a, err := pgx.CollectRows(rows, pgx.RowToStructByName[booking.Booking])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return a, err
}

func NewBookingStorage(
	dbp *pgxpool.Pool,
) BookingStorage {
	return &storage{
		dbp: dbp,
	}
}
