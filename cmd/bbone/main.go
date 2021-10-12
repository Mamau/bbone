package main

import (
	"bbone/internal/config"
	"bbone/internal/router"
	"bbone/internal/tracing"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var conf *config.Config

func main() {
	if err := godotenv.Load(); err != nil {
		log.Err(err).Msg(".env file not found")
	}
	conf = config.InitConfig(config.CONFIG_NAME)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tracer io.Closer
	if os.Getenv("ENVIRONMENT") == "prod" {
		tracer = tracing.InitTracing()
	}

	//conn := db_connection.Connection(ctx)
	//defer func() {
	//	if err := conn.Close(); err != nil {
	//		log.Info().Msgf("error while close connection DB, err: %s", err.Error())
	//	}
	//}()

	srv := server()
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()

	// Stop by signal
	<-c

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msg(err.Error())
	}

	if err := tracer.Close(); err != nil {
		log.Err(err).Msg(err.Error())
	}
}

func server() *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", conf.REST.Port),
		Handler:      router.NewRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return srv
}
