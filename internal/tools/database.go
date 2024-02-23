package tools

import "github.com/rs/zerolog/log"

// Databse collections
type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	return &database, nil
}
