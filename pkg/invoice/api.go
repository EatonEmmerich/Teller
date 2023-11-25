package invoice

import "database/sql"

type Invoice struct {
	Id     int64
	IsPaid bool
}

func NewInvoice(dbc *sql.DB) (Invoice, error) {
	res := dbc.QueryRow("SELECT nextval('invoice_id_seq')")
	var invoiceID int64
	err := res.Scan(&invoiceID)
	if err != nil {
		return Invoice{}, err
	}
	_, err = dbc.Exec("INSERT INTO invoice VALUES ($1)", invoiceID)
	if err != nil {
		return Invoice{}, err
	}
	return Invoice{invoiceID, false}, nil
}

func (i *Invoice) AddItem(dbc *sql.DB, amount int64, description string) error {
	_, err := dbc.Exec("INSERT INTO invoice_item VALUES ($1, $2, $3)", i.Id, amount, description)
	return err
}

func (i *Invoice) Paid(dbc *sql.DB) error {
	_, err := dbc.Exec("UPDATE invoice SET paid=true WHERE id=$1", i.Id)
	i.IsPaid = true
	return err
}
