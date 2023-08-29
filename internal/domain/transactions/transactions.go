package transactions

import (
	"time"

	"github.com/google/uuid"
)

// Transaction Транзакция
type Transaction struct {
	// Amount Сумма.
	Amount float64 `json:"amount"`

	// Id Идентификатор транзакции.
	Id uuid.UUID `json:"id"`

	// Timestamp Время оплаты.
	Timestamp *time.Time `json:"timestamp,omitempty" db:"created_at"`

	// UserProfileId Идентификатор пользователя.
	UserProfileId uuid.UUID `json:"user_profile_id" db:"user_profiles_id"`
}
