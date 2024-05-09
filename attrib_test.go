package htmx_test

import (
	"testing"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/internal/assert"
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
			a := htmx.Async()
			assert.Equal(t, test.want, a)
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
			a := htmx.AutoFocus()
			assert.Equal(t, test.want, a)
		})
	}
}
