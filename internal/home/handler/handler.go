package handlers

import (
	"github.com/waduhek/website/internal/telemetry"
	tplsvc "github.com/waduhek/website/internal/templates/service"
)

type HomeHandler struct {
	logger          telemetry.Logger
	templateService tplsvc.TemplateService
}
