package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/telemetry"
)

// NewEducationPgxRepository creates a new instance of [EducationPgxRepository].
func NewEducationPgxRepository(
	dbConn database.Connection,
	logger telemetry.Logger,
) *EducationPgxRepository {
	return &EducationPgxRepository{dbConn, logger}
}
