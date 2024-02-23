package middleware

import (
	"errors"
	"net/http"

	"github.com/alexezm1/go_api/api"
	"github.com/alexezm1/go_api/internal/tools"
	"github.com/rs/zerolog/log"
)

var UnAuhtorizedError = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var username string = req.URL.Query().Get("username")
		var token = req.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error().Err(UnAuhtorizedError).Msg("")
			api.RequestErrorHandler(res, UnAuhtorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(res)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error().Err(UnAuhtorizedError).Msg("")
			api.RequestErrorHandler(res, UnAuhtorizedError)
			return
		}

		next.ServeHTTP(res, req)
	})
}
