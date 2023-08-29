package transactions

import (
	"context"
	"fmt"
	"log"

	"homework/internal/domain/transactions"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionsStorage interface {
	GetTransactionsList(context.Context) ([]transactions.Transaction, error)
	CreateTransaction(ctx context.Context, amount float64, userProfileId uuid.UUID) (uuid.UUID, error)
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

func (s storage) CreateTransaction(ctx context.Context, amount float64, userProfileId uuid.UUID) (uuid.UUID, error) {
	tx, err := s.dbp.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return uuid.Nil, err
	}

	defer tx.Rollback(context.Background())

	transactionId := uuid.New()
	ct, err := tx.Exec(
		ctx,
		"INSERT INTO transactions (id, amount, user_profiles_id) VALUES($1, $2, $3)",
		transactionId,
		amount,
		userProfileId,
	)
	if err != nil {
		return uuid.Nil, err
	}
	if ct.RowsAffected() != 1 {
		return uuid.Nil, fmt.Errorf(
			"transaction storage error: при добавлении количество затронутых строк не равно 1. затронутых строк %d",
			ct.RowsAffected(),
		)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return transactionId, nil
}

func NewTransactionsStorage(
	dbp *pgxpool.Pool,
) TransactionsStorage {
	return &storage{
		dbp: dbp,
	}
}
