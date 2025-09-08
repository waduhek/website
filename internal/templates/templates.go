package templates

type TemplateName = string

const (
	Header     TemplateName = "header"     // Header is the header section used in all templates.
	Footer     TemplateName = "footer"     // Footer is the footer used in all templates.
	Home       TemplateName = "home"       // Home is the template for the home page.
	Experience TemplateName = "experience" // Experience is the template for experience section.
	Education  TemplateName = "education"  // Education is the template for education section.
	Projects   TemplateName = "projects"   // Projects is the template for projects section.
)
