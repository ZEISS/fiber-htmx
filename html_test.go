package htmx

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElementA(t *testing.T) {
	e := Element("a")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<a></a>", e.String())
}

func ExampleElement_a() {
	_ = Element("a").Render(os.Stdout)
	// Output: <a></a>
}

func ExampleElement_div() {
	_ = Element("div").Render(os.Stdout)
	// Output: <div></div>
}
