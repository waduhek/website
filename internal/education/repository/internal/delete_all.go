package internal

import "context"

func (r *EducationPgxRepository) DeleteAll(ctx context.Context) error {
	query := "TRUNCATE TABLE education;"
	_, err := r.dbConn.Exec(ctx, query)
	return err
}
