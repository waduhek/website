package internal

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/waduhek/website/internal/experience/models"
)

func (r *ExperiencePgxRepository) GetById(
	ctx context.Context,
	id int32,
) (*models.ExperienceOutputModel, error) {
	query := "SELECT " +
		"(id, title, company_name, start_date, end_date, is_current, location, description, skills) " +
		"FROM experience WHERE id = $1;"

	rows, err := r.dbConn.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToAddrOfStructByName[models.ExperienceOutputModel],
	)
}
