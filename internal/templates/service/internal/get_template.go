package tplsvcint

import (
	"html/template"

	"github.com/waduhek/website/internal/templates"
)

func (t *TemplateServiceImpl) GetTemplate(
	name templates.TemplateName,
) *template.Template {
	mappedName := t.templateNameMapper[name]
	return t.templateBox.Lookup(mappedName)
}
