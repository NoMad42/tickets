package transactions

import (
	"context"
	"log"

	"homework/internal/domain/transactions"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionsStorage interface {
	GetTransactionsList(context.Context) (transactions.TransactionsList, error)
}

type storage struct {
	dbp *pgxpool.Pool
}

func (s storage) GetTransactionsList(ctx context.Context) (transactions.TransactionsList, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from transactions limit 100")
	defer rows.Close()

	a, err := pgx.CollectRows(rows, pgx.RowToStructByName[transactions.Transaction])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return a, err
}

func NewTransactionsStorage(
	dbp *pgxpool.Pool,
) TransactionsStorage {
	return &storage{
		dbp: dbp,
	}
}
