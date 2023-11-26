package payment

import (
	"Teller/pkg/invoice"
	"database/sql"
)

func Pay(dbc *sql.DB, invoice *invoice.Invoice, amount int64) error {
	res := dbc.QueryRow("SELECT nextval('payment_id_seq')")
	var paymentId int64
	err := res.Scan(&paymentId)
	if err != nil {
		return err
	}

	_, err = dbc.Exec("INSERT INTO payment VALUES ($1, $2, $3)", paymentId, amount, invoice.Id)
	return err
}

func OutstandingBalance(dbc *sql.DB, invoice *invoice.Invoice) (int64, error) {
	res, err := dbc.Query("SELECT amount FROM payment WHERE invoice=$1", invoice.Id)
	if err != nil {
		return 0, err
	}
	total := int64(0)
	for res.Next() {
		var amount int64
		err := res.Scan(&amount)
		if err != nil {
			return 0, err
		}
		total += amount
	}
	return total, nil
}
