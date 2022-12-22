package transactions

import "time"

// Транзакция
type Transaction struct {
	// Сумма.
	Amount float64 `json:"amount"`

	// Идентификатор транзакции.
	Id string `json:"id"`

	// Время оплаты.
	Timestamp *time.Time `json:"timestamp,omitempty" db:"created_at"`

	// Идентификатор пользователя.
	UserProfileId string `json:"user_profile_id" db:"user_profiles_id"`
}
