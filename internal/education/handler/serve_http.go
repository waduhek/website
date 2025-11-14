package handler

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"

	"github.com/waduhek/website/internal/education/models"
	"github.com/waduhek/website/internal/telemetry"
	"github.com/waduhek/website/internal/templates"
)

type pageData struct {
	Institute string
	Degree    string
	Major     string
	Grade     string
	Location  string
	Start     string
	End       string
}

func (h *EducationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := telemetry.ExtractContext(r.Context(), r.Header)

	spanCtx, span := telemetry.NewSpan(ctx, "GET /education")
	defer span.End()

	span.SetAttributes(
		semconv.HTTPRequestMethodGet,
		semconv.URLPath(r.URL.Path),
		semconv.URLScheme(r.URL.Scheme),
	)

	educations, err := h.getAndMapEducations(spanCtx)
	if err != nil {
		span.SetStatus(codes.Error, "error while getting and mapping education")
		span.SetAttributes(
			semconv.ErrorType(err),
			semconv.HTTPResponseStatusCode(500),
		)

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	template := h.templateService.GetTemplate(templates.Education)
	if template == nil {
		h.logger.ErrorContext(
			spanCtx,
			"could not get education section template",
		)

		span.SetStatus(codes.Error, "did not get template for education page")
		span.SetAttributes(semconv.HTTPResponseStatusCode(500))

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	err = template.Execute(w, educations)
	if err != nil {
		h.logger.ErrorContext(
			spanCtx,
			"error while executing education template",
			"err", err,
		)

		span.SetStatus(codes.Error, "error while executing education template")
		span.SetAttributes(
			semconv.ErrorType(err),
			semconv.HTTPResponseStatusCode(500),
		)

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	span.SetStatus(200, "education template executed")
	span.SetAttributes(semconv.HTTPResponseStatusCode(200))
}

func (h *EducationHandler) getAllEducations(
	ctx context.Context,
) ([]models.EducationOutputModel, error) {
	educations, err := h.eduRepo.GetAll(ctx)
	if err != nil {
		h.logger.ErrorContext(
			ctx,
			"error while getting all educations",
			"err", err,
		)
	}

	return educations, err
}

func (h *EducationHandler) mapEducations(
	educations []models.EducationOutputModel,
) []pageData {
	dateFormat := "January 2006"
	mappedData := make([]pageData, len(educations))

	for i, edu := range educations {
		mappedData[i] = pageData{
			Institute: edu.Institute,
			Degree:    edu.Degree,
			Major:     edu.Major,
			Grade:     edu.Grade,
			Location:  edu.Location,
			Start:     edu.StartDate.Format(dateFormat),
			End:       edu.EndDate.Format(dateFormat),
		}
	}

	return mappedData
}

func (h *EducationHandler) getAndMapEducations(
	ctx context.Context,
) ([]pageData, error) {
	educations, err := h.getAllEducations(ctx)
	if err != nil {
		return []pageData{}, err
	}

	mappedEdus := h.mapEducations(educations)
	return mappedEdus, nil
}
