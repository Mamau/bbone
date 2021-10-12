package db_connection

import (
	"bbone/internal/config"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func Connection(ctx context.Context) *sqlx.DB {
	settings := config.InitConfig(config.CONFIG_NAME).Database

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Password,
		settings.DbName,
		settings.SSLMode,
	)

	connect, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		log.Error().Err(err).Msg("error while connection to db")
	}

	if err = connect.Ping(); err != nil {
		log.Error().Msgf("error while ping to database, err: %v", err.Error())
	}

	return connect
}
