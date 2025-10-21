package handlers

import (
	"github.com/waduhek/website/internal/experience/repository"
	"github.com/waduhek/website/internal/telemetry"
	tplsvc "github.com/waduhek/website/internal/templates/service"
)

type ExperienceHandler struct {
	logger               telemetry.Logger
	experienceRepository repository.ExperienceRepository
	templateService      tplsvc.TemplateService
}
