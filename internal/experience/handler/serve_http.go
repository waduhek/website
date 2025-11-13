package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"

	"github.com/waduhek/website/internal/experience/models"
	"github.com/waduhek/website/internal/telemetry"
	"github.com/waduhek/website/internal/templates"
)

type pageData struct {
	Title       string
	Company     string
	Start       string
	End         string
	IsCurrent   bool
	Location    string
	Description []string
	Skills      string
}

func (h *ExperienceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := telemetry.ExtractContext(r.Context(), r.Header)

	spanCtx, span := telemetry.NewSpan(ctx, "GET /experiences")
	defer span.End()

	span.SetAttributes(
		semconv.HTTPRequestMethodGet,
		semconv.URLPath(r.URL.Path),
		semconv.URLScheme(r.URL.Scheme),
	)

	data, err := h.getAndMapExperiencesToPageData(spanCtx)
	if err != nil {
		span.SetStatus(codes.Error, "error while getting and mapping experience")
		span.SetAttributes(
			semconv.ErrorType(err),
			semconv.HTTPResponseStatusCode(500),
		)

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	template := h.templateService.GetTemplate(templates.Experience)
	if template == nil {
		h.logger.ErrorContext(
			spanCtx,
			"did not get the template for experience page",
		)

		span.SetStatus(codes.Error, "did not get template for experience page")
		span.SetAttributes(semconv.HTTPResponseStatusCode(500))

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	err = template.Execute(w, data)
	if err != nil {
		h.logger.ErrorContext(
			spanCtx,
			"error while executing template",
			"err", err,
		)

		span.SetStatus(codes.Error, "error while executing experience template")
		span.SetAttributes(
			semconv.ErrorType(err),
			semconv.HTTPResponseStatusCode(500),
		)

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	span.SetStatus(codes.Ok, "experience page template executed")
	span.SetAttributes(semconv.HTTPResponseStatusCode(200))
}

func (h *ExperienceHandler) getExperiences(
	ctx context.Context,
) ([]models.ExperienceOutputModel, error) {
	data, err := h.experienceRepository.GetAll(ctx)
	if err != nil {
		h.logger.ErrorContext(
			ctx,
			"error while getting experience data",
			"err", err,
		)

		return []models.ExperienceOutputModel{}, err
	}

	return data, nil
}

func (h *ExperienceHandler) mapExperiencesToPageData(
	experiences []models.ExperienceOutputModel,
) []pageData {
	experienceDateFormat := "January 2006"
	data := make([]pageData, len(experiences))

	for i, experience := range experiences {
		skills := strings.Join(experience.Skills, ", ")

		data[i] = pageData{
			Title:       experience.Title,
			Company:     experience.CompanyName,
			Start:       experience.StartDate.Format(experienceDateFormat),
			End:         experience.EndDate.Format(experienceDateFormat),
			IsCurrent:   experience.IsCurrent,
			Location:    experience.Location,
			Description: experience.Description,
			Skills:      skills,
		}
	}

	return data
}

func (h *ExperienceHandler) getAndMapExperiencesToPageData(
	ctx context.Context,
) ([]pageData, error) {
	experiences, err := h.getExperiences(ctx)
	if err != nil {
		return []pageData{}, err
	}

	return h.mapExperiencesToPageData(experiences), nil
}
