package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/logger"
)

// NewExperiencePgxRepository creates a new instance of
// [ExperiencePgxRepository].
func NewExperiencePgxRepository(
	dbConn database.Connection,
	logger logger.Logger,
) *ExperiencePgxRepository {
	return &ExperiencePgxRepository{dbConn, logger}
}
