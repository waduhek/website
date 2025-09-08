package tplsvcint

import (
	"html/template"

	"github.com/waduhek/website/internal/templates"
)

type TemplateServiceImpl struct {
	templateBox        *template.Template
	templateNameMapper map[templates.TemplateName]string
}
