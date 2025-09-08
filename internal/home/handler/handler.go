package handlers

import (
	"github.com/waduhek/website/internal/logger"
	tplsvc "github.com/waduhek/website/internal/templates/service"
)

type HomeHandler struct {
	logger          logger.Logger
	templateService tplsvc.TemplateService
}
