package coins

import (
	"database/sql"
	"fmt"
	"merch-shop/internal/utils"
	"sync"
)

func (repo *CoinsDBRepostitory) GetHistory(userID string) (*History, error) {
	hst := &History{}
	query := `SELECT sender_id, amount FROM coin_history WHERE receiver_id = $1;`
	errStr := "error while selecting the history of the input transactions"
	transactions, err := GetTransactions(repo.dtb, repo.mutex, query, errStr, userID)
	if err != nil {
		return nil, err
	}

	hst.Received = make([]Input, len(transactions))
	for idx, transaction := range transactions {
		hst.Received[idx].FromUser = transaction.User
		hst.Received[idx].Amount = transaction.Amount
	}

	errStr = "error while selecting the history of the output transactions"
	query = `SELECT receiver_id, amount FROM coin_history WHERE sender_id = $1;`
	transactions, err = GetTransactions(repo.dtb, repo.mutex, query, errStr, userID)
	if err != nil {
		return nil, err
	}
	hst.Sent = make([]Output, len(transactions))
	for idx, transaction := range transactions {
		hst.Sent[idx].ToUser = transaction.User
		hst.Sent[idx].Amount = transaction.Amount
	}

	return hst, nil
}

func GetTransactions(dtb *sql.DB, mutex *sync.Mutex, query, errStr string, userID string) ([]Transaction, error) {
	mutex.Lock()
	defer mutex.Unlock()

	rows, err := dtb.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errStr, err)
	}
	defer rows.Close()

	transactions := make([]Transaction, 0)
	for rows.Next() {
		transaction := Transaction{}
		err := rows.Scan(&transaction.User, &transaction.Amount)
		if err != nil {
			return nil, fmt.Errorf("error from method `Scan`, package sql: %v", err)
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating over rows returned by query: %v", err)
	}
	return transactions, nil
}

func GetOutput(dtb *sql.DB, mutex *sync.Mutex, userID string) ([]Input, error) {
	mutex.Lock()
	defer mutex.Unlock()

	query := `SELECT receiver_id, amount FROM coin_history WHERE sender_id = $1;`
	rows, err := dtb.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the history of the output transactions: %v", err)
	}
	defer rows.Close()

	inputs := make([]Input, 0)
	for rows.Next() {
		inp := Input{}
		var receiverID string
		err := rows.Scan(&receiverID, &inp.Amount)
		if err != nil {
			return nil, fmt.Errorf("error from method `Scan`, package sql: %v", err)
		}

		username, err := utils.GetUsername(dtb, receiverID)
		if err != nil {
			return nil, err
		}

		inp.FromUser = *username
		inputs = append(inputs, inp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating over rows returned by query: %v", err)
	}
	return inputs, nil
}
