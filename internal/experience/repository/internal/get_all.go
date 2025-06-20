package internal

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/waduhek/website/internal/experience/models"
)

func (r *ExperiencePgxRepository) GetAll(
	ctx context.Context,
) ([]models.ExperienceOutputModel, error) {
	query := "SELECT " +
		"id, title, company_name, start_date, end_date, is_current, location, description, skills " +
		"FROM experience ORDER BY start_date DESC;"

	rows, err := r.dbConn.Query(ctx, query)
	if err != nil {
		return []models.ExperienceOutputModel{}, err
	}

	return pgx.CollectRows(
		rows,
		pgx.RowToStructByName[models.ExperienceOutputModel],
	)
}
