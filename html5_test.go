package htmx_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	htmx "github.com/zeiss/fiber-htmx"
)

func Benchmark_HTML5_Render(b *testing.B) {
	ctx := htmx.DefaultCtx()

	nodes := []htmx.Node{}

	for i := 0; i < 10000; i++ {
		nodes = append(nodes, htmx.Element("div"))
	}

	for i := 0; i < b.N; i++ {
		err := htmx.HTML5(
			ctx,
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
