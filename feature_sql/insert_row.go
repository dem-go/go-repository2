package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InserRow(ctx context.Context, conn *pgx.Conn) error {
	quarySQL := `
	INSERT INTO tasks (TITLE, DESCRIPTION, COMPLETED, CREATED_AT, COMPLETED_AT)
	VALUES ('DERGAU', 'NU TUT NE DOBAVIT NE UBAVIT', TRUE, '2026-06-03 20:20:05', NULL);
	`

	_, err := conn.Exec(ctx, quarySQL)
	if err != nil {
		panic(err)
	}

	return nil

}
