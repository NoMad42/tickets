package transactions

import "time"

// Transaction defines model for Transaction.
type Transaction struct {
	// Сумма.
	Amount float64 `json:"amount"`

	// Идентификатор транзакции.
	Id string `json:"id"`

	// Время оплаты.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// TransactionsList defines model for TransactionsList.
type TransactionsList []Transaction
