package repository

import (
	"context"

	"github.com/waduhek/website/internal/projects/models"
)

type ProjectsRepository interface {
	// GetAll gets all the projects sorted in the descending order of their IDs.
	GetAll(ctx context.Context) ([]models.ProjectOutputModel, error)
}
