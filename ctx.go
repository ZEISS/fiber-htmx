package htmx

import (
	"github.com/gofiber/fiber/v2"
)

// Ctx is a struct that contains the context of a htmx context.
type Ctx interface {
	// Context is the fiber.Ctx instance.
	Context(...*fiber.Ctx) *fiber.Ctx
	// Locals ...
	Locals(key any, value ...any) (val any)
	// Reset ...
	Reset()
}

// DefaultCtx is the default implementation of the Ctx interface.
type DefaultCtx struct {
	localValues map[any]any
	ctx         *fiber.Ctx
}

// Locals is a method that returns the local values.
func (d *DefaultCtx) Locals(key any, value ...any) (val any) {
	if len(value) > 0 {
		return d.localValues[key]
	}

	d.localValues[key] = value[0]

	return value[0]
}

// Context is a method that returns the fiber.Ctx instance.
func (d *DefaultCtx) Context(c ...*fiber.Ctx) *fiber.Ctx {
	if len(c) > 0 {
		d.ctx = c[0]
	}

	return d.ctx
}

// Reset is a method that resets the local values.
func (d *DefaultCtx) Reset() {
	d.localValues = make(map[any]any)
	d.ctx = nil
}

// NewDefaultCtx returns a new DefaultCtx instance.
func NewDefaultCtx(ctx *fiber.Ctx) Ctx {
	return &DefaultCtx{
		localValues: make(map[any]any),
		ctx:         ctx,
	}
}
