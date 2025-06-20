package templates_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/waduhek/website/internal/templates"
)

const templateFileName string = "temp-template.txt.tmpl"

func TestSuccessfulGetTemplate(t *testing.T) {
	templateService := setup(t)

	template := templateService.GetTemplate(templates.Home)
	if template == nil {
		t.Errorf("expected template to be non-nil value")
	}
}

func TestUnknownGetTemplate(t *testing.T) {
	templateService := setup(t)

	template := templateService.GetTemplate(templates.Header)
	if template != nil {
		t.Errorf("expected template to be nil value")
	}
}

func TestTemplateParseError(t *testing.T) {
	tempDir := t.TempDir()
	nonExistentTemplate := path.Join(tempDir, templateFileName)

	_, err := templates.NewTemplateService(map[templates.TemplateName]string{
		templates.Home: nonExistentTemplate,
	})
	if err == nil {
		t.Errorf("expected an error when parsing a non-existent template")
	}
}

func setup(t *testing.T) *templates.TemplateService {
	tempDir := t.TempDir()
	pathToTemplate := path.Join(tempDir, templateFileName)

	file, err := os.Create(pathToTemplate)
	if err != nil {
		t.Fatalf("error while creating template file: %v", err)
	}

	fmt.Fprint(file, "{{.Text}}")

	templateService, err := templates.NewTemplateService(
		map[templates.TemplateName]string{
			templates.Home: pathToTemplate,
		},
	)
	if err != nil {
		t.Fatalf("error while creating template service: %v\n", err)
	}

	return templateService
}
