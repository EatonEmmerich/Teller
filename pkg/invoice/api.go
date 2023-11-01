package invoice

import "database/sql"

type Invoice = int64

func NewInvoice(dbc *sql.DB) (Invoice, error) {
	res := dbc.QueryRow("SELECT nextval('invoice_id_seq')")
	var invoiceID int64
	err := res.Scan(&invoiceID)
	if err != nil {
		return 0, err
	}
	_, err = dbc.Exec("INSERT INTO invoice VALUES ($1)", invoiceID)
	if err != nil {
		return 0, err
	}
	return invoiceID, nil
}
