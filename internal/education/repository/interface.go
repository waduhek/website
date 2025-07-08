package repository

import (
	"context"

	"github.com/waduhek/website/internal/education/models"
)

type EducationRepository interface {
	// InsertOne inserts a single education entry.
	InsertOne(ctx context.Context, education *models.EducationInputModel) error

	// InsertMany inserts multiple education entries.
	InsertMany(
		ctx context.Context,
		educations ...models.EducationInputModel,
	) error

	// GetAll gets all education entries. Sorts the entries in descending order
	// of their start dates.
	GetAll(ctx context.Context) ([]models.EducationOutputModel, error)

	// GetByID gets an education by it's ID.
	GetByID(ctx context.Context, id int32) (*models.EducationOutputModel, error)

	// DeleteAll deletes all education entries.
	DeleteAll(ctx context.Context) error

	// DeleteByID deletes an education entry by it's ID.
	DeleteByID(ctx context.Context, id int32) error
}
