package internal

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/waduhek/website/internal/education/models"
)

func (r *EducationPgxRepository) InsertOne(
	ctx context.Context,
	education *models.EducationInputModel,
) error {
	query := "INSERT INTO education " +
		"(institute, degree, major, grade, location, start_date, end_date) " +
		"VALUES (@institute, @degree, @major, @grade, @location, @start_date, @end_date);"
	args := pgx.NamedArgs{
		"institute":  education.Institute,
		"degree":     education.Degree,
		"major":      education.Major,
		"grade":      education.Grade,
		"location":   education.Location,
		"start_date": education.StartDate,
		"end_date":   education.EndDate,
	}

	_, err := r.dbConn.Exec(ctx, query, args)

	return err
}
