package handler

import (
	"github.com/waduhek/website/internal/education/repository"
	"github.com/waduhek/website/internal/logger"
	"github.com/waduhek/website/internal/templates"
)

type EducationHandler struct {
	logger          logger.Logger
	eduRepo         repository.EducationRepository
	templateService *templates.TemplateService
}
