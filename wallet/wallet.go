package wallet

type Wallet struct {
	ID      int     `json:"id"`
	Owner   string  `json:"name"`
	Balance float64 `json:"balance"`
}

type GetWalletResponse struct {
	ID      int
	Owner   string
	Balance float64
}

type SaveWalletRequest struct {
	Owner   string
	Balance float64
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

type BalanceRequest struct {
	Balance float64 `json:"balance"`
}
