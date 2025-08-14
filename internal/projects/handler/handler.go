package handler

import (
	"github.com/waduhek/website/internal/logger"
	"github.com/waduhek/website/internal/projects/repository"
	"github.com/waduhek/website/internal/templates"
)

type ProjectsHandler struct {
	logger          logger.Logger
	projRepo        repository.ProjectsRepository
	templateService *templates.TemplateService
}
