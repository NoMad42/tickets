package transactions

import (
	"context"
	"log"

	"homework/internal/domain/transactions"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionsStorage interface {
	GetTransactionsList(context.Context) ([]transactions.Transaction, error)
	CreateTransaction(ctx context.Context, amount float64, userProfileId string) (string, error)
}

type storage struct {
	dbp *pgxpool.Pool
}

func (s storage) GetTransactionsList(ctx context.Context) ([]transactions.Transaction, error) {
	rows, _ := s.dbp.Query(context.Background(), "select * from transactions limit 100")
	defer rows.Close()

	a, err := pgx.CollectRows(rows, pgx.RowToStructByName[transactions.Transaction])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}

	return a, err
}

func (s storage) CreateTransaction(ctx context.Context, amount float64, userProfileId string) (string, error) {
	tx, err := s.dbp.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", err
	}

	defer tx.Rollback(context.Background())

	lastInsertId := ""
	err = tx.QueryRow(
		ctx,
		"INSERT INTO transactions (amount, user_profiles_id) VALUES($1, $2) RETURNING id",
		amount,
		userProfileId,
	).Scan(&lastInsertId)

	if err != nil {
		return lastInsertId, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return lastInsertId, nil
}

func NewTransactionsStorage(
	dbp *pgxpool.Pool,
) TransactionsStorage {
	return &storage{
		dbp: dbp,
	}
}
