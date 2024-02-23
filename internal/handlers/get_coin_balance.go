package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alexezm1/go_api/api"
	"github.com/alexezm1/go_api/internal/tools"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog/log"
)

func GetCoinBalance(res http.ResponseWriter, req *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, req.URL.Query())

	if err != nil {
		log.Error().Err(err).Msg("")
		api.InternalErrorHandler(res)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(res)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error().Err(err).Msg("")
		api.InternalErrorHandler(res)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		log.Error().Err(err).Msg("")
		api.InternalErrorHandler(res)
		return
	}
}
