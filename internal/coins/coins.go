package coins

type CoinsRepo interface {
	GetBalance(userID string) (*int, error)
	GetHistory(userID string) (*History, error)
	SendCoins(transaction TransactionInDetail) (int, error)
}

type History struct {
	Received []Input  `json:"received"`
	Sent     []Output `json:"sent"`
}

type Input struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type Output struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

type Transaction struct {
	User   string
	Amount int
}

type TransactionInDetail struct {
	SenderID     string
	ReceiverName string
	Balance      int
	Amount       int
}
