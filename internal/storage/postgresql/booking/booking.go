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
	CreateBooking(ctx context.Context, flightId, seatId, userId string) (string, error)
	GetBookingById(ctx context.Context, id string) (booking.Booking, error)
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

func (s storage) CreateBooking(ctx context.Context, flightId, seatId, userId string) (string, error) {
	tx, err := s.dbp.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", err
	}

	defer tx.Rollback(context.Background())

	lastInsertId := ""
	err = tx.QueryRow(
		ctx,
		"INSERT INTO booking (flight_id, seats_id, user_profiles_id, status) VALUES($1, $2, $3, $4) RETURNING id",
		flightId,
		seatId,
		userId,
		booking.BookingStatusBooking,
	).Scan(&lastInsertId)

	if err != nil {
		return lastInsertId, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return lastInsertId, nil
}

func (s storage) GetBookingById(ctx context.Context, id string) (booking.Booking, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from booking where id = $1 limit 1", id)
	defer rows.Close()

	b, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[booking.Booking])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return b, err
}

func NewBookingStorage(
	dbp *pgxpool.Pool,
) BookingStorage {
	return &storage{
		dbp: dbp,
	}
}
