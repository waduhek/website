package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"

	"github.com/waduhek/website/internal/experience/models"
	"github.com/waduhek/website/internal/telemetry"
)

func (r *ExperiencePgxRepository) GetAll(
	ctx context.Context,
) ([]models.ExperienceOutputModel, error) {
	spanCtx, span := telemetry.NewSpan(ctx, "get experiences repo")
	defer span.End()

	query := "SELECT " +
		"id, title, company_name, start_date, end_date, is_current, location, description, skills " +
		"FROM experience ORDER BY start_date DESC;"

	span.SetAttributes(
		semconv.DBCollectionName("website.experience"),
		semconv.DBQueryText(query),
		semconv.DBOperationName("SELECT"),
	)

	rows, err := r.dbConn.Query(spanCtx, query)
	if err != nil {
		span.SetStatus(codes.Error, "query returned error")
		span.SetAttributes(semconv.ErrorType(err))

		return []models.ExperienceOutputModel{}, err
	}

	collectedRows, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[models.ExperienceOutputModel],
	)
	if err != nil {
		span.SetStatus(codes.Error, "error while collecting rows")
		span.SetAttributes(semconv.ErrorType(err))

		return []models.ExperienceOutputModel{}, err
	}

	span.SetStatus(codes.Ok, "experiences fetched")
	return collectedRows, nil
}
