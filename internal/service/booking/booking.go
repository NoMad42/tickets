package airports

import (
	"context"

	"homework/internal/domain/booking"
)

type BookingService interface {
	GetBookingList(context.Context) (booking.BookingList, error)
}

type BookingStorage interface {
	GetBookingList(context.Context) (booking.BookingList, error)
}

type service struct {
	bookingStorage BookingStorage
}

func (s service) GetBookingList(ctx context.Context) (booking.BookingList, error) {
	return s.bookingStorage.GetBookingList(ctx)
}

func NewBookingService(
	bookingStorage BookingStorage,
) BookingService {
	return &service{
		bookingStorage: bookingStorage,
	}
}
