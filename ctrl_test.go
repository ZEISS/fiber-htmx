package htmx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	htmx "github.com/zeiss/fiber-htmx"
)

func TestNewTransactionControl(t *testing.T) {
	ctrl := htmx.NewTransactionController()
	require.NotNil(t, ctrl)
	require.Implements(t, (*htmx.TransactionController)(nil), ctrl)
}
