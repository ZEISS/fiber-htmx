package toasts

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

// DefaultErrorHandler is the default error handler for toasts.
var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var te Toast
	if !errors.As(err, &te) {
		te = Error("there has been an unexpected error")
	}

	var e *fiber.Error // if this is not a toast then use the error message
	if errors.As(err, &e) {
		code = e.Code
		te = Error(e.Message)
	}

	if te.Level != SUCCESS {
		htmx.ReSwap(c, "none")
	}

	te.SetHXTriggerHeader(c)

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	return c.Status(code).SendString(err.Error())
}
