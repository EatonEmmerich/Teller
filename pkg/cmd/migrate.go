package main

import (
	"Teller/pkg/catalogue"
	"Teller/pkg/db"
	"Teller/pkg/invoice"
	"context"
)

func main() {
	dbc, err := db.InitConnection()
	if err != nil {
		panic(err)
	}
	err = catalogue.SetupSchema(dbc, context.Background())
	if err != nil {
		panic(err)
	}
	err = invoice.SetupSchema(dbc, context.Background())
	if err != nil {
		panic(err)
	}
}
