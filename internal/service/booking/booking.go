package airports

import (
	"context"
	"errors"

	"homework/internal/domain/booking"
	"homework/internal/domain/seats"

	"github.com/google/uuid"
)

type BookingService interface {
	GetBookingList(context.Context) ([]booking.Booking, error)
	CreateBooking(ctx context.Context, seatId, userId uuid.UUID) (uuid.UUID, error)
	GetBookingById(ctx context.Context, id uuid.UUID) (booking.Booking, error)
	Approve(ctx context.Context, bookingId, transactionId uuid.UUID) error
}

type BookingStorage interface {
	GetBookingList(context.Context) ([]booking.Booking, error)
	CreateBooking(ctx context.Context, flightId, seatId, userId uuid.UUID) (uuid.UUID, error)
	GetBookingById(ctx context.Context, id uuid.UUID) (booking.Booking, error)
	Approve(ctx context.Context, bookingId, transactionId uuid.UUID) error
}

type SeatService interface {
	GetSeatById(ctx context.Context, id uuid.UUID) (seats.Seat, error)
}

type service struct {
	bookingStorage BookingStorage
	seatService    SeatService
}

func (s service) GetBookingList(ctx context.Context) ([]booking.Booking, error) {
	return s.bookingStorage.GetBookingList(ctx)
}

func (s service) CreateBooking(ctx context.Context, seatId, userId uuid.UUID) (uuid.UUID, error) {
	seat, err := s.seatService.GetSeatById(ctx, seatId)
	if err != nil {
		return uuid.Nil, err
	}

	if seat.FlightId == uuid.Nil {
		return uuid.Nil, errors.New("seat flight id is missing")
	}

	id, err := s.bookingStorage.CreateBooking(ctx, seat.FlightId, seat.Id, userId)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s service) GetBookingById(ctx context.Context, id uuid.UUID) (booking.Booking, error) {
	return s.bookingStorage.GetBookingById(ctx, id)
}

func (s service) Approve(ctx context.Context, bookingId, transactionId uuid.UUID) error {
	return s.bookingStorage.Approve(ctx, bookingId, transactionId)
}

func NewBookingService(
	bookingStorage BookingStorage,
	seatService SeatService,
) BookingService {
	return &service{
		bookingStorage: bookingStorage,
		seatService:    seatService,
	}
}
