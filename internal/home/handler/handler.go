package handlers

import (
	"github.com/waduhek/website/internal/logger"
	"github.com/waduhek/website/internal/templates"
)

type HomeHandler struct {
	logger          logger.Logger
	templateService *templates.TemplateService
}
