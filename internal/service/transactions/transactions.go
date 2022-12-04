package transactions

import (
	"context"

	"homework/internal/domain/transactions"
)

type TransactionsService interface {
	GetTransactionsList(context.Context) (transactions.TransactionsList, error)
}

type TransactionsStorage interface {
	GetTransactionsList(context.Context) (transactions.TransactionsList, error)
}

type service struct {
	transactionsStorage TransactionsStorage
}

func (s service) GetTransactionsList(ctx context.Context) (transactions.TransactionsList, error) {
	return s.transactionsStorage.GetTransactionsList(ctx)
}

func NewTransactionsService(
	transactionsStorage TransactionsStorage,
) TransactionsService {
	return &service{
		transactionsStorage: transactionsStorage,
	}
}
