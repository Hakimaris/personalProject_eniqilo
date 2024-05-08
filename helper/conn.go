// db/db.go
package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func Config() (*pgxpool.Config, error) {
	const (
		defaultMaxConns          = 4
		defaultMinConns          = 0
		defaultMaxConnLifetime   = time.Hour
		defaultMaxConnIdleTime   = time.Minute * 30
		defaultHealthCheckPeriod = time.Minute
		defaultConnectTimeout    = time.Second * 5
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

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig, nil
}

func Init() error { // this doenst run first actually, summon the func
	config, err := Config()
	if err != nil {
		return err
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Error while creating connection to the database: %s", err)
	}

	connection, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Error while acquiring connection from the database pool: %s",err)
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatalf("Could not ping database : %s",err)
	}

	fmt.Println("Connected to the database!!")

	return nil
}

// func Query(query string, args ...interface{}) (pgx.Rows, error) { // make it less hassle to run query
// 	rows, err := pool.Query(context.Background(), query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rows, nil
// }
