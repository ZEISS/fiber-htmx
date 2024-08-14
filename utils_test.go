package htmx_test

import (
	"testing"

	htmx "github.com/zeiss/fiber-htmx"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		in   []htmx.ClassNames
		out  htmx.ClassNames
	}{
		{
			name: "merge",
			in: []htmx.ClassNames{
				{"a": true, "b": false},
				{"b": true, "c": false},
			},
			out: htmx.ClassNames{"a": true, "b": true, "c": false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, htmx.Merge(tt.in...))
		})
	}
}
