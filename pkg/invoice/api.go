package invoice

import (
	"database/sql"
)

type Invoice struct {
	Id    int64
	Items []Item
}

type Item struct {
	Amount      int64
	Description string
}

func NewInvoice(dbc *sql.DB) (*Invoice, error) {
	res := dbc.QueryRow("SELECT nextval('invoice_id_seq')")
	var invoiceID int64
	err := res.Scan(&invoiceID)
	if err != nil {
		return &Invoice{}, err
	}
	_, err = dbc.Exec("INSERT INTO invoice VALUES ($1)", invoiceID)
	if err != nil {
		return &Invoice{}, err
	}
	return &Invoice{invoiceID, nil}, nil
}

func (i *Invoice) AddItem(dbc *sql.DB, amount int64, description string) error {
	_, err := dbc.Exec("INSERT INTO invoice_item VALUES ($1, $2, $3)", i.Id, amount, description)
	if err != nil {
		return err
	}
	i.Items = append(i.Items, Item{amount, description})
	return nil
}
