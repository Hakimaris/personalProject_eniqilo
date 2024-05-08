// db/db.go
package db

import (
	"context"
	"fmt"
	"os"
	"time"

	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func Config() (*pgxpool.Config, error) {
	const (
		defaultMaxConns         = 4
		defaultMinConns         = 0
		defaultMaxConnLifetime  = time.Hour
		defaultMaxConnIdleTime  = time.Minute * 30
		defaultHealthCheckPeriod = time.Minute
		defaultConnectTimeout   = time.Second * 5
	)

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PARAMS"),
	)

	dbConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	return dbConfig, nil
}

func Init() error { // this doenst run first actually, summon the func
	config, err := Config()
	if err != nil {
		return err
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}

	return nil
}

// func Query(query string, args ...interface{}) (pgx.Rows, error) { // make it less hassle to run query
// 	rows, err := pool.Query(context.Background(), query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rows, nil
// }
