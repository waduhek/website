package tplsvc

import (
	"html/template"

	"github.com/waduhek/website/internal/templates"
)

type TemplateService interface {
	// GetTemplate gets the template by the provided name. Returns nil if the
	// template was not found.
	GetTemplate(name templates.TemplateName) *template.Template
}
