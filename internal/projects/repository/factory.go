package repository

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/projects/repository/internal"
	"github.com/waduhek/website/internal/telemetry"
)

// NewProjectsRepository creates a new implementation of ProjectsRepository.
func NewProjectsRepository(
	dbConn database.Connection,
	logger telemetry.Logger,
) ProjectsRepository {
	return internal.NewProjectsPgxRepository(dbConn, logger)
}
