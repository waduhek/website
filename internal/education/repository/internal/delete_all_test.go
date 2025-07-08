package internal_test

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/pashagolub/pgxmock/v4"

	"github.com/waduhek/website/internal/education/repository/internal"
)

func TestDeleteAll(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectExec("TRUNCATE").
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	expRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	if err := expRepo.DeleteAll(context.Background()); err != nil {
		t.Fatalf("error while deleting all educations: %s", err)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}

func TestDeleteAllError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectExec("TRUNCATE").
		WillReturnError(errors.New("test error"))

	expRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	if err := expRepo.DeleteAll(context.Background()); err == nil {
		t.Fatal("expected error when deleting all educations")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}
