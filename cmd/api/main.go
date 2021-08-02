package main

import (
	"flag"
	"log"
	"os"

	"github.com/valentedev/elements/internal/data"
	"github.com/valentedev/elements/internal/jsonlog"

	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	models data.Models
	logger *jsonlog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-open-conns", 25, "PostgresQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgresQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgresQL max connection idle time")
	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := &application{
		config: cfg,
		models: data.NewModels(db),
		logger: logger,
	}

	err = app.serve()
	if err != nil {
		log.Println(err)
	}

	log.Printf("Listening on port: %v\n", cfg.port)
}
