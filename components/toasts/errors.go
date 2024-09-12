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

	if te.Level != SUCCESS {
		htmx.ReSwap(c, "none")
	}

	te.SetHXTriggerHeader(c)

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	return c.Status(code).SendString(err.Error())
}
