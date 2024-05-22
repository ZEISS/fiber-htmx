package htmx

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	// StatusStopPolling is a helper status code to stop polling.
	StatusStopPolling = 286
)

// HxRequestHeader is a helper type for htmx request headers.
type HxRequestHeader string

// String returns the header as a string.
func (h HxRequestHeader) String() string {
	return string(h)
}

// HxResponseHeader ...
type HxResponseHeader string

// HxResponseHeaders ...
type HxResponseHeaders struct {
	headers http.Header
}

// String returns the header as a string.
func (h HxResponseHeader) String() string {
	return string(h)
}

// Set is a helper function to set a header.
func (h *HxResponseHeaders) Set(k HxResponseHeader, val string) {
	h.headers.Set(k.String(), val)
}

// Get is a helper function to get a header.
func (h *HxResponseHeaders) Get(k HxResponseHeader) string {
	return h.headers.Get(k.String())
}

const (
	HXLocation           HxResponseHeader = "HX-Location"             // Allows you to do a client-side redirect that does not do a full page reload
	HXPushUrl            HxResponseHeader = "HX-Push-Url"             // pushes a new url into the history stack
	HXRedirect           HxResponseHeader = "HX-Redirect"             // can be used to do a client-side redirect to a new location
	HXRefresh            HxResponseHeader = "HX-Refresh"              // if set to "true" the client side will do a full refresh of the page
	HXReplaceUrl         HxResponseHeader = "HX-Replace-Url"          // replaces the current URL in the location bar
	HXReswap             HxResponseHeader = "HX-Reswap"               // Allows you to specify how the response will be swapped. See hx-swap for possible values
	HXRetarget           HxResponseHeader = "HX-Retarget"             // A CSS selector that updates the target of the content update to a different element on the page
	HXReselect           HxResponseHeader = "HX-Reselect"             // A CSS selector that allows you to choose which part of the response is used to be swapped in. Overrides an existing hx-select on the triggering element
	HXTrigger            HxResponseHeader = "HX-Trigger"              // allows you to trigger client side events, see the documentation for more info
	HXTriggerAfterSettle HxResponseHeader = "HX-Trigger-After-Settle" // allows you to trigger client side events, see the documentation for more info
	HXTriggerAfterSwap   HxResponseHeader = "HX-Trigger-After-Swap"   // allows you to trigger client side events, see the documentation for more info
)

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

// Redirect is a helper function to redirect the client.
func Redirect(c *fiber.Ctx, url string) {
	c.Set(HXRedirect.String(), url)
}

// ReplaceURL is a helper function to replace the current URL.
func ReplaceURL(c *fiber.Ctx, url string) {
	c.Set(HXReplaceUrl.String(), url)
}

// ReSwap is a helper function to swap the response.
func ReSwap(c *fiber.Ctx, target string) {
	c.Set(HXReswap.String(), target)
}

// ReTarget is a helper function to retarget the response.
func ReTarget(c *fiber.Ctx, target string) {
	c.Set(HXRetarget.String(), target)
}

// ReSelect is a helper function to reselect the response.
func ReSelect(c *fiber.Ctx, target string) {
	c.Set(HXReselect.String(), target)
}

// Triggers is a helper function to trigger an event.
func HxTriggers(c *fiber.Ctx, target string) {
	c.Set(HXTrigger.String(), target)
}

// Boosted returns true if the request is boosted.
func Boosted(c *fiber.Ctx) bool {
	return AsBool(c.Get(HxRequestHeaderBoosted.String()))
}

// CurrentURL returns the current URL.
func CurrentURL(c *fiber.Ctx) string {
	return c.Get(HxRequestHeaderCurrentURL.String())
}

// HistoryRestoreRequest returns true if the request is a history restore request.
func HistoryRestoreRequest(c *fiber.Ctx) bool {
	return AsBool(c.Get(HxRequestHeaderHistoryRestoreRequest.String()))
}

// Prompt returns the prompt.
func Prompt(c *fiber.Ctx) string {
	return c.Get(HxRequestHeaderPrompt.String())
}

// Request returns true if the request is an htmx request.
func Request(c *fiber.Ctx) bool {
	return AsBool(c.Get(HxRequestHeaderRequest.String()))
}

// Targets is a helper function to get the target.
func Targets(c *fiber.Ctx) string {
	return c.Get(HxRequestHeaderTarget.String())
}

// TriggerName is a helper function to get the trigger name.
func TriggerName(c *fiber.Ctx) string {
	return c.Get(HxRequestHeaderTriggerName.String())
}

// Trigger is a helper function to trigger an event.
func Trigger(c *fiber.Ctx, target string) {
	c.Set(HxRequestHeaderTrigger.String(), target)
}

// RenderPartial returns true if the request is an htmx request.
func RenderPartial(c *fiber.Ctx) bool {
	return (Request(c) || Boosted(c)) && !HistoryRestoreRequest(c)
}

// RenderComp is a helper function to render a component.
func RenderComp(c *fiber.Ctx, n Node, opt ...RenderOpt) error {
	for _, o := range opt {
		o(c)
	}

	return n.Render(c)
}

// StopPolling is a helper function to stop polling.
func StopPolling(c *fiber.Ctx) error {
	return c.SendStatus(StatusStopPolling)
}

// FilterFunc is a function that filters the context.
type FilterFunc func(c *fiber.Ctx) error

// Config ...
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	Next func(c *fiber.Ctx) bool

	// Filters is a list of filters that filter the context.
	Filters []FilterFunc

	// Resolvers is a list of resolvers that resolve the context.
	Resolvers []ResolveFunc

	// ErrorHandler is executed when an error is returned from fiber.Handler.
	//
	// Optional. Default: DefaultErrorHandler
	ErrorHandler fiber.ErrorHandler
}

// ConfigDefault is the default config.
var ConfigDefault = Config{
	ErrorHandler: defaultErrorHandler,
	Filters:      []FilterFunc{},
	Resolvers:    []ResolveFunc{},
}

// default ErrorHandler that process return error from fiber.Handler
func defaultErrorHandler(_ *fiber.Ctx, _ error) error {
	return fiber.ErrBadRequest
}

// RenderOpt is helper function to configure the render.
type RenderOpt func(c *fiber.Ctx)

// RenderStatusCode is a helper function to set the status code.
func RenderStatusCode(err error) RenderOpt {
	return func(c *fiber.Ctx) {
		var e *fiber.Error
		ok := errors.As(err, &e)
		if !ok {
			e = fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint("%w", err))
		}

		c.Status(e.Code)
	}
}

// NewCompHandler returns a new comp handler.
func NewCompHandler(n Node, config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		err := n.Render(c)
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		return nil
	}
}

// CompFunc is a helper type for component functions.
type CompFunc func(c *fiber.Ctx) (Node, error)

// NewCompFuncHandler returns a new comp handler.
func NewCompFuncHandler(handler CompFunc, config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		n, err := handler(c)
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		err = n.Render(c)
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		return nil
	}
}

// NewHxControllerHandler returns a new htmx controller handler.
// Deprecated: use NewControllerHandler instead.
func NewHxControllerHandler(ctrl Controller, config ...Config) fiber.Handler {
	return NewControllerHandler(ctrl, config...)
}

// NewControllerHandler returns a new htmx controller handler.
// nolint:gocyclo
func NewControllerHandler(ctrl Controller, config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) (err error) {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		defer func() {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)
				if !ok {
					err = fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("%v", r))
				}
				err = ctrl.Error(err)
			}
		}()

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		err = Resolve(c, cfg.Resolvers...)
		if err != nil {
			return ctrl.Error(err)
		}

		for _, f := range cfg.Filters {
			err = f(c)
			if err != nil {
				return ctrl.Error(err)
			}
		}

		err = ctrl.Init(c)
		if err != nil {
			return ctrl.Error(err)
		}

		err = ctrl.Prepare()
		if err != nil {
			return ctrl.Error(err)
		}

		switch c.Method() {
		case fiber.MethodGet:
			err = ctrl.Get()
		case fiber.MethodPost:
			err = ctrl.Post()
		case fiber.MethodPut:
			err = ctrl.Put()
		case fiber.MethodPatch:
			err = ctrl.Patch()
		case fiber.MethodDelete:
			err = ctrl.Delete()
		case fiber.MethodOptions:
			err = ctrl.Options()
		case fiber.MethodTrace:
			err = ctrl.Trace()
		case fiber.MethodHead:
			err = ctrl.Head()
		default:
			err = fiber.ErrMethodNotAllowed
		}

		if err != nil {
			return ctrl.Error(err)
		}

		err = ctrl.Finalize()
		if err != nil {
			return ctrl.Error(err)
		}

		return nil
	}
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.ErrorHandler == nil {
		cfg.ErrorHandler = ConfigDefault.ErrorHandler
	}

	if cfg.Filters == nil {
		cfg.Filters = ConfigDefault.Filters
	}

	if cfg.Resolvers == nil {
		cfg.Resolvers = ConfigDefault.Resolvers
	}

	return cfg
}
