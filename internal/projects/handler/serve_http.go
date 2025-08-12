package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/waduhek/website/internal/projects/models"
	"github.com/waduhek/website/internal/templates"
)

type pageData struct {
	Name         string
	PublicURL    string
	RepoURL      string
	Description  []string
	Technologies []string
}

func (h *ProjectsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projects, err := h.getAndMapProjects(ctx)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	template := h.templateService.GetTemplate(templates.Projects)
	if template == nil {
		h.logger.ErrorContext(ctx, "did not get template for projects page")

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	err = template.Execute(w, projects)
	if err != nil {
		h.logger.ErrorContext(ctx, "error while executing template", "err", err)

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}
}

func (h *ProjectsHandler) getAndMapProjects(
	ctx context.Context,
) ([]pageData, error) {
	projects, err := h.getProjects(ctx)
	if err != nil {
		return []pageData{}, err
	}

	return h.mapProjects(projects), nil
}

func (h *ProjectsHandler) getProjects(
	ctx context.Context,
) ([]models.ProjectOutputModel, error) {
	data, err := h.projRepo.GetAll(ctx)
	if err != nil {
		h.logger.ErrorContext(
			ctx,
			"error while getting projects from database",
			"err", err,
		)

		return []models.ProjectOutputModel{}, err
	}

	return data, nil
}

func (h *ProjectsHandler) mapProjects(
	projects []models.ProjectOutputModel,
) []pageData {
	mappedProjects := make([]pageData, len(projects))

	for i, project := range projects {
		mappedProjects[i] = pageData{
			Name:         project.Name,
			PublicURL:    project.PublicURL,
			RepoURL:      project.RepoURL,
			Description:  project.Description,
			Technologies: project.Technologies,
		}
	}

	return mappedProjects
}
