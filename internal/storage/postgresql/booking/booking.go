package booking

import (
	"context"
	"errors"
	"log"

	"homework/internal/domain/booking"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingStorage interface {
	GetBookingList(context.Context) ([]booking.Booking, error)
	CreateBooking(ctx context.Context, flightId, seatId, userId uuid.UUID) (uuid.UUID, error)
	GetBookingById(ctx context.Context, id uuid.UUID) (booking.Booking, error)
	Approve(ctx context.Context, bookingId, transactionId uuid.UUID) error
}

type storage struct {
	dbp *pgxpool.Pool
}

func (s storage) GetBookingList(ctx context.Context) ([]booking.Booking, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from booking limit 100")
	defer rows.Close()

	a, err := pgx.CollectRows(rows, pgx.RowToStructByName[booking.Booking])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return a, err
}

func (s storage) CreateBooking(ctx context.Context, flightId, seatId, userId uuid.UUID) (uuid.UUID, error) {
	tx, err := s.dbp.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return uuid.Nil, err
	}

	defer tx.Rollback(context.Background())

	lastInsertId := uuid.New()
	err = tx.QueryRow(
		ctx,
		"INSERT INTO booking (flight_id, seats_id, user_profiles_id, status) VALUES($1, $2, $3, $4) RETURNING id",
		flightId,
		seatId,
		userId,
		booking.BookingStatusBooking,
	).Scan(&lastInsertId)

	if err != nil {
		return uuid.Nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return lastInsertId, nil
}

func (s storage) GetBookingById(ctx context.Context, id uuid.UUID) (booking.Booking, error) {
	b := booking.Booking{}
	err := s.dbp.QueryRow(
		context.Background(),
		"select id, flight_id, seats_id, transaction_id, user_profiles_id, status from booking where id = $1 limit 1",
		id,
	).Scan(
		&b.Id,
		&b.FlightId,
		&b.SeatId,
		&b.TransactionId,
		&b.UserProfileId,
		&b.Status,
	)
	if err != nil {
		log.Printf("scan error at GetBookingById: %v", err)
	}

	return b, err
}

func (s storage) Approve(ctx context.Context, bookingId, transactionId uuid.UUID) error {
	tx, err := s.dbp.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	commandTag, err := tx.Exec(
		ctx,
		"UPDATE booking SET transaction_id = $1, status = $2 WHERE id = $3",
		transactionId,
		booking.BookingStatusBooked,
		bookingId,
	)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		log.Printf("%#v\n", commandTag)
		log.Printf(
			"%#v | %#v | %#v\n",
			bookingId,
			booking.BookingStatusBooked,
			transactionId,
		)
		return errors.New("не одной строки не обнавлено")
	}
	if commandTag.RowsAffected() > 1 {
		log.Printf("%#v\n", commandTag)
		log.Printf(
			"%#v | %#v | %#v\n",
			bookingId,
			booking.BookingStatusBooked,
			transactionId,
		)
		return errors.New("обнавлено больше одной строки")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func NewBookingStorage(
	dbp *pgxpool.Pool,
) BookingStorage {
	return &storage{
		dbp: dbp,
	}
}
