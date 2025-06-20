package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/waduhek/website/internal/experience/models"
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
	data, err := h.getAndMapExperiencesToPageData(r.Context())
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	template := h.templateService.GetTemplate(templates.Experience)
	if template == nil {
		h.logger.Error("did not get the template for experience page")

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}

	err = template.Execute(w, data)
	if err != nil {
		h.logger.Error("error while executing template", "err", err)

		w.WriteHeader(500)
		fmt.Fprint(w, "oops an error occurred")

		return
	}
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
