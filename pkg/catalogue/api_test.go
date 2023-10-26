package catalogue_test

import (
	"Teller/pkg/catalogue"
	"Teller/pkg/db/testutils"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupSchema(t *testing.T) {
	dbc := testutils.ConnectForTesting(t)

	err := catalogue.SetupSchema(dbc, context.Background())
	if err != nil {
		t.Errorf("SetupSchema() error = %v, want none", err)
	}
}

func Test_addItem(t *testing.T) {
	dbc := testutils.ConnectForTesting(t)
	err := catalogue.SetupSchema(dbc, context.Background())
	assert.NoError(t, err)
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"new catalogue item",
			args{name: "apple"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := catalogue.AddItem(dbc, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("addItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
