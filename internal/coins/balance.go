package coins

import (
	"fmt"
)

func (repo *CoinsDBRepostitory) GetBalance(userID int) (*int, error) {
	var balance int
	err := repo.dtb.QueryRow("SELECT balance FROM coins_balance WHERE user_id = $1", userID).Scan(&balance)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the balance of coins: %v", err)
	}
	return &balance, nil
}
