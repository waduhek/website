package internal_test

import (
	"context"
	"errors"
	"log/slog"
	"reflect"
	"testing"

	"github.com/pashagolub/pgxmock/v4"

	"github.com/waduhek/website/internal/projects/models"
	"github.com/waduhek/website/internal/projects/repository/internal"
)

func TestGetAll(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating a new mock pool: %s", err)
	}

	rows := pgxmock.NewRows([]string{
		"id", "name", "public_url", "repo_url", "description", "technologies",
	}).
		AddRow(
			1, "Test", "https://example.com", "https://example.com",
			[]string{"Test"}, []string{"Test"},
		)

	expectedProjects := []models.ProjectOutputModel{
		{
			ID:           1,
			Name:         "Test",
			PublicURL:    "https://example.com",
			RepoURL:      "https://example.com",
			Description:  []string{"Test"},
			Technologies: []string{"Test"},
		},
	}

	mockConn.ExpectQuery("SELECT").WillReturnRows(rows)

	projRepo := internal.NewProjectsPgxRepository(mockConn, slog.Default())

	gotProjects, err := projRepo.GetAll(context.Background())
	if err != nil {
		t.Fatalf("error while getting projects: %s", err)
	}

	if !reflect.DeepEqual(expectedProjects, gotProjects) {
		t.Fatalf(
			"projects do not match. expected: %v, got: %v",
			expectedProjects, gotProjects,
		)
	}

	if err = mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}

func TestGetAllError(t *testing.T) {
	mockConn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error while creating mock connection: %s", err)
	}

	mockConn.ExpectQuery("SELECT").WillReturnError(errors.New("test error"))

	projRepo := internal.NewProjectsPgxRepository(mockConn, slog.Default())

	if _, err = projRepo.GetAll(context.Background()); err == nil {
		t.Fatalf("expected error while getting all projects")
	}

	if err = mockConn.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations were not met: %s", err)
	}
}
