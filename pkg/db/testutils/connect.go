package testutils

import (
	"Teller/pkg/db"
	"database/sql"
	"github.com/stretchr/testify/require"
	"math/rand"
	"sync"
	"testing"
)

func ConnectForTesting(t *testing.T) *sql.DB {
	databaseName := generateDBName()
	createDB(t, databaseName)

	return connect(t, databaseName)
}

func ConnectForTestingPool(t *testing.T, poolSize int) []*sql.DB {
	databaseName := generateDBName()
	createDB(t, databaseName)

	pool := make([]*sql.DB, poolSize)
	for x := 0; x < poolSize; x++ {
		pool[x] = connect(t, databaseName)
	}
	return pool
}

func connect(t *testing.T, databaseName string) *sql.DB {
	dbc, err := db.Connect("localhost", 5432, "postgres", "postgres", databaseName)
	require.NoError(t, err)
	t.Cleanup(
		func() {
			dbc.Close()
		})
	return dbc
}

func createDB(t *testing.T, databaseName string) {
	dbc_root, err := db.Connect("localhost", 5432, "postgres", "postgres", "test")
	require.NoError(t, err)
	_, err = dbc_root.Exec("CREATE DATABASE " + databaseName)
	require.NoError(t, err)

	t.Cleanup(
		func() {
			_, err = dbc_root.Exec("DROP DATABASE " + databaseName)
			require.NoError(t, err)
			dbc_root.Close()
		})
}

var names = new(sync.Map)

func generateDBName() string {
	var resp string
	for a := 0; a < 12; a++ {
		resp = resp + string('a'+rune(rand.Intn(26)))
	}
	if _, ok := names.Load(resp); ok {
		return generateDBName()
	}
	names.Store(names, nil)
	return resp
}
