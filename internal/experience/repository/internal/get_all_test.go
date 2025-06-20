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

func TestGetAll(t *testing.T) {
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

	expectedExperiences := []models.ExperienceOutputModel{
		{
			ID:          1,
			Title:       "Test",
			CompanyName: "Test",
			StartDate:   time.Time{},
			EndDate:     time.Time{},
			IsCurrent:   true,
			Location:    "Test",
			Description: []string{"Test"},
			Skills:      []string{"Test"},
		},
	}

	mockConn.ExpectQuery("SELECT").
		WillReturnRows(rows)

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	gotExperiences, err := expRepo.GetAll(context.Background())
	if err != nil {
		t.Fatalf("error while getting all experiences: %s", err)
	}

	if !reflect.DeepEqual(expectedExperiences, gotExperiences) {
		t.Fatalf(
			"experiences do not match expected: %v, got: %v",
			expectedExperiences, gotExperiences,
		)
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}

func TestGetAllReturnsError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a mock pool: %s", err)
	}

	mockConn.ExpectQuery("SELECT").WillReturnError(errors.New("test error"))

	expRepo := internal.NewExperiencePgxRepository(mockConn, slog.Default())

	if _, err := expRepo.GetAll(context.Background()); err == nil {
		t.Fatalf("expected error while getting all experiences")
	}

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("expectation were not met: %s", err)
	}
}
