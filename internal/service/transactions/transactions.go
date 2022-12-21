package transactions

import (
	"context"
	"log"

	"homework/internal/domain/booking"
	"homework/internal/domain/seats"
	"homework/internal/domain/transactions"
)

type TransactionsService interface {
	GetTransactionsList(context.Context) (transactions.TransactionsList, error)
	CreateTransaction(ctx context.Context, bookingId string) (string, error)
}

type TransactionsStorage interface {
	GetTransactionsList(context.Context) (transactions.TransactionsList, error)
	CreateTransaction(ctx context.Context, amount float64, userProfileId string) (string, error)
}

type BookingService interface {
	GetBookingById(ctx context.Context, id string) (booking.Booking, error)
}

type SeatService interface {
	GetSeatById(ctx context.Context, id string) (seats.Seat, error)
}

type service struct {
	transactionsStorage TransactionsStorage
	bookingService      BookingService
	seatService         SeatService
}

func (s service) GetTransactionsList(ctx context.Context) (transactions.TransactionsList, error) {
	return s.transactionsStorage.GetTransactionsList(ctx)
}

func (s service) CreateTransaction(ctx context.Context, bookingId string) (string, error) {
	log.Println(bookingId)
	b, err := s.bookingService.GetBookingById(ctx, bookingId)
	if err != nil {
		log.Println("problem is here")
		return "", err
	}
	log.Println(b)

	seat, err := s.seatService.GetSeatById(ctx, b.SeatId)
	if err != nil {
		return "", err
	}

	log.Println(seat.Price, b.UserProfileId)
	// TODO add calc for seat options
	id, err := s.transactionsStorage.CreateTransaction(ctx, seat.Price, b.UserProfileId)
	if err != nil {
		return "", err
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
