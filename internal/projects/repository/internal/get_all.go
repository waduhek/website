package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"

	"github.com/waduhek/website/internal/projects/models"
	"github.com/waduhek/website/internal/telemetry"
)

func (r *ProjectsPgxRepository) GetAll(
	ctx context.Context,
) ([]models.ProjectOutputModel, error) {
	spanCtx, span := telemetry.NewSpan(ctx, "get projects repo")
	defer span.End()

	query := "SELECT id, name, public_url, repo_url, description, technologies " +
		"FROM project ORDER BY id DESC"

	span.SetAttributes(
		semconv.DBCollectionName("website.project"),
		semconv.DBQueryText(query),
		semconv.DBOperationName("SELECT"),
	)

	rows, err := r.dbConn.Query(spanCtx, query)
	if err != nil {
		span.SetStatus(codes.Error, "query returned error")
		span.SetAttributes(semconv.ErrorType(err))

		return []models.ProjectOutputModel{}, err
	}

	collectedRows, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[models.ProjectOutputModel],
	)
	if err != nil {
		span.SetStatus(codes.Error, "error while collecting rows")
		span.SetAttributes(semconv.ErrorType(err))

		return []models.ProjectOutputModel{}, err
	}

	span.SetStatus(codes.Ok, "got all projects")
	return collectedRows, nil
}
