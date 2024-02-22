package api

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success code
	Code int

	// Account Balance
	Balance int64
}

type Error struct {
	// Error Code
	Code int

	// Error Message
	Message string
}
