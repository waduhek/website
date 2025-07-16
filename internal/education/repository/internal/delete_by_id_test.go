package internal_test

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/pashagolub/pgxmock/v4"

	"github.com/waduhek/website/internal/education/repository/internal"
)

func TestDeleteByID(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectExec("DELETE FROM education").
		WithArgs(int32(1)).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	expRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	if err := expRepo.DeleteByID(context.Background(), 1); err != nil {
		t.Fatalf("error while deleting education by id: %s", err)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}

func TestDeleteByIDError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectExec("DELETE FROM education").
		WithArgs(int32(1)).
		WillReturnError(errors.New("test error"))

	expRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	if err := expRepo.DeleteByID(context.Background(), 1); err == nil {
		t.Fatalf("expected error while deleting education by id")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}
