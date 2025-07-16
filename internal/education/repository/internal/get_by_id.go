package internal

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/waduhek/website/internal/education/models"
)

func (r *EducationPgxRepository) GetByID(
	ctx context.Context,
	id int32,
) (*models.EducationOutputModel, error) {
	query := "SELECT " +
		"id, institute, degree, major, grade, location, start_date, end_date " +
		"FROM education WHERE id = $1;"

	rows, err := r.dbConn.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return pgx.CollectOneRow(
		rows,
		pgx.RowToAddrOfStructByName[models.EducationOutputModel],
	)
}
