package internal

import "github.com/waduhek/website/internal/templates"

// TemplateNameFileMap is the map of the template name and the path to the
// template. Use this while building [templates.TemplateService].
var TemplateNameFileMap = map[templates.TemplateName]string{
	templates.Header:     "templates/common/header.html.tmpl",
	templates.Footer:     "templates/common/footer.html.tmpl",
	templates.Home:       "templates/home.html.tmpl",
	templates.Experience: "templates/experience.html.tmpl",
	templates.Education:  "templates/education.html.tmpl",
	templates.Projects:   "templates/projects.html.tmpl",
}
