package tplsvc

import (
	"github.com/waduhek/website/internal/templates"
	tplsvcint "github.com/waduhek/website/internal/templates/service/internal"
)

// NewTemplateService parses all the defined templates and returns a service
// for interacting with the templates.
func NewTemplateService(
	templateNameFileMap map[templates.TemplateName]string,
) (TemplateService, error) {
	return tplsvcint.NewTemplateServiceImpl(templateNameFileMap)
}
