package api

import (
	"encoding/json"
	"net/http"
)

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

func writeError(res http.ResponseWriter, message string, code int) {
	errRes := Error{
		Code:    code,
		Message: message,
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)

	json.NewEncoder(res).Encode(errRes)
}

var (
	RequestErrorHandler = func(res http.ResponseWriter, err error) {
		writeError(res, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(res http.ResponseWriter) {
		writeError(res, "An Unexpected Error Ocurred", http.StatusInternalServerError)
	}
)
