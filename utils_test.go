package htmx_test

import (
	"context"
	"testing"
	"time"

	htmx "github.com/zeiss/fiber-htmx"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAsBool(t *testing.T) {
	assert.True(t, htmx.AsBool("true"))
	assert.False(t, htmx.AsBool("false"))
	assert.True(t, htmx.AsBool("True"))
	assert.False(t, htmx.AsBool("False"))
}

func TestAsStr(t *testing.T) {
	assert.Equal(t, "true", htmx.AsStr(true))
	assert.Equal(t, "false", htmx.AsStr(false))
}

func TestIntAsString(t *testing.T) {
	assert.Equal(t, "1", htmx.IntAsString(1))
	assert.Equal(t, "0", htmx.IntAsString(0))
}

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

func TestFromContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resolver := func(ctx context.Context) (interface{}, interface{}, error) {
		return "key", "value", nil
	}

	c, err := htmx.FromContext(ctx, resolver)
	require.NoError(t, err)
	assert.NotNil(t, c)

	value := c.Locals("key")
	assert.Equal(t, "value", value)
}
