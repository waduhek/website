package tplsvcint

import (
	"html/template"
	"path/filepath"

	"github.com/waduhek/website/internal/templates"
)

func NewTemplateServiceImpl(
	templateNameFileMap map[templates.TemplateName]string,
) (*TemplateServiceImpl, error) {
	templateBox := template.New("box")
	templateNameMapper := make(map[templates.TemplateName]string)

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

	return &TemplateServiceImpl{templateBox, templateNameMapper}, nil
}
