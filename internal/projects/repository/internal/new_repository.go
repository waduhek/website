package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/logger"
)

// NewProjectsPgxRepository creates a new ProjectsPgxRepository.
func NewProjectsPgxRepository(
	dbConn database.Connection,
	logger logger.Logger,
) *ProjectsPgxRepository {
	return &ProjectsPgxRepository{
		dbConn: dbConn,
		logger: logger,
	}
}
