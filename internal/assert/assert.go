package assert

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	htmx "github.com/zeiss/fiber-htmx"
)

// Equal asserts that the provided string matches the rendered Node output.
func Equal(t *testing.T, expected string, actual htmx.Node) {
	t.Helper()

	var b strings.Builder
	err := actual.Render(&b)
	require.NoError(t, err)

	if expected != b.String() {
		t.Errorf("expected %q, got %q", expected, b.String())
	}
}
