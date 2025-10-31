package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/telemetry"
)

type EducationPgxRepository struct {
	dbConn database.Connection
	logger telemetry.Logger
}
