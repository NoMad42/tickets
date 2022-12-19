package seats

import (
	"context"

	"homework/internal/domain/seats"
)

type SeatsService interface {
	GetSeatsList(context.Context) (seats.SeatsList, error)
	GetSeatOptionsList(context.Context) (seats.SeatOptionsList, error)
}

type SeatsStorage interface {
	GetSeatsList(context.Context) (seats.SeatsList, error)
	GetSeatOptionsList(context.Context) (seats.SeatOptionsList, error)
}

type service struct {
	seatsStorage SeatsStorage
}

func (s service) GetSeatsList(ctx context.Context) (seats.SeatsList, error) {
	return s.seatsStorage.GetSeatsList(ctx)
}

func (s service) GetSeatOptionsList(ctx context.Context) (seats.SeatOptionsList, error) {
	return s.seatsStorage.GetSeatOptionsList(ctx)
}

func NewSeatsService(
	seatsStorage SeatsStorage,
) SeatsService {
	return &service{
		seatsStorage: seatsStorage,
	}
}
