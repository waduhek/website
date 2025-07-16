package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/logger"
)

type EducationPgxRepository struct {
	dbConn database.Connection
	logger logger.Logger
}
