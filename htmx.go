package htmx

import (
	"github.com/gofiber/fiber/v2"
)

// HxRequestHeader ...
type HxRequestHeader string

// String ...
func (h HxRequestHeader) String() string {
	return string(h)
}

const (
	HxRequestHeaderBoosted               HxRequestHeader = "HX-Boosted"
	HxRequestHeaderCurrentURL            HxRequestHeader = "HX-Current-URL"
	HxRequestHeaderHistoryRestoreRequest HxRequestHeader = "HX-History-Restore-Request"
	HxRequestHeaderPrompt                HxRequestHeader = "HX-Prompt"
	HxRequestHeaderRequest               HxRequestHeader = "HX-Request"
	HxRequestHeaderTarget                HxRequestHeader = "HX-Target"
	HxRequestHeaderTrigger               HxRequestHeader = "HX-Trigger"
	HxRequestHeaderTriggerName           HxRequestHeader = "HX-Trigger-Name"
)

// Hx ...
type Hx struct {
	HxBoosted               bool
	HxCurrentURL            string
	HxHistoryRestoreRequest bool
	HxPrompt                string
	HxRequest               bool
	HxTarget                string
	HxTriggerName           string
	HxTrigger               string
}

// HxFromContext ...
func HxFromContext(c *fiber.Ctx) *Hx {
	return &Hx{
		HxBoosted:               AsBool(c.Get(HxRequestHeaderBoosted.String())),
		HxCurrentURL:            c.Get(HxRequestHeaderCurrentURL.String()),
		HxHistoryRestoreRequest: AsBool(c.Get(HxRequestHeaderHistoryRestoreRequest.String())),
		HxPrompt:                c.Get(HxRequestHeaderPrompt.String()),
		HxRequest:               AsBool(c.Get(HxRequestHeaderRequest.String())),
		HxTarget:                c.Get(HxRequestHeaderTarget.String()),
		HxTriggerName:           c.Get(HxRequestHeaderTriggerName.String()),
		HxTrigger:               c.Get(HxRequestHeaderTrigger.String()),
	}
}

// ContextWithHx ...
func ContextWithHx(c *fiber.Ctx) *fiber.Ctx {
	c.Locals(htmxContext, HxFromContext(c))

	return c
}

// The contextKey type is unexported to prevent collisions with context keys defined in
// other packages.
type contextKey int

// The keys for the values in context
const (
	htmxContext contextKey = iota
)

// Config ...
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	Next func(c *fiber.Ctx) bool

	// ErrorHandler is executed when an error is returned from fiber.Handler.
	//
	// Optional. Default: DefaultErrorHandler
	ErrorHandler fiber.ErrorHandler
}

// ConfigDefault is the default config.
var ConfigDefault = Config{
	ErrorHandler: defaultErrorHandler,
}

// default ErrorHandler that process return error from fiber.Handler
func defaultErrorHandler(_ *fiber.Ctx, _ error) error {
	return fiber.ErrBadRequest
}

// New ...
func New(config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		return ContextWithHx(c).Next()
	}
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	return cfg
}
