package htmx

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/katallaxie/pkg/utils"
)

// Locals is a method that returns the local values.
func Locals[V any](c *fiber.Ctx, key any, value ...V) V {
	var v V
	var ok bool

	if len(value) > 0 {
		v, ok = c.Locals(key, value).(V)
	} else {
		v, ok = c.Locals(key).(V)
	}

	if !ok {
		return utils.Zero[V]()
	}

	return v
}

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

// Hx is an interface for htmx requests.
type Hx interface {
	// HxBoosted returns true if the request is boosted.
	HxBoosted() bool
	// HxCurrentURL returns the current URL.
	HxCurrentURL() string
	// HxHistoryRestoreRequest returns true if the request is a history restore request.
	HxHistoryRestoreRequest() bool
	// HxPrompt returns the prompt.
	HxPrompt() string
	// HxRequest returns true if the request is an htmx request.
	HxRequest() bool
	// HxTarget returns the target.
	HxTarget() string
	// HxTriggerName returns the trigger name.
	HxTriggerName() string
	// HxTrigger returns the trigger.
	HxTrigger() string
	// HxRenderPartial returns true if the request is an htmx request.
	RenderPartial() bool
	// Write writes a response.
	Write(data []byte) (n int, err error)
	// WriteHTML writes an HTML response.
	WriteHTML(html template.HTML) (n int, err error)
	// WriteJSON writes a JSON response.
	WriteJSON(data any) (n int, err error)
	// WriteString writes a string.
	WriteString(s string) (n int, err error)
	// RenderComp renders a component.
	RenderComp(n Node, opt ...RenderOpt) error
	// StopPolling stops polling.
	StopPolling() error
	// Redirect redirects the client.
	Redirect(url string)
	// ReplaceURL replaces the current URL.
	ReplaceURL(url string)
	// ReSwap ...
	ReSwap(target string)
	// ReTarget ...
	ReTarget(target string)
	// ReSelect ...
	ReSelect(target string)
	// Trigger ...
	Trigger(target string)
}

// Redirect is a helper function to redirect the client.
func (h *hx) Redirect(url string) {
	h.ctx.Set(HXRedirect.String(), url)
}

// ReplaceURL is a helper function to replace the current URL.
func (h *hx) ReplaceURL(url string) {
	h.ctx.Set(HXReplaceUrl.String(), url)
}

// ReSwap ...
func (h *hx) ReSwap(target string) {
	h.ctx.Set(HXReswap.String(), target)
}

// ReTarget ...
func (h *hx) ReTarget(target string) {
	h.ctx.Set(HXRetarget.String(), target)
}

// ReSelect ...
func (h *hx) ReSelect(target string) {
	h.ctx.Set(HXReselect.String(), target)
}

// Trigger ...
func (h *hx) Trigger(target string) {
	h.ctx.Set(HXTrigger.String(), target)
}

type hx struct {
	ctx *fiber.Ctx
}

// HxBoosted returns true if the request is boosted.
func (h *hx) HxBoosted() bool {
	return AsBool(h.ctx.Get(HxRequestHeaderBoosted.String()))
}

// HxCurrentURL returns the current URL.
func (h *hx) HxCurrentURL() string {
	return h.ctx.Get(HxRequestHeaderCurrentURL.String())
}

// HxHistoryRestoreRequest returns true if the request is a history restore request.
func (h *hx) HxHistoryRestoreRequest() bool {
	return AsBool(h.ctx.Get(HxRequestHeaderHistoryRestoreRequest.String()))
}

// HxPrompt returns the prompt.
func (h *hx) HxPrompt() string {
	return h.ctx.Get(HxRequestHeaderPrompt.String())
}

// HxRequest returns true if the request is an htmx request.
func (h *hx) HxRequest() bool {
	return AsBool(h.ctx.Get(HxRequestHeaderRequest.String()))
}

// HxTarget returns the target.
func (h *hx) HxTarget() string {
	return h.ctx.Get(HxRequestHeaderTarget.String())
}

// HxTriggerName returns the trigger name.
func (h *hx) HxTriggerName() string {
	return h.ctx.Get(HxRequestHeaderTriggerName.String())
}

// HxTrigger returns the trigger.
func (h *hx) HxTrigger() string {
	return h.ctx.Get(HxRequestHeaderTrigger.String())
}

// RenderPartial returns true if the request is an htmx request.
func (h *hx) RenderPartial() bool {
	return (h.HxRequest() || h.HxBoosted()) && !h.HxHistoryRestoreRequest()
}

// Write writes a response.
func (h *hx) Write(data []byte) (n int, err error) {
	return h.ctx.Write(data)
}

// WriteString writes a string.
func (h *hx) WriteString(s string) (n int, err error) {
	return h.ctx.WriteString(s)
}

// WriteHTML writes an HTML response.
func (h *hx) WriteHTML(html template.HTML) (n int, err error) {
	return h.WriteString(string(html))
}

// WriteJSON writes a JSON response.
func (h *hx) WriteJSON(data any) (n int, err error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	h.ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return h.Write(payload)
}

// RenderComp is a helper function to render a component.
func (h *hx) RenderComp(n Node, opt ...RenderOpt) error {
	for _, o := range opt {
		o(h)
	}

	return n.Render(h)
}

// StopPolling ...
func (h *hx) StopPolling() error {
	return h.ctx.SendStatus(StatusStopPolling)
}

// NewHx returns a new htmx request.
func NewHx(c *fiber.Ctx) Hx {
	return &hx{c}
}

// ContextWithHx ...
func ContextWithHx(c *fiber.Ctx) *fiber.Ctx {
	_ = c.Locals(htmxContext, NewHx(c))

	return c
}

// HxFromContext is a helper function to get the htmx from the context.
func HxFromContext(c *fiber.Ctx) Hx {
	return Locals[Hx](c, htmxContext)
}

// The contextKey type is unexported to prevent collisions with context keys defined in
// other packages.
type contextKey int

// The keys for the values in context
const (
	htmxContext contextKey = iota
)

// FilterFunc is a function that filters the context.
type FilterFunc func(c *fiber.Ctx) error

// Config ...
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	Next func(c *fiber.Ctx) bool

	// Filters is a list of filters that filter the context.
	Filters []FilterFunc

	// ErrorHandler is executed when an error is returned from fiber.Handler.
	//
	// Optional. Default: DefaultErrorHandler
	ErrorHandler fiber.ErrorHandler
}

// ConfigDefault is the default config.
var ConfigDefault = Config{
	ErrorHandler: defaultErrorHandler,
	Filters:      []FilterFunc{},
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

// RenderOpt is helper function to configure the render.
type RenderOpt func(h *hx)

// RenderStatusCode is a helper function to set the status code.
func RenderStatusCode(err error) RenderOpt {
	return func(h *hx) {
		var e *fiber.Error
		ok := errors.As(err, &e)
		if !ok {
			e = fiber.NewError(fiber.StatusInternalServerError, fmt.Sprint("%w", err))
		}

		h.ctx.Status(e.Code)
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
// nolint:gocyclo
func NewHxControllerHandler(ctrl Controller, config ...Config) fiber.Handler {
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

		c = ContextWithHx(c)

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

	return cfg
}
