package handlers

import "github.com/waduhek/website/internal"

// NewHomeHandler creates a new instance of home page handler.
func NewHomeHandler(deps *internal.Dependencies) *HomeHandler {
	return &HomeHandler{
		logger:          deps.Logger,
		templateService: deps.TemplateService,
	}
}
