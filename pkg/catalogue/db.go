package catalogue

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed schema.sql
var schema string

func SetupSchema(dbc *sql.DB, ctx context.Context) error {
	_, err := dbc.ExecContext(ctx, schema)
	return err
}
