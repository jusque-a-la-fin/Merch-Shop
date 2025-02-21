package coins

import (
	"fmt"
)

func (repo *CoinsDBRepostitory) GetBalance(userID string) (*int, error) {
	var balance int
	err := repo.dtb.QueryRow("SELECT balance FROM coins_balance WHERE user_id = $1", userID).Scan(&balance)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the coins balance: %v", err)
	}
	return &balance, nil
}

func (repo *CoinsDBRepostitory) UpdateBalance(userID string) error {
	query := `
	INSERT INTO coins_balance (balance, user_id)
    VALUES (200, $1)
    ON CONFLICT (user_id) DO UPDATE
    SET balance = excluded.balance, user_id = excluded.user_id;`

	_, err := repo.dtb.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error while updating the coins balances: %v", err)
	}
	return nil
}
