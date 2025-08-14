package templates

import (
	"html/template"
	"path/filepath"
)

type TemplateName = string

const (
	Header     TemplateName = "header"     // Header is the header section used in all templates.
	Footer     TemplateName = "footer"     // Footer is the footer used in all templates.
	Home       TemplateName = "home"       // Home is the template for the home page.
	Experience TemplateName = "experience" // Experience is the template for experience section.
	Education  TemplateName = "education"  // Education is the template for education section.
	Projects   TemplateName = "projects"   // Projects is the template for projects section.
)

// TemplateService is a storage container for all the templates registered.
type TemplateService struct {
	templateBox        *template.Template
	templateNameMapper map[TemplateName]string
}

// GetTemplate gets the template by the provided name. Returns nil if the
// template was not found.
func (t *TemplateService) GetTemplate(name TemplateName) *template.Template {
	mappedName := t.templateNameMapper[name]
	return t.templateBox.Lookup(mappedName)
}

// NewTemplateService parses all the defined templates and returns a service
// for interacting with the templates.
func NewTemplateService(
	templateNameFileMap map[TemplateName]string,
) (*TemplateService, error) {
	templateBox := template.New("box")
	templateNameMapper := make(map[TemplateName]string)

	for nameType, templatePath := range templateNameFileMap {
		_, err := templateBox.ParseFiles(templatePath)
		if err != nil {
			return nil, err
		}

		// This is the actual name of the template that is stored.
		fileName := filepath.Base(templatePath)
		// This mapping is required since `template.ParseFiles` names the
		// templates as the file name. I require something that allows me to
		// dynamically parse templates and associate them with some fixed string
		// so that there won't be magic strings where ever the templates are
		// being executed.
		templateNameMapper[nameType] = fileName
	}

	return &TemplateService{templateBox, templateNameMapper}, nil
}
