package handler

import (
	"github.com/waduhek/website/internal/education/repository"
	"github.com/waduhek/website/internal/telemetry"
	tplsvc "github.com/waduhek/website/internal/templates/service"
)

type EducationHandler struct {
	logger          telemetry.Logger
	eduRepo         repository.EducationRepository
	templateService tplsvc.TemplateService
}
