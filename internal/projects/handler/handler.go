package handler

import (
	"github.com/waduhek/website/internal/logger"
	"github.com/waduhek/website/internal/projects/repository"
	tplsvc "github.com/waduhek/website/internal/templates/service"
)

type ProjectsHandler struct {
	logger          logger.Logger
	projRepo        repository.ProjectsRepository
	templateService tplsvc.TemplateService
}
