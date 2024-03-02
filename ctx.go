package htmx

import (
	"github.com/gofiber/fiber/v2"
)

// ContextCtx is a struct that contains the context of a htmx context.
type Context interface {
	// Context is the fiber.Ctx instance.
	Context(...*fiber.Ctx) *fiber.Ctx
	// Locals is a method that returns the local values.
	Locals(key any, value ...any) (val any)
	// Copy is a method that returns a new Ctx instance with the same properties.
	Copy() Ctx
	// Reset is a method that resets the local values.
	Reset()
}

var _ Context = (*Ctx)(nil)

// Ctx is a struct that contains the properties of a htmx context.
type Ctx struct {
	localValues map[any]any
	ctx         *fiber.Ctx
}

// Locals is a method that returns the local values.
func (c *Ctx) Locals(key any, value ...any) (val any) {
	if len(value) == 0 {
		return c.localValues[key]
	}

	c.localValues[key] = value[0]

	return value[0]
}

// Context is a method that returns the fiber.Ctx instance.
func (c *Ctx) Context(ctx ...*fiber.Ctx) *fiber.Ctx {
	if c.ctx == nil {
		c.ctx = &fiber.Ctx{}
	}

	if len(ctx) == 0 {
		return c.ctx
	}

	c.ctx = ctx[0]

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

// DefaultCtx is a function that returns a new Ctx instance.
func DefaultCtx() Ctx {
	return Ctx{
		localValues: make(map[any]any),
	}
}
