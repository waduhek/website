package handler

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.39.0"

	"github.com/waduhek/website/internal/projects/models"
	"github.com/waduhek/website/internal/telemetry"
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
	ctx := telemetry.ExtractContext(r.Context(), r.Header)

	spanCtx, span := telemetry.NewSpan(ctx, "GET /projects")
	defer span.End()

	span.SetAttributes(
		semconv.HTTPRequestMethodGet,
		semconv.URLPath(r.URL.Path),
		semconv.URLScheme(r.URL.Scheme),
	)

	projects, err := h.getAndMapProjects(spanCtx)
	if err != nil {
		span.SetStatus(codes.Error, "error while getting and mapping projects")
		span.SetAttributes(
			semconv.ErrorType(err),
			semconv.HTTPResponseStatusCode(500),
		)

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	template := h.templateService.GetTemplate(templates.Projects)
	if template == nil {
		h.logger.ErrorContext(spanCtx, "did not get template for projects page")

		span.SetStatus(codes.Error, "did not get template for projects page")
		span.SetAttributes(semconv.HTTPResponseStatusCode(500))

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	err = template.Execute(w, projects)
	if err != nil {
		h.logger.ErrorContext(
			spanCtx,
			"error while executing template",
			"err", err,
		)

		span.SetStatus(codes.Error, "error while executing projects template")
		span.SetAttributes(
			semconv.ErrorType(err),
			semconv.HTTPResponseStatusCode(500),
		)

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	span.SetStatus(codes.Ok, "projects template executed")
	span.SetAttributes(semconv.HTTPResponseStatusCode(200))
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
