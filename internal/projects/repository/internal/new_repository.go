package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/telemetry"
)

// NewProjectsPgxRepository creates a new ProjectsPgxRepository.
func NewProjectsPgxRepository(
	dbConn database.Connection,
	logger telemetry.Logger,
) *ProjectsPgxRepository {
	return &ProjectsPgxRepository{
		dbConn: dbConn,
		logger: logger,
	}
}
