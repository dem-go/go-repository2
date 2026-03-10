package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateROW(ctx context.Context, conn *pgx.Conn) error {
	updateSQL := `
	UPDATE tasks
	SET completed = FALSE
	WHERE ID = 1
	`
	_, err := conn.Exec(ctx, updateSQL)
	if err != nil {
		panic(err)
	}
	return nil

}
