package main

import (
	"fmt"
	"net/http"

	"github.com/alexezm1/go_api/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Hello World")

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API Service...")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
	}
}
