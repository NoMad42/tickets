package transactions

import (
	"context"
	"log"

	"homework/internal/domain/booking"
	"homework/internal/domain/seats"
	"homework/internal/domain/transactions"

	"github.com/google/uuid"
)

type TransactionsService interface {
	GetTransactionsList(context.Context) ([]transactions.Transaction, error)
	CreateTransaction(ctx context.Context, bookingId uuid.UUID) (uuid.UUID, error)
}

type TransactionsStorage interface {
	GetTransactionsList(context.Context) ([]transactions.Transaction, error)
	CreateTransaction(ctx context.Context, amount float64, userProfileId uuid.UUID) (uuid.UUID, error)
}

type BookingService interface {
	GetBookingById(ctx context.Context, id uuid.UUID) (booking.Booking, error)
	Approve(ctx context.Context, bookingId, transactionId uuid.UUID) error
}

type SeatService interface {
	GetSeatById(ctx context.Context, id uuid.UUID) (seats.Seat, error)
}

type service struct {
	transactionsStorage TransactionsStorage
	bookingService      BookingService
	seatService         SeatService
}

func (s service) GetTransactionsList(ctx context.Context) ([]transactions.Transaction, error) {
	return s.transactionsStorage.GetTransactionsList(ctx)
}

func (s service) CreateTransaction(ctx context.Context, bookingId uuid.UUID) (uuid.UUID, error) {
	b, err := s.bookingService.GetBookingById(ctx, bookingId)
	if err != nil {
		return uuid.Nil, err
	}
	log.Println(b)

	seat, err := s.seatService.GetSeatById(ctx, b.SeatId)
	if err != nil {
		return uuid.Nil, err
	}

	log.Println(seat.Price, b.UserProfileId)
	// TODO add calc for seat options
	id, err := s.transactionsStorage.CreateTransaction(ctx, seat.Price, b.UserProfileId)
	if err != nil {
		return uuid.Nil, err
	}

	err = s.bookingService.Approve(ctx, b.Id, id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func NewTransactionsService(
	transactionsStorage TransactionsStorage,
	bookingService BookingService,
	seatService SeatService,
) TransactionsService {
	return &service{
		transactionsStorage: transactionsStorage,
		bookingService:      bookingService,
		seatService:         seatService,
	}
}
