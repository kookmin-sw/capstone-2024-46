package main

import (
	"context"
	"log"
	"os"

	"private-llm-backend/internal/config"
	"private-llm-backend/internal/database"
)

var tables = []interface{}{
	// Add GORM models here
}

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	c, _, err := config.Load(context.Background(), "./configs/", env)
	if err != nil {
		log.Fatalf("failed to load config: %+v", err)
	}
	dbConfig := database.NewMySQLConfig(*c.MySqlDsn)
	db, err := database.OpenConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	for _, table := range tables {
		log.Printf("[%T] migrating...\n", table)
		err = db.AutoMigrate(table)
		if err != nil {
			log.Fatalf("[%T] failed to migrate table: %+v", table, err)
		}
		log.Printf("[%T] done!\n", table)
	}
}
