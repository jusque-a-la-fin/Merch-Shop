package coins

import (
	"fmt"
	"merch-shop/internal/utils"
)

func (repo *CoinsDBRepostitory) SendCoins(transaction TransactionInDetail) (int, error) {
	if transaction.Balance-transaction.Amount < 0 {
		return 400, fmt.Errorf("insufficient balance")
	}

	receiverID, err := utils.GeReceiverID(repo.dtb, transaction.ReceiverName)
	if err != nil {
		return -1, nil
	}

	exists, err := utils.CheckUser(repo.dtb, transaction.ReceiverName)
	if err != nil {
		return -1, err
	}

	if !exists {
		exists, err := utils.CheckShop(repo.dtb, transaction.ReceiverName)
		if err != nil {
			return -1, err
		}

		if !exists {
			return 400, fmt.Errorf("user or shop does not exist")
		}
	}

	query := `
		     UPDATE coins_balance
		     SET balance = $1
		     WHERE user_id = $2;`

	_, err = repo.dtb.Exec(query, transaction.Balance, transaction.SenderID)
	if err != nil {
		return 500, fmt.Errorf("error while updating the coins balance: %v", err)
	}

	query = `INSERT INTO coin_history (sender_id, receiver_id, amount) VALUES ($1, $2, $3);`
	_, err = repo.dtb.Exec(query, transaction.SenderID, *receiverID, transaction.Amount)
	if err != nil {
		return 500, fmt.Errorf("error while adding new transaction: %v", err)
	}
	return 200, nil
}
