package htmx_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	htmx "github.com/zeiss/fiber-htmx"
)

func Benchmark_HTML5_Render(b *testing.B) {
	nodes := []htmx.Node{}

	for i := 0; i < 10000; i++ {
		nodes = append(nodes, htmx.Element("div"))
	}

	for i := 0; i < b.N; i++ {
		err := htmx.HTML5(
			htmx.HTML5Props{
				Title:       "Hello, World!",
				Description: "An example of an HTML5 document.",
				Language:    "en",
				Head: []htmx.Node{
					htmx.Meta(htmx.Charset("utf-8")),
					htmx.Meta(htmx.Name("viewport"), htmx.Content("width=device-width, initial-scale=1")),
					htmx.TitleElement(htmx.Text("Hello, World!")),
				},
			}, nodes...).Render(io.Discard)
		assert.NoError(b, err)
	}
}

func Test_ClassNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		dst  htmx.ClassNames
		src  []htmx.ClassNames
		want htmx.ClassNames
	}{
		{
			name: "empty",
			dst:  htmx.ClassNames{},
			src:  []htmx.ClassNames{},
			want: htmx.ClassNames{},
		},
		{
			name: "single",
			dst:  htmx.ClassNames{},
			src:  []htmx.ClassNames{{"foo": true}},
			want: htmx.ClassNames{"foo": true},
		},
		{
			name: "multiple",
			dst:  htmx.ClassNames{},
			src:  []htmx.ClassNames{{"foo": true}, {"bar": true}},
			want: htmx.ClassNames{"foo": true, "bar": true},
		},
		{
			name: "overwrite",
			dst:  htmx.ClassNames{"foo": true},
			src:  []htmx.ClassNames{{"foo": false}},
			want: htmx.ClassNames{"foo": false},
		},
		{
			name: "merge",
			dst:  htmx.ClassNames{"foo": true},
			src:  []htmx.ClassNames{{"bar": true}},
			want: htmx.ClassNames{"foo": true, "bar": true},
		},
		{
			name: "merge overwrite",
			dst:  htmx.ClassNames{"foo": true},
			src:  []htmx.ClassNames{{"foo": false}},
			want: htmx.ClassNames{"foo": false},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.dst.Copy(test.src...)
			assert.Equal(t, test.want, test.dst)
		})
	}
}

func Test_ClassNames_Render(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		c    htmx.ClassNames
		want string
	}{
		{
			name: "empty",
			c:    htmx.ClassNames{},
			want: " class=\"\"",
		},
		{
			name: "single",
			c:    htmx.ClassNames{"foo": true},
			want: " class=\"foo\"",
		},
		{
			name: "multiple",
			c:    htmx.ClassNames{"foo": true, "bar": true},
			want: " class=\"bar foo\"",
		},
		{
			name: "overwrite",
			c:    htmx.ClassNames{"foo": false},
			want: " class=\"\"",
		},
		{
			name: "merge",
			c:    htmx.ClassNames{"foo": true},
			want: " class=\"foo\"",
		},
		{
			name: "merge overwrite",
			c:    htmx.ClassNames{"foo": true},
			want: " class=\"foo\"",
		},
		{
			name: "trim space",
			c:    htmx.ClassNames{" foo ": true},
			want: " class=\"foo\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var b bytes.Buffer
			err := test.c.Render(&b)

			require.NoError(t, err)
			assert.Equal(t, test.want, b.String())
		})
	}
}
