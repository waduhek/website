package repository

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/education/repository/internal"
	"github.com/waduhek/website/internal/telemetry"
)

func NewEducationRepository(
	dbConn database.Connection,
	logger telemetry.Logger,
) EducationRepository {
	return internal.NewEducationPgxRepository(dbConn, logger)
}
