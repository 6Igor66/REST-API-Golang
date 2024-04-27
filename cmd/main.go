package main

import (
	"log"
	"myapp/internal/app/apiserver"
	"myapp/internal/storage"

	"myapp/internal/config"

	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.ReadConfigFromFile("../config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := storage.NewPostgreSQL(cfg.Postgresql.ConnString)

	if err != nil {
		log.Printf("failed to connect db: %v", err)
	}

	a, err := apiserver.New(db)

	if err != nil {
		log.Printf("failed to init server: %v", err)
	}

	err = a.Start(cfg.Http.Port)
	if err != nil {
		log.Printf("failed to start server: %v", err)
	}
}
