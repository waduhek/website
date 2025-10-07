package internal

import (
	"github.com/waduhek/website/internal/database"
	"github.com/waduhek/website/internal/telemetry"
)

type ExperiencePgxRepository struct {
	dbConn database.Connection
	logger telemetry.Logger
}
