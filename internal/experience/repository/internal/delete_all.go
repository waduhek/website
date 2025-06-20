package internal

import "context"

func (r *ExperiencePgxRepository) DeleteAll(ctx context.Context) error {
	query := "TRUNCATE TABLE experience;"
	_, err := r.dbConn.Exec(ctx, query)
	return err
}
