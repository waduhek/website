package repository

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/logger"
	"github.com/waduhek/website/internal/projects/repository/internal"
)

// NewProjectsRepository creates a new implementation of ProjectsRepository.
func NewProjectsRepository(
	dbConn database.Connection,
	logger logger.Logger,
) ProjectsRepository {
	return internal.NewProjectsPgxRepository(dbConn, logger)
}
