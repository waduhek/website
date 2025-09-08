package handler

import (
	"github.com/waduhek/website/internal/education/repository"
	"github.com/waduhek/website/internal/logger"
	tplsvc "github.com/waduhek/website/internal/templates/service"
)

type EducationHandler struct {
	logger          logger.Logger
	eduRepo         repository.EducationRepository
	templateService tplsvc.TemplateService
}
