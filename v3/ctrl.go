package htmx

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// ErrUnimplemented is returned when a method is not implemented.
var ErrUnimplemented = fiber.NewError(fiber.StatusNotImplemented, "method not implemented")

// Controller is the interface for the htmx controller.
type Controller interface {
	// Prepare is called before the controller is executed.
	Prepare(c fiber.Ctx) error
	// Finalize is called after the controller is executed.
	Finalize(c fiber.Ctx) error
	// Error is called when the controller returns an error.
	Error(c fiber.Ctx, err ...error) error
	// Get is called when the controller is executed with the GET method.
	Get(c fiber.Ctx) error
	// Post is called when the controller is executed with the POST method.
	Post(c fiber.Ctx) error
	// Put is called when the controller is executed with the PUT method.
	Put(c fiber.Ctx) error
	// Patch is called when the controller is executed with the PATCH method.
	Patch(c fiber.Ctx) error
	// Delete is called when the controller is executed with the DELETE method.
	Delete(c fiber.Ctx) error
	// Options is called when the controller is executed with the OPTIONS method.
	Options(c fiber.Ctx) error
	// Trace is called when the controller is executed with the TRACE method.
	Trace(c fiber.Ctx) error
	// Head is called when the controller is executed with the HEAD method.
	Head(c fiber.Ctx) error
	// Reset resets the controller.
	Reset()
}

var _ Controller = (*UnimplementedController)(nil)

// UnimplementedController is the default controller implementation.
type UnimplementedController struct{}

// Prepare is called before the controller is executed.
func (ctrl *UnimplementedController) Prepare(c fiber.Ctx) error {
	return nil
}

// Finalize is called after the controller is executed.
func (ctrl *UnimplementedController) Finalize(c fiber.Ctx) error {
	return nil
}

// Get is called when the controller is executed with the GET method.
func (ctrl *UnimplementedController) Get(c fiber.Ctx) error {
	return ErrUnimplemented
}

// Post is called when the controller is executed with the POST method.
func (ctrl *UnimplementedController) Post(c fiber.Ctx) error {
	return ErrUnimplemented
}

// Put is called when the controller is executed with the PUT method.
func (ctrl *UnimplementedController) Put(c fiber.Ctx) error {
	return ErrUnimplemented
}

// Patch is called when the controller is executed with the PATCH method.
func (ctrl *UnimplementedController) Patch(c fiber.Ctx) error {
	return ErrUnimplemented
}

// Delete is called when the controller is executed with the DELETE method.
func (ctrl *UnimplementedController) Delete(c fiber.Ctx) error {
	return ErrUnimplemented
}

// Options is called when the controller is executed with the OPTIONS method.
func (ctrl *UnimplementedController) Options(c fiber.Ctx) error {
	return ErrUnimplemented
}

// Head is called when the controller is executed with the HEAD method.
func (ctrl *UnimplementedController) Head(c fiber.Ctx) error {
	return ErrUnimplemented
}

// Trace is called when the controller is executed with the TRACE method.
func (ctrl *UnimplementedController) Trace(c fiber.Ctx) error {
	return ErrUnimplemented
}

// Error is called when the controller returns an error.
func (ctrl *UnimplementedController) Error(c fiber.Ctx, err ...error) error {
	return ErrUnimplemented
}

// Reset resets the controller.
func (ctrl *UnimplementedController) Reset() {}

var _ TransactionController = (*UnimplementedTransactionController)(nil)

// TransactionController is the interface for a controller that also does database transactions.
type TransactionController interface {
	// Returns the transaction.
	Tx() *gorm.DB
	// Begin begins a transaction.
	Begin(*gorm.DB) error
	// Commit commits the transaction.
	Commmit() error
	// Rollback rolls back the transaction.
	Rollback() error
}

// UnimplementedTransactionController is the default implementation of a TransactionController.
type UnimplementedTransactionController struct {
	TransactionController
	tx *gorm.DB
}

// Begin begins a transaction.
func (c *UnimplementedTransactionController) Begin(conn *gorm.DB) error {
	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	c.tx = tx

	return nil
}

// Commit commits the transaction.
func (c *UnimplementedTransactionController) Commit() error {
	if c.tx == nil {
		return nil
	}

	c.tx.Commit()
	if err := c.tx.Error; err != nil {
		return err
	}

	c.tx = nil

	return nil
}

// Rollback rolls back the transaction.
func (c *UnimplementedTransactionController) Rollback() error {
	if c.tx == nil {
		return nil
	}

	c.tx.Rollback()
	if err := c.tx.Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}
	c.tx = nil

	return nil
}

// Txn returns the transaction.
func (c *UnimplementedTransactionController) Tx() *gorm.DB {
	if c.tx != nil {
		return c.tx
	}

	return nil
}
