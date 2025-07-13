package handler

import "github.com/waduhek/website/internal"

// NewEducationHandler creates a new instance of EducationHandler.
func NewEducationHandler(deps *internal.Dependencies) *EducationHandler {
	return &EducationHandler{
		logger:          deps.Logger,
		eduRepo:         deps.EducationRepository,
		templateService: deps.TemplateService,
	}
}
