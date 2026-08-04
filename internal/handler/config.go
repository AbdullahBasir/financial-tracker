package handler

import (
	"database/sql"
	"log"
	"os"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	dbQueries *sqlc.Queries
	JwtSecret string
}

func LoadConfig() apiConfig {
	godotenv.Load()

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL must be set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if dbUrl == "" {
		log.Fatal("JWT_SECRET must be set")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}

	dbQueries := sqlc.New(db)

	return apiConfig{
		dbQueries: dbQueries,
		JwtSecret: jwtSecret,
	}
}
