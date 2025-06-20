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

var expectedColumns = []string{
	"title",
	"company_name",
	"start_date",
	"end_date",
	"is_current",
	"location",
	"description",
	"skills",
}

var inputExperiences = []models.ExperienceInputModel{
	{
		Title:       "Test",
		CompanyName: "Test",
		StartDate:   time.Now(),
		IsCurrent:   true,
		Location:    "Test",
		Description: []string{"Test", "Test"},
		Skills:      []string{"Test", "Test"},
	},
}

func TestInsertMany(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectCopyFrom(pgx.Identifier{"experience"}, expectedColumns).
		WillReturnResult(1)

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if err := expRepo.InsertMany(context.Background(), inputExperiences...); err != nil {
		t.Fatalf("error while inserting experiences: %s", err)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}

func TestInsertManyReturnsError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectCopyFrom(pgx.Identifier{"experience"}, expectedColumns).
		WillReturnError(errors.New("test error"))

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if err := expRepo.InsertMany(context.Background(), inputExperiences...); err == nil {
		t.Fatalf("expected error while inserting experiences")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}
