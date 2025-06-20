package handlers

import (
	"github.com/waduhek/website/internal/experience/repository"
	"github.com/waduhek/website/internal/logger"
	"github.com/waduhek/website/internal/templates"
)

type ExperienceHandler struct {
	logger               logger.Logger
	experienceRepository repository.ExperienceRepository
	templateService      *templates.TemplateService
}
