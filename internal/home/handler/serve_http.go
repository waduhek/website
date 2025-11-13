package handlers

import (
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/codes"

	"github.com/waduhek/website/internal/telemetry"
	"github.com/waduhek/website/internal/templates"
)

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := telemetry.ExtractContext(r.Context(), r.Header)

	_, span := telemetry.NewSpan(ctx, "GET /")
	defer span.End()

	template := h.templateService.GetTemplate(templates.Home)
	if template == nil {
		h.logger.ErrorContext(ctx, "did not get the template for home page")

		span.AddEvent("got nil template, returning status 500")
		span.SetStatus(codes.Error, "got nil template")

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	template.Execute(w, nil)

	span.SetStatus(codes.Ok, "home template found and executed successfully")
}
