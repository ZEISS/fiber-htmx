package htmx_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	htmx "github.com/zeiss/fiber-htmx"
)

func Test_Async(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{
			name: "async",
			want: " async",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var b bytes.Buffer

			a := htmx.Async()
			assert.NotNil(t, a)

			err := a.Render(&b)
			require.NoError(t, err)
			assert.Equal(t, test.want, b.String())
		})
	}
}

func Test_AutoFocus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{
			name: "autofocus",
			want: " autofocus",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var b bytes.Buffer

			a := htmx.AutoFocus()
			assert.NotNil(t, a)

			err := a.Render(&b)
			require.NoError(t, err)
			assert.Equal(t, test.want, b.String())
		})
	}
}
