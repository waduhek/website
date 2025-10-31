package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/telemetry"
)

// NewExperiencePgxRepository creates a new instance of
// [ExperiencePgxRepository].
func NewExperiencePgxRepository(
	dbConn database.Connection,
	logger telemetry.Logger,
) *ExperiencePgxRepository {
	return &ExperiencePgxRepository{dbConn, logger}
}
