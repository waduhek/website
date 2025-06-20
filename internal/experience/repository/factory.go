package repository

import (
	"github.com/waduhek/website/internal/database"
	expRepoInternal "github.com/waduhek/website/internal/experience/repository/internal"
	"github.com/waduhek/website/internal/logger"
)

// NewExperienceRepository creates a new [ExperienceRepository].
func NewExperienceRepository(
	dbConn database.Connection,
	logger logger.Logger,
) ExperienceRepository {
	return expRepoInternal.NewExperiencePgxRepository(dbConn, logger)
}
