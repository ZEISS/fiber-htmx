package htmx

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Controller is the interface for the htmx controller.
type Controller interface {
	// Init is called when the controller is initialized.
	Init(hx *Htmx) error
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
	hx *Htmx

	sync.RWMutex
}

// Hx returns the htmx instance.
func (c *DefaultController) Hx() *Htmx {
	return c.hx
}

// Init is called when the controller is initialized.
func (c *DefaultController) Init(hx *Htmx) error {
	c.hx = hx

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
	return c.Hx().Ctx().BodyParser(obj)
}

// BindParams binds the params to the given struct.
func (c *DefaultController) BindParams(obj interface{}) error {
	return c.Hx().Ctx().ParamsParser(obj)
}

// BindQuery binds the query to the given struct.
func (c *DefaultController) BindQuery(obj interface{}) error {
	return c.Hx().Ctx().QueryParser(obj)
}
