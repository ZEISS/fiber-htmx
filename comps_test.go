package htmx_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	htmx "github.com/zeiss/fiber-htmx"
)

func BenchmarkElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := htmx.Element("div").Render(io.Discard)
		assert.NoError(b, err)
	}
}

func ExampleElement_tag() {
	_ = htmx.Element("div").Render(os.Stdout)
	// Output: <div></div>
}

func ExampleAttribute_bool() {
	_ = htmx.Attribute("disabled", htmx.AsStr(true)).Render(os.Stdout)
	// Output: disabled="true"
}

func ExampleAttribute_name_value() {
	_ = htmx.Attribute("href", "/").Render(os.Stdout)
	// Output: href="/"
}

func BenchmarkAttribute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		htmx.Attribute("href", "/")
		err := htmx.Attribute("href", "/").Render(io.Discard)
		assert.NoError(b, err)
	}
}

func TestNodeFunc(t *testing.T) {
	t.Run("implements fmt.String", func(t *testing.T) {
		fn := htmx.NodeFunc(func(w io.Writer) error {
			_, _ = w.Write([]byte("hello"))
			return nil
		})

		assert.Implements(t, (*fmt.Stringer)(nil), fn)
		assert.Equal(t, "hello", fmt.Sprint(fn))
	})
}
