package seats

import (
	"context"

	"homework/internal/domain/seats"
)

type SeatsService interface {
	GetSeatsList(context.Context) (seats.SeatsList, error)
}

type SeatsStorage interface {
	GetSeatsList(context.Context) (seats.SeatsList, error)
}

type service struct {
	seatsStorage SeatsStorage
}

func (s service) GetSeatsList(ctx context.Context) (seats.SeatsList, error) {
	return s.seatsStorage.GetSeatsList(ctx)
}

func NewSeatsService(
	seatsStorage SeatsStorage,
) SeatsService {
	return &service{
		seatsStorage: seatsStorage,
	}
}
