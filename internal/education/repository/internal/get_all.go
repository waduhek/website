package internal

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/waduhek/website/internal/education/models"
)

func (r *EducationPgxRepository) GetAll(
	ctx context.Context,
) ([]models.EducationOutputModel, error) {
	query := "SELECT " +
		"id, institute, degree, major, grade, location, start_date, end_date " +
		"FROM education ORDER BY start_date DESC;"

	rows, err := r.dbConn.Query(ctx, query)
	if err != nil {
		return []models.EducationOutputModel{}, err
	}

	return pgx.CollectRows(
		rows,
		pgx.RowToStructByName[models.EducationOutputModel],
	)
}
