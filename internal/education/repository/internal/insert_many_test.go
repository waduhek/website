package internal_test

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"

	"github.com/waduhek/website/internal/education/models"
	"github.com/waduhek/website/internal/education/repository/internal"
)

var expectedInsertManyColumns = []string{
	"institute",
	"degree",
	"major",
	"grade",
	"location",
	"start_date",
	"end_date",
}

var testInsertManyInput = []models.EducationInputModel{
	{
		Institute: "Test",
		Degree:    "Test",
		Major:     "Test",
		Grade:     "CGPI: 10/10",
		Location:  "Test",
		StartDate: time.Time{},
		EndDate:   time.Time{},
	},
}

func TestInsertMany(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating connection pool: %s", err)
	}
	defer mockConn.Close()

	mockConn.ExpectCopyFrom(
		pgx.Identifier{"education"},
		expectedInsertManyColumns,
	).
		WillReturnResult(1)

	eduRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	err = eduRepo.InsertMany(context.Background(), testInsertManyInput...)
	if err != nil {
		t.Fatalf("error while inserting multiple educations: %s", err)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}

func TestInsertManyError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating connection pool: %s", err)
	}
	defer mockConn.Close()

	mockConn.ExpectCopyFrom(pgx.Identifier{"education"}, expectedInsertManyColumns).
		WillReturnError(errors.New("test error"))

	eduRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	err = eduRepo.InsertMany(context.Background(), testInsertManyInput...)
	if err == nil {
		t.Fatalf("expected error while inserting multiple educations")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}
