package repository

import (
	"github.com/waduhek/website/internal/database"
	expRepoInternal "github.com/waduhek/website/internal/experience/repository/internal"
	"github.com/waduhek/website/internal/telemetry"
)

// NewExperienceRepository creates a new [ExperienceRepository].
func NewExperienceRepository(
	dbConn database.Connection,
	logger telemetry.Logger,
) ExperienceRepository {
	return expRepoInternal.NewExperiencePgxRepository(dbConn, logger)
}
