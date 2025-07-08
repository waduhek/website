package internal

import "context"

func (r *EducationPgxRepository) DeleteByID(
	ctx context.Context,
	id int32,
) error {
	query := "DELETE FROM education WHERE id = $1"
	_, err := r.dbConn.Exec(ctx, query, id)
	return err
}
