package htmx

import (
	"context"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type token struct{}

// ContextCtx is a struct that contains the context of a htmx context.
type Context interface {
	// Context is the fiber.Ctx instance.
	Context() *fiber.Ctx
	// Locals is a method that returns the local values.
	Locals(key any, value ...any) (val any)
	// Reset is a method that resets the local values.
	Reset()
}

var _ Context = (*Ctx)(nil)

// Ctx is a struct that contains the properties of a htmx context.
type Ctx struct {
	localValues map[any]any
	ctx         *fiber.Ctx

	wg      sync.WaitGroup
	errOnce sync.Once
	err     error
	cancel  func(error)

	sem chan token

	sync.RWMutex
}

// Err is a method that returns the error of the context.
func (c *Ctx) Err() error {
	return c.err
}

// Error ...
type Error struct {
	Err error
}

// Error ...
func (s *Error) Error() string { return fmt.Sprintf("server: %s", s.Err) }

// Unwrap ...
func (s *Error) Unwrap() error { return s.Err }

// NewError returns a new error.
func NewError(err error) *Error {
	return &Error{Err: err}
}

// Wait is waiting for all reservers to finish.
func (c *Ctx) Wait() error {
	c.wg.Wait()

	if c.cancel != nil {
		c.cancel(c.err)
	}

	return c.err
}

// ResolveFunc is a function that resolves locals for the context.
type ResolveFunc func(c context.Context) (interface{}, interface{}, error)

// Resolve is a method that resolves locals for the context.
func (c *Ctx) Resolve(f ResolveFunc) {
	if c.sem != nil {
		c.sem <- token{}
	}

	c.wg.Add(1)
	go func() {
		defer c.done()

		k, v, err := f(c.ctx.Context())
		if err != nil {
			c.errOnce.Do(func() {
				c.err = err
				if c.cancel != nil {
					c.cancel(c.err)
				}
			})
		}

		c.Locals(k, v)
	}()
}

func (c *Ctx) done() {
	if c.sem != nil {
		<-c.sem
	}

	c.wg.Done()
}

// Locals is a method that returns the local values.
func (c *Ctx) Locals(key any, value ...any) (val any) {
	c.Lock()
	defer c.Unlock()

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
	c.Lock()
	defer c.Unlock()

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
