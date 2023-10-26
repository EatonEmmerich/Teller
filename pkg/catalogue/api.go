package catalogue

import "database/sql"

func AddItem(dbc *sql.DB, name string) error {
	stmt, err := dbc.Prepare("INSERT INTO items (name) VALUES ($1)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name)
	return err
}
