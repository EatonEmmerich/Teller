package invoice_test

import (
	"Teller/pkg/db/testutils"
	"Teller/pkg/invoice"
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInvoice(t *testing.T) {
	dbc := testutils.ConnectForTesting(t)
	err := invoice.SetupSchema(dbc, context.Background())
	assert.NoError(t, err)

	type args struct {
		dbc *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    invoice.Invoice
		wantErr bool
	}{
		{
			"create new invoice",
			args{dbc},
			1,
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := invoice.NewInvoice(tt.args.dbc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewInvoice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateMultipleInvoices(t *testing.T) {
	dbc := testutils.ConnectForTesting(t)
	err := invoice.SetupSchema(dbc, context.Background())
	got, err := invoice.NewInvoice(dbc)
	assert.NoError(t, err)
	assert.Equal(t, invoice.Invoice(1), got)
	got, err = invoice.NewInvoice(dbc)
	assert.NoError(t, err)
	assert.Equal(t, invoice.Invoice(2), got)
}

func TestCreateMultipleInvoicesAsync(t *testing.T) {
	poolsize := 10
	dbc := testutils.ConnectForTestingPool(t, poolsize)
	err := invoice.SetupSchema(dbc[0], context.Background())
	assert.NoError(t, err)

	invoices := make(chan invoice.Invoice, 2)
	for x := 0; x < poolsize; x++ {
		go newInvoice(t, dbc[x], invoices)
	}

	finalList := make([]invoice.Invoice, poolsize)
	for x := 0; x < poolsize; x++ {
		finalList[x] = <-invoices
	}

	assert.ElementsMatch(t, finalList, []invoice.Invoice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func newInvoice(t *testing.T, dbc *sql.DB, cha chan invoice.Invoice) {
	got, err := invoice.NewInvoice(dbc)
	assert.NoError(t, err)
	cha <- got
}
