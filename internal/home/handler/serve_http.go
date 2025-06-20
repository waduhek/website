package handlers

import (
	"fmt"
	"net/http"

	"github.com/waduhek/website/internal/templates"
)

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	template := h.templateService.GetTemplate(templates.Home)
	if template == nil {
		h.logger.Error("did not get the template for home page")

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")
	}

	template.Execute(w, nil)
}
