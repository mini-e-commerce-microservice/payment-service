package main

import (
	"context"
	"errors"
	"github.com/mini-e-commerce-microservice/payment-service/internal/conf"
	"github.com/mini-e-commerce-microservice/payment-service/internal/presenter"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"net/http"
	"os/signal"
	"syscall"
)

var restApi = &cobra.Command{
	Use:   "restApi",
	Short: "restApi",
	Run: func(cmd *cobra.Command, args []string) {
		appConf := conf.LoadAppConf()

		server := presenter.New(&presenter.Presenter{
			Port: int(appConf.AppPort),
		})

		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		go func() {
			if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				ctx.Done()
			}
		}()

		<-ctx.Done()
		log.Info().Msg("Received shutdown signal, shutting down server gracefully...")
	},
}
