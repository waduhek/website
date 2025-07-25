package internal

import (
	"context"
	"os"

	"github.com/waduhek/website/internal/database"
	eduRepo "github.com/waduhek/website/internal/education/repository"
	experienceRepository "github.com/waduhek/website/internal/experience/repository"
	"github.com/waduhek/website/internal/logger"
	"github.com/waduhek/website/internal/templates"
)

// Dependencies is a structure that is used to store all the dependencies that
// may be required.
type Dependencies struct {
	Logger               logger.Logger
	DbConn               database.Connection
	ExperienceRepository experienceRepository.ExperienceRepository
	EducationRepository  eduRepo.EducationRepository
	TemplateService      *templates.TemplateService
}

// BuildDependencies builds all the dependencies of the service.
func BuildDependencies(templateNamePathMap map[string]string) *Dependencies {
	logger := logger.NewLogger()
	dbConn, err := database.Connect(context.Background())
	if err != nil {
		logger.Error("error while connecting to the database", "err", err)
		os.Exit(1)
	}
	eduRepo := eduRepo.NewEducationRepository(dbConn, logger)
	expRepo := experienceRepository.NewExperienceRepository(dbConn, logger)
	templateService, err := templates.NewTemplateService(templateNamePathMap)
	if err != nil {
		logger.Error("error while building template service", "err", err)
		os.Exit(1)
	}

	return &Dependencies{
		Logger:               logger,
		DbConn:               dbConn,
		EducationRepository:  eduRepo,
		ExperienceRepository: expRepo,
		TemplateService:      templateService,
	}
}
