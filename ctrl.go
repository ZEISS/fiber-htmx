package htmx

import "github.com/gofiber/fiber/v2"

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
}

var _ Controller = (*UnimplementedController)(nil)

// UnimplementedController is the default controller implementation.
type UnimplementedController struct {
	hx *Htmx
}

// Hx returns the htmx instance.
func (c *UnimplementedController) Hx() *Htmx {
	return c.hx
}

// Init is called when the controller is initialized.
func (c *UnimplementedController) Init(hx *Htmx) error {
	c.hx = hx

	return nil
}

// Prepare is called before the controller is executed.
func (c *UnimplementedController) Prepare() error {
	return nil
}

// Finalize is called after the controller is executed.
func (c *UnimplementedController) Finalize() error {
	return nil
}

// Get is called when the controller is executed with the GET method.
func (c *UnimplementedController) Get() error {
	return c.Hx().Ctx().SendStatus(fiber.StatusNotImplemented)
}

// Post is called when the controller is executed with the POST method.
func (c *UnimplementedController) Post() error {
	return c.Hx().Ctx().SendStatus(fiber.StatusNotImplemented)
}

// Put is called when the controller is executed with the PUT method.
func (c *UnimplementedController) Put() error {
	return c.Hx().Ctx().SendStatus(fiber.StatusNotImplemented)
}

// Patch is called when the controller is executed with the PATCH method.
func (c *UnimplementedController) Patch() error {
	return c.Hx().Ctx().SendStatus(fiber.StatusNotImplemented)
}

// Delete is called when the controller is executed with the DELETE method.
func (c *UnimplementedController) Delete() error {
	return c.Hx().Ctx().SendStatus(fiber.StatusNotImplemented)
}

// Options is called when the controller is executed with the OPTIONS method.
func (c *UnimplementedController) Options() error {
	return c.Hx().Ctx().SendStatus(fiber.StatusNotImplemented)
}
