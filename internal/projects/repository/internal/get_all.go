package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/waduhek/website/internal/projects/models"
)

func (r *ProjectsPgxRepository) GetAll(
	ctx context.Context,
) ([]models.ProjectOutputModel, error) {
	query := "SELECT id, name, public_url, repo_url, description, technologies " +
		"FROM project ORDER BY id DESC"

	rows, err := r.dbConn.Query(ctx, query)
	if err != nil {
		return []models.ProjectOutputModel{}, err
	}

	return pgx.CollectRows(
		rows,
		pgx.RowToStructByName[models.ProjectOutputModel],
	)
}
