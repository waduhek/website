package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/waduhek/website/internal/experience/models"
)

func (r *ExperiencePgxRepository) InsertMany(
	ctx context.Context,
	experiences ...models.ExperienceInputModel,
) error {
	count, err := r.dbConn.CopyFrom(
		ctx,
		pgx.Identifier{"experience"},
		[]string{
			"title",
			"company_name",
			"start_date",
			"end_date",
			"is_current",
			"location",
			"description",
			"skills",
		},
		pgx.CopyFromSlice(len(experiences), func(i int) ([]any, error) {
			experience := experiences[i]
			retVal := []any{
				experience.Title,
				experience.CompanyName,
				experience.StartDate,
				experience.EndDate,
				experience.IsCurrent,
				experience.Location,
				experience.Description,
				experience.Skills,
			}

			return retVal, nil
		}),
	)

	r.logger.Debug("inserted multiple experiences", "count", count)

	return err
}
