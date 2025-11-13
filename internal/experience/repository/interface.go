package repository

import (
	"context"

	"github.com/waduhek/website/internal/experience/models"
)

type ExperienceRepository interface {
	// GetAll gets all the experience objects stored. Sorts the returned entries
	// by the experience start date in descending order.
	GetAll(ctx context.Context) ([]models.ExperienceOutputModel, error)
}
