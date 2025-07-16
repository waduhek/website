package internal

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/waduhek/website/internal/education/models"
)

func (r *EducationPgxRepository) InsertMany(
	ctx context.Context,
	educations ...models.EducationInputModel,
) error {
	count, err := r.dbConn.CopyFrom(
		ctx,
		pgx.Identifier{"education"},
		[]string{
			"institute",
			"degree",
			"major",
			"grade",
			"location",
			"start_date",
			"end_date",
		},
		pgx.CopyFromSlice(len(educations), func(i int) ([]any, error) {
			education := educations[i]

			retVal := []any{
				education.Institute,
				education.Degree,
				education.Major,
				education.Grade,
				education.Location,
				education.StartDate,
				education.EndDate,
			}

			return retVal, nil
		}),
	)

	r.logger.InfoContext(ctx, "inserted multiple educations", "count", count)

	return err
}
