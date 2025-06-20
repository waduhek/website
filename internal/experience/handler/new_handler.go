package handlers

import "github.com/waduhek/website/internal"

// NewExperienceHandler creates a new experience handler.
func NewExperienceHandler(deps *internal.Dependencies) *ExperienceHandler {
	return &ExperienceHandler{
		logger:               deps.Logger,
		experienceRepository: deps.ExperienceRepository,
		templateService:      deps.TemplateService,
	}
}
