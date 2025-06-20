package internal

import "context"

func (r *ExperiencePgxRepository) DeleteById(
	ctx context.Context,
	id int32,
) error {
	query := "DELETE FROM experience WHERE id = $1;"
	_, err := r.dbConn.Exec(ctx, query, id)
	return err
}
