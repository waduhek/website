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

var testInsertOneInput = models.EducationInputModel{
	Institute: "Test",
	Degree:    "Test",
	Major:     "Test",
	Grade:     "CGPI: 10/10",
	Location:  "Test, Test",
	StartDate: time.Time{},
	EndDate:   time.Time{},
}

var testInsertOneNamedArgs = pgx.NamedArgs{
	"institute":  "Test",
	"degree":     "Test",
	"major":      "Test",
	"grade":      "CGPI: 10/10",
	"location":   "Test, Test",
	"start_date": time.Time{},
	"end_date":   time.Time{},
}

func TestInsertOne(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating connection pool: %s", err)
	}
	defer mockConn.Close()

	mockConn.ExpectExec("INSERT INTO education").
		WithArgs(testInsertOneNamedArgs).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	eduRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	err = eduRepo.InsertOne(context.Background(), &testInsertOneInput)
	if err != nil {
		t.Fatalf("error while inserting education: %s", err)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}

func TestInsertOneError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating connection pool: %s", err)
	}
	defer mockConn.Close()

	mockConn.ExpectExec("INSERT INTO education").
		WithArgs(testInsertOneNamedArgs).
		WillReturnError(errors.New("test error"))

	eduRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	err = eduRepo.InsertOne(context.Background(), &testInsertOneInput)
	if err == nil {
		t.Fatalf("expected error while inserting education")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}
