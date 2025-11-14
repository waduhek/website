package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"

	"github.com/waduhek/website/internal/education/models"
	"github.com/waduhek/website/internal/telemetry"
)

func (r *EducationPgxRepository) GetAll(
	ctx context.Context,
) ([]models.EducationOutputModel, error) {
	spanCtx, span := telemetry.NewSpan(ctx, "get education repo")
	defer span.End()

	query := "SELECT " +
		"id, institute, degree, major, grade, location, start_date, end_date " +
		"FROM education ORDER BY start_date DESC;"

	span.SetAttributes(
		semconv.DBCollectionName("website.education"),
		semconv.DBQueryText(query),
		semconv.DBOperationName("SELECT"),
	)

	rows, err := r.dbConn.Query(spanCtx, query)
	if err != nil {
		span.SetStatus(codes.Error, "query returned error")
		span.SetAttributes(semconv.ErrorType(err))

		return []models.EducationOutputModel{}, err
	}

	collectedRows, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[models.EducationOutputModel],
	)
	if err != nil {
		span.SetStatus(codes.Error, "error while collecting rows")
		span.SetAttributes(semconv.ErrorType(err))

		return []models.EducationOutputModel{}, err
	}

	span.SetStatus(codes.Ok, "got all education entries")
	return collectedRows, nil
}
