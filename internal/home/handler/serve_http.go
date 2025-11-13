package handlers

import (
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"

	"github.com/waduhek/website/internal/telemetry"
	"github.com/waduhek/website/internal/templates"
)

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := telemetry.ExtractContext(r.Context(), r.Header)

	spanCtx, span := telemetry.NewSpan(ctx, "GET /")
	defer span.End()

	span.SetAttributes(
		semconv.HTTPRequestMethodGet,
		semconv.URLPath(r.URL.Path),
		semconv.URLScheme(r.URL.Scheme),
	)

	template := h.templateService.GetTemplate(templates.Home)
	if template == nil {
		h.logger.ErrorContext(spanCtx, "did not get the template for home page")

		span.AddEvent("got nil template, returning status 500")
		span.SetStatus(codes.Error, "got nil template")
		span.SetAttributes(semconv.HTTPResponseStatusCode(500))

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	template.Execute(w, nil)

	span.SetAttributes(semconv.HTTPResponseStatusCode(200))
	span.SetStatus(codes.Ok, "home template found and executed successfully")
}
