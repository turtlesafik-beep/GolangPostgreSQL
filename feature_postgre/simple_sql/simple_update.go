package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRaw(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = TRUE
	WHERE id = 3;
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
