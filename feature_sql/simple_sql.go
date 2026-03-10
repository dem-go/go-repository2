package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		ID SERIAL PRIMARY KEY,
		TITLE VARCHAR(200) NOT NULL,
		DESCRIPTION VARCHAR(1000),
		COMPLETED BOOLEAN NOT NULL,
		CREATED_AT TIMESTAMP NOT NULL,
		COMPLETED_AT TIMESTAMP
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)

	if err != nil {
		return err
	}

	return nil
}
