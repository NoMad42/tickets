package airports

import (
	"context"
	"errors"

	"homework/internal/domain/booking"
	"homework/internal/domain/seats"
)

type BookingService interface {
	GetBookingList(context.Context) (booking.BookingList, error)
	CreateBooking(ctx context.Context, seatId, userId string) (string, error)
	GetBookingById(ctx context.Context, id string) (booking.Booking, error)
}

type BookingStorage interface {
	GetBookingList(context.Context) (booking.BookingList, error)
	CreateBooking(ctx context.Context, flightId, seatId, userId string) (string, error)
	GetBookingById(ctx context.Context, id string) (booking.Booking, error)
}

type SeatService interface {
	GetSeatById(ctx context.Context, id string) (seats.Seat, error)
}

type service struct {
	bookingStorage BookingStorage
	seatService    SeatService
}

func (s service) GetBookingList(ctx context.Context) (booking.BookingList, error) {
	return s.bookingStorage.GetBookingList(ctx)
}

func (s service) CreateBooking(ctx context.Context, seatId, userId string) (string, error) {
	seat, err := s.seatService.GetSeatById(ctx, seatId)
	if err != nil {
		return "", err
	}

	if seat.FlightId == nil {
		return "", errors.New("seat flight id is missing")
	}

	id, err := s.bookingStorage.CreateBooking(ctx, *seat.FlightId, seat.Id, userId)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s service) GetBookingById(ctx context.Context, id string) (booking.Booking, error) {
	return s.bookingStorage.GetBookingById(ctx, id)
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
