package htmx

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ContextCtx is a struct that contains the context of a htmx context.
type Context interface {
	// Context is the fiber.Ctx instance.
	Context() *fiber.Ctx
	// Locals is a method that returns the local values.
	Locals(key any, value ...any) (val any)
	// Copy is a method that returns a new Ctx instance with the same properties.
	Copy() Ctx
	// Reset is a method that resets the local values.
	Reset()
}

var (
	_ Context         = (*Ctx)(nil)
	_ context.Context = (*Ctx)(nil)
)

// Ctx is a struct that contains the properties of a htmx context.
type Ctx struct {
	localValues map[any]any
	ctx         *fiber.Ctx
	err         error
	done        chan struct{}
}

// Value is a method that returns the value of the provided key.
func (c *Ctx) Value(key interface{}) interface{} {
	return c.Locals(key)
}

// Deadline is a method that returns the deadline of the context.
func (c *Ctx) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Done is a method that returns a channel that is closed when the context is done.
func (c *Ctx) Done() <-chan struct{} {
	return c.done
}

// Err is a method that returns the error of the context.
func (c *Ctx) Err() error {
	return c.err
}

// Locals is a method that returns the local values.
func (c *Ctx) Locals(key any, value ...any) (val any) {
	if len(value) == 0 {
		return c.localValues[key]
	}

	c.localValues[key] = value[0]

	return value[0]
}

// Locals is a method that returns the local values.
func Locals[V any](c Context, key any, value ...V) V {
	var v V
	var ok bool

	if len(value) == 0 {
		v, ok = c.Locals(key).(V)
	} else {
		v, ok = c.Locals(key, value[0]).(V)
	}

	if !ok {
		return v // return zero of type V
	}

	return v
}

// Context is a method that returns the fiber.Ctx instance.
func (c *Ctx) Context() *fiber.Ctx {
	return c.ctx
}

// Reset is a method that resets the local values.
func (c *Ctx) Reset() {
	c.localValues = make(map[any]any)
	c.ctx = nil
}

// Copy is a method that returns a new Ctx instance with the same properties.
func (c *Ctx) Copy() Ctx {
	return Ctx{
		localValues: c.localValues,
		ctx:         c.ctx,
	}
}

// FromContext is a function that returns the Ctx instance from the provided context.
func FromContext(c *fiber.Ctx) *Ctx {
	return &Ctx{
		localValues: make(map[any]any),
		ctx:         c,
	}
}
