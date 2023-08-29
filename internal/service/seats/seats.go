package seats

import (
	"context"

	"homework/internal/domain/seats"

	"github.com/google/uuid"
)

type SeatsService interface {
	GetSeatsList(context.Context) (seats.SeatsList, error)
	GetSeatOptionsList(context.Context) (seats.SeatOptionsList, error)
	GetSeatById(ctx context.Context, id uuid.UUID) (seats.Seat, error)
}

type SeatsStorage interface {
	GetSeatsList(context.Context) (seats.SeatsList, error)
	GetSeatOptionsList(context.Context) (seats.SeatOptionsList, error)
	GetSeatById(ctx context.Context, id uuid.UUID) (seats.Seat, error)
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

func (s service) GetSeatById(ctx context.Context, id uuid.UUID) (seats.Seat, error) {
	return s.seatsStorage.GetSeatById(ctx, id)
}

func NewSeatsService(
	seatsStorage SeatsStorage,
) SeatsService {
	return &service{
		seatsStorage: seatsStorage,
	}
}
