package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

const timeout = 5

type Handler struct {
	bulletService *BulletService
}

func (handler *Handler) rootHandler(w http.ResponseWriter, _ *http.Request) {
	games, err := handler.bulletService.GetGames()
	if err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	checked := 0
	for _, game := range games {
		if game.Checked {
			checked++
		}
	}
	//nolint:mnd // allow 100
	percentage := float64(checked) / float64(len(games)) * 100
	fmt.Fprintf(w, "Geschafft: %.1f%%", percentage)
}

func StartServer(bulletService *BulletService) error {
	handler := &Handler{
		bulletService: bulletService,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.rootHandler)
	server := &http.Server{
		Addr:              ":80",
		Handler:           mux,
		ReadHeaderTimeout: timeout * time.Second,
	}

	err := server.ListenAndServe()
	return fmt.Errorf("error running server: %w", err)
}
