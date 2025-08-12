package handler

import "github.com/waduhek/website/internal"

// NewProjectsHandler creates a new project section handler.
func NewProjectsHandler(deps *internal.Dependencies) *ProjectsHandler {
	return &ProjectsHandler{
		logger:          deps.Logger,
		projRepo:        deps.ProjectsRepository,
		templateService: deps.TemplateService,
	}
}
