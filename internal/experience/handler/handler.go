package handlers

import (
	"github.com/waduhek/website/internal/experience/repository"
	"github.com/waduhek/website/internal/logger"
	tplsvc "github.com/waduhek/website/internal/templates/service"
)

type ExperienceHandler struct {
	logger               logger.Logger
	experienceRepository repository.ExperienceRepository
	templateService      tplsvc.TemplateService
}
