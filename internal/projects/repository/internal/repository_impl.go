package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/logger"
)

type ProjectsPgxRepository struct {
	dbConn database.Connection
	logger logger.Logger
}
