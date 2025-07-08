package repository

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/education/repository/internal"
	"github.com/waduhek/website/internal/logger"
)

func NewEducationRepository(
	dbConn database.Connection,
	logger logger.Logger,
) EducationRepository {
	return internal.NewEducationPgxRepository(dbConn, logger)
}
