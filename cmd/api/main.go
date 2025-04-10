package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type Config struct {
	port string
	dsn  string
}

type Application struct {
	config Config
	logger *log.Logger
}

func main() {
	cfg := Config{
		port: os.Getenv("API_PORT"),
		dsn:  os.Getenv("DATABASE_URL"),
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	logger.Printf("database connection established")

	err = migrateDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("database migration ran successfully")

	app := &Application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting server on port %s", srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrateDB(cfg Config) error {
	m, err := migrate.New("file://migrations", cfg.dsn)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && errors.Is(err, migrate.ErrNoChange) == false {
		return err
	}

	return nil
}
