package repository

import (
	"context"

	"github.com/waduhek/website/internal/education/models"
)

type EducationRepository interface {
	// GetAll gets all education entries. Sorts the entries in descending order
	// of their start dates.
	GetAll(ctx context.Context) ([]models.EducationOutputModel, error)
}
