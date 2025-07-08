package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/logger"
)

// NewEducationPgxRepository creates a new instance of [EducationPgxRepository].
func NewEducationPgxRepository(
	dbConn database.Connection,
	logger logger.Logger,
) *EducationPgxRepository {
	return &EducationPgxRepository{dbConn, logger}
}
