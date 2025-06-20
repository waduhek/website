package internal_test

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/pashagolub/pgxmock/v4"

	"github.com/waduhek/website/internal/experience/repository/internal"
)

func TestDeleteAll(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectExec("TRUNCATE").
		WillReturnResult(pgxmock.NewResult("DELTE", 1))

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if err := expRepo.DeleteAll(context.Background()); err != nil {
		t.Fatalf("error while deleting all experiences: %s", err)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}

func TestDeleteAllReturnsError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectExec("TRUNCATE").
		WillReturnError(errors.New("test error"))

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if err := expRepo.DeleteAll(context.Background()); err == nil {
		t.Fatalf("expected error while deleting all experiences")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}
