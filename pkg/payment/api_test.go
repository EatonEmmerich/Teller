package payment_test

import (
	"Teller/pkg/db/testutils"
	"Teller/pkg/invoice"
	"Teller/pkg/payment"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentOnInvoice(t *testing.T) {
	dbc := testutils.ConnectForTesting(t)
	err := payment.SetupSchema(dbc, context.Background())
	assert.NoError(t, err)
	invoice := &invoice.Invoice{
		1,
		[]invoice.Item{{1000, "description"}},
	}

	err = payment.Pay(
		dbc,
		invoice,
		500,
	)
	assert.NoError(t, err)
	got, err := payment.OutstandingBalance(dbc, invoice)
	assert.NoError(t, err)
	assert.Equal(t, int64(500), got)
}
