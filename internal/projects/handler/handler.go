package handler

import (
	"github.com/waduhek/website/internal/projects/repository"
	"github.com/waduhek/website/internal/telemetry"
	tplsvc "github.com/waduhek/website/internal/templates/service"
)

type ProjectsHandler struct {
	logger          telemetry.Logger
	projRepo        repository.ProjectsRepository
	templateService tplsvc.TemplateService
}
