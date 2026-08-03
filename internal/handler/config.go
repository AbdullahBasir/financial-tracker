package handler

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	dbQueries *sqlc.Queries
}

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

func LoadConfig() apiConfig {
	godotenv.Load()

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL must be set")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}

	dbQueries := sqlc.New(db)

	return apiConfig{
		dbQueries: dbQueries,
	}
}
