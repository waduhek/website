package internal_test

import (
	"context"
	"errors"
	"log/slog"
	"reflect"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"

	"github.com/waduhek/website/internal/education/models"
	"github.com/waduhek/website/internal/education/repository/internal"
)

func TestGetByID(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating connection pool: %s", err)
	}
	defer mockConn.Close()

	rows := pgxmock.NewRows(
		[]string{
			"id", "institute", "degree", "major", "grade", "location",
			"start_date", "end_date",
		},
	).
		AddRow(
			1, "Test", "Test", "Test", "CGPI: 10/10", "Test", time.Time{},
			time.Time{},
		)

	mockConn.ExpectQuery("SELECT").
		WithArgs(int32(1)).
		WillReturnRows(rows)

	eduRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	expectedEducation := &models.EducationOutputModel{
		ID:        1,
		Institute: "Test",
		Degree:    "Test",
		Major:     "Test",
		Grade:     "CGPI: 10/10",
		Location:  "Test",
		StartDate: time.Time{},
		EndDate:   time.Time{},
	}

	gotEducation, err := eduRepo.GetByID(context.Background(), 1)
	if err != nil {
		t.Fatalf("error while getting all educations: %s", err)
	}

	if !reflect.DeepEqual(expectedEducation, gotEducation) {
		t.Fatalf(
			"educations did not match.\nexpected: %v\ngot: %v",
			expectedEducation, gotEducation,
		)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}

func TestGetByIDNoRows(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating connection pool: %s", err)
	}
	defer mockConn.Close()

	rows := pgxmock.NewRows(
		[]string{
			"id", "institute", "degree", "major", "grade", "location",
			"start_date", "end_date",
		},
	)

	mockConn.ExpectQuery("SELECT").
		WithArgs(int32(1)).
		WillReturnRows(rows)

	eduRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	_, err = eduRepo.GetByID(context.Background(), 1)
	if !errors.Is(pgx.ErrNoRows, err) {
		t.Fatalf("expected pgx.ErrNoRows, got: %s", err)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}

func TestGetByIDError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating connection pool: %s", err)
	}
	defer mockConn.Close()

	mockConn.ExpectQuery("SELECT").
		WithArgs(int32(1)).
		WillReturnError(errors.New("test error"))

	eduRepo := internal.NewEducationPgxRepository(mockConn, slog.Default())

	_, err = eduRepo.GetByID(context.Background(), 1)
	if err == nil {
		t.Fatalf("expected error while getting all educations")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}
