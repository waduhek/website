package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/waduhek/website/internal/experience/models"
)

func (r *ExperiencePgxRepository) InsertOne(
	ctx context.Context,
	experience *models.ExperienceInputModel,
) error {
	query := "INSERT INTO experience " +
		"(title, company_name, start_date, end_date, is_current, location, description, skills) " +
		"VALUES (@title, @company_name, @start_date, @end_date, @is_current, @location, @description, @skills);"
	args := pgx.NamedArgs{
		"title":        experience.Title,
		"company_name": experience.CompanyName,
		"start_date":   experience.StartDate,
		"end_date":     experience.EndDate,
		"is_current":   experience.IsCurrent,
		"location":     experience.Location,
		"description":  experience.Description,
		"skills":       experience.Skills,
	}

	_, err := r.dbConn.Exec(ctx, query, args)

	return err
}
