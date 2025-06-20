package database

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect establishes connection to the database using the connection string
// specified in the DATABASE_URL environment variable.
func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("environment variable DATABASE_URL is required")
	}

	return pgxpool.New(ctx, databaseUrl)
}
