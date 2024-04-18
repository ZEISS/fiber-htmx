package htmx

import (
	"context"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Controller is the interface for the htmx controller.
type Controller interface {
	// Init is called when the controller is initialized.
	Init(ctx *fiber.Ctx) error
	// Prepare is called before the controller is executed.
	Prepare() error
	// Finalize is called after the controller is executed.
	Finalize() error
	// Get is called when the controller is executed with the GET method.
	Get() error
	// Post is called when the controller is executed with the POST method.
	Post() error
	// Put is called when the controller is executed with the PUT method.
	Put() error
	// Patch is called when the controller is executed with the PATCH method.
	Patch() error
	// Delete is called when the controller is executed with the DELETE method.
	Delete() error
	// Options is called when the controller is executed with the OPTIONS method.
	Options() error
	// Trace is called when the controller is executed with the TRACE method.
	Trace() error
	// Head  is called when the controller is executed with the HEAD method.
	Head() error
	// Error is called when an error occurs.
	Error(err error) error
}

var _ Controller = (*DefaultController)(nil)

// UnimplementedController is not to be used anymore.
// Deprecated: Use DefaultController instead.
type UnimplementedController = DefaultController

// NewDefaultController returns a new default controller.
func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

// UnimplementedController is the default controller implementation.
type DefaultController struct {
	ctx *fiber.Ctx
}

// BindFunc is a function that returns a context.
type BindFunc func(ctx *fiber.Ctx) (any, any, error)

// Init is called when the controller is initialized.
func (c *DefaultController) Init(ctx *fiber.Ctx) error {
	c.Reset()
	c.ctx = ctx

	return nil
}

// Prepare is called before the controller is executed.
func (c *DefaultController) Prepare() error {
	return nil
}

// Finalize is called after the controller is executed.
func (c *DefaultController) Finalize() error {
	return nil
}

// Get is called when the controller is executed with the GET method.
func (c *DefaultController) Get() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Post is called when the controller is executed with the POST method.
func (c *DefaultController) Post() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Put is called when the controller is executed with the PUT method.
func (c *DefaultController) Put() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Patch is called when the controller is executed with the PATCH method.
func (c *DefaultController) Patch() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Delete is called when the controller is executed with the DELETE method.
func (c *DefaultController) Delete() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Options is called when the controller is executed with the OPTIONS method.
func (c *DefaultController) Options() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Error is called when an error occurs.
func (c *DefaultController) Error(err error) error {
	return fiber.NewError(fiber.StatusInternalServerError)
}

// Head is called when the controller is executed with the HEAD method.
func (c *DefaultController) Head() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// Trace is called when the controller is executed with the TRACE method.
func (c *DefaultController) Trace() error {
	return fiber.NewError(fiber.StatusNotImplemented)
}

// BindForm binds the form to the given struct.
func (c *DefaultController) BindBody(obj interface{}) error {
	return c.ctx.BodyParser(obj)
}

// BindParams binds the params to the given struct.
func (c *DefaultController) BindParams(obj interface{}) error {
	return c.ctx.ParamsParser(obj)
}

// BindQuery binds the query to the given struct.
func (c *DefaultController) BindQuery(obj interface{}) error {
	return c.ctx.QueryParser(obj)
}

// Values is a helper function to get the values from the context.
func (c *DefaultController) Values(key any, value ...any) (val any) {
	return c.ctx.Locals(key, value...)
}

// ValuesString is a helper function to get the values as a string from the context.
func (c *DefaultController) ValuesString(key any, value ...any) (val string) {
	return c.ctx.Locals(key, value...).(string)
}

// ValuesInt is a helper function to get the values as an int from the context.
func (c *DefaultController) ValuesInt(key any, value ...any) (val int) {
	return c.ctx.Locals(key, value...).(int)
}

// ValuesBool is a helper function to get the values as a bool from the context.
func (c *DefaultController) ValuesBool(key any, value ...any) (val bool) {
	return c.ctx.Locals(key, value...).(bool)
}

// Path is a helper function to get the path from the context.
func (c *DefaultController) Path() string {
	return c.ctx.Path()
}

// Hx is a helper function to get the htmx from the context.
func (c *DefaultController) Hx() Hx {
	return HxFromContext(c.ctx)
}

// Context returns the context.
func (c *DefaultController) Context() context.Context {
	return c.ctx.Context()
}

// Ctx returns the fiber.Ctx.
func (c *DefaultController) Ctx() *fiber.Ctx {
	return c.ctx
}

// BindValues binds the values to the context.
func (c *DefaultController) BindValues(funcs ...BindFunc) error {
	var wg sync.WaitGroup
	var errOnce sync.Once
	var err error

	for _, f := range funcs {
		f := f

		wg.Add(1)
		go func() {
			defer wg.Done()

			k, v, errr := f(c.ctx)
			if errr != nil {
				errOnce.Do(func() {
					err = errr
				})
			}

			if errr == nil {
				c.ctx.Locals(k, v)
			}
		}()
	}

	wg.Wait()

	return err
}

// Reset resets the controller.
func (c *DefaultController) Reset() {
	c.ctx = nil
}
