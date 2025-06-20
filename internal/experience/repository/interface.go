package repository

import (
	"context"

	"github.com/waduhek/website/internal/experience/models"
)

type ExperienceRepository interface {
	// InsertOne inserts a single experience entry.
	InsertOne(
		ctx context.Context,
		experience *models.ExperienceInputModel,
	) error

	// InsertMany inserts multiple experience entries.
	InsertMany(
		ctx context.Context,
		experiences ...models.ExperienceInputModel,
	) error

	// GetAll gets all the experience objects stored. Sorts the returned entries
	// by the experience start date in descending order.
	GetAll(ctx context.Context) ([]models.ExperienceOutputModel, error)

	// GetById gets an experience by it's ID.
	GetById(
		ctx context.Context,
		id int32,
	) (*models.ExperienceOutputModel, error)

	// DeleteAll deletes all the entries of experiences present.
	DeleteAll(ctx context.Context) error

	// DeleteById deletes a single experience by it's ID.
	DeleteById(ctx context.Context, id int32) error
}
