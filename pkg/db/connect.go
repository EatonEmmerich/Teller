package db

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
)

var host = flag.String("databaseHost", "localhost", "the host dns")
var port = flag.Int("databasePort", 5432, "The database network port")
var database = flag.String("databaseName", "teller", "The name of the db")
var user = flag.String("databaseUser", "postgres", "The name of the db user for the app")
var password = flag.String("databasePassword", "postgres", "The password of the db user specified in 'user' parameter")

func InitConnection() (*sql.DB, error) {
	return Connect(*host, *port, *user, *password, *database)
}

func Connect(host string, port int, user string, password string, database string) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, database))
}

type Conn interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)

	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}
