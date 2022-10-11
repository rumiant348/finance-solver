package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)


func Connection() *sql.DB {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "postgres://localhost:5432/postgres"
	}

	db, err := sql.Open("pgx", databaseUrl)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return db
}
