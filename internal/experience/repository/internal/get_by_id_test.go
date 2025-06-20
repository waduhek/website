package internal_test

import (
	"context"
	"errors"
	"log/slog"
	"reflect"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v4"

	"github.com/waduhek/website/internal/experience/models"
	"github.com/waduhek/website/internal/experience/repository/internal"
)

func TestGetById(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	rows := pgxmock.NewRows([]string{
		"id", "title", "company_name", "start_date", "end_date", "is_current",
		"location", "description", "skills",
	}).
		AddRow(
			1, "Test", "Test", time.Time{}, time.Time{}, true, "Test",
			[]string{"Test"}, []string{"Test"},
		)

	expectedExperience := &models.ExperienceOutputModel{
		ID:          1,
		Title:       "Test",
		CompanyName: "Test",
		StartDate:   time.Time{},
		EndDate:     time.Time{},
		IsCurrent:   true,
		Location:    "Test",
		Description: []string{"Test"},
		Skills:      []string{"Test"},
	}

	mockConn.ExpectQuery("SELECT").
		WithArgs(int32(1)).
		WillReturnRows(rows)

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	experience, err := expRepo.GetById(context.Background(), 1)
	if err != nil {
		t.Fatalf("error while getting experience: %s", err)
	}

	if !reflect.DeepEqual(expectedExperience, experience) {
		t.Fatalf(
			"expected experience to match. expected: %v got: %v",
			expectedExperience, experience,
		)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}

func TestGetByIdReturnsError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectQuery("SELECT").
		WithArgs(int32(1)).
		WillReturnError(errors.New("test error"))

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if _, err := expRepo.GetById(context.Background(), 1); err == nil {
		t.Fatalf("expected error to not be nil")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}
