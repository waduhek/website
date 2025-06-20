package internal_test

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"

	"github.com/waduhek/website/internal/experience/models"
	"github.com/waduhek/website/internal/experience/repository/internal"
)

var now = time.Now()

var testInput = models.ExperienceInputModel{
	Title:       "Test",
	CompanyName: "Test",
	StartDate:   now,
	IsCurrent:   true,
	Location:    "Test",
	Description: []string{"Test", "Test"},
	Skills:      []string{"Test", "Test"},
}

var testInputNamedArgs = pgx.NamedArgs{
	"title":        "Test",
	"company_name": "Test",
	"start_date":   now,
	"end_date":     time.Time{},
	"is_current":   true,
	"location":     "Test",
	"description":  []string{"Test", "Test"},
	"skills":       []string{"Test", "Test"},
}

func TestInsertOne(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("failed to create a new mock pool: %s", err)
	}
	defer mockConn.Close()

	mockConn.ExpectExec("INSERT INTO experience").
		WithArgs(testInputNamedArgs).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if err := expRepo.InsertOne(context.Background(), &testInput); err != nil {
		t.Fatalf("error while inserting experience: %s", err)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectation: %s", err)
	}
}

func TestInsertOneReturnsError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("failed to create a new mock pool: %s", err)
	}
	defer mockConn.Close()

	mockConn.ExpectExec("INSERT INTO experience").
		WithArgs(testInputNamedArgs).
		WillReturnError(errors.New("test error"))

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if err := expRepo.InsertOne(context.Background(), &testInput); err == nil {
		t.Fatalf("expected error while inserting experience")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectation: %s", err)
	}
}
