package internal_test

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/pashagolub/pgxmock/v4"

	"github.com/waduhek/website/internal/experience/repository/internal"
)

func TestDeleteById(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectExec("DELETE FROM experience").
		WithArgs(int32(1)).
		WillReturnResult(pgxmock.NewResult("DELTE", 1))

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if err := expRepo.DeleteById(context.Background(), 1); err != nil {
		t.Fatalf("error while deleting experience by id: %s", err)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}

func TestDeleteByIdReturnsError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectExec("DELETE FROM experience").
		WithArgs(int32(1)).
		WillReturnError(errors.New("test error"))

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if err := expRepo.DeleteById(context.Background(), 1); err == nil {
		t.Fatalf("expected error while deleting experience by id")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}
