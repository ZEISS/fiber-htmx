package htmx

import (
	"encoding/json"
	"html/template"
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
	_ = c.Locals(htmxContext, HxFromContext(c))

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

// Htmx is a helper struct for htmx requests.
type Htmx struct {
	request *Hx
	ctx     *fiber.Ctx
}

// IsHxRequest returns true if the request is an htmx request.
func (h *Htmx) IsHxRequest() bool {
	return h.request.HxRequest
}

// IsHxBoosted returns true if the request is an htmx request.
func (h *Htmx) IsHxBoosted() bool {
	return h.request.HxBoosted
}

// IsHxHistoryRestoreRequest returns true if the request is an htmx request.
func (h *Htmx) IsHxHistoryRestoreRequest() bool {
	return h.request.HxHistoryRestoreRequest
}

// RenderPartial reutrns true if the request is an htmx request.
func (h *Htmx) RenderPartial() bool {
	return (h.request.HxRequest || h.request.HxBoosted) && !h.request.HxHistoryRestoreRequest
}

// Write writes a response.
func (h *Htmx) Write(data []byte) (n int, err error) {
	return h.ctx.Write(data)
}

// WriteHTML writes an HTML response.
func (h *Htmx) WriteHTML(html template.HTML) (n int, err error) {
	return h.WriteString(string(html))
}

// WriteJSON writes a JSON response.
func (h *Htmx) WriteJSON(data any) (n int, err error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	return h.Write(payload)
}

// RenderComp is a helper function to render a component.
func (h *Htmx) RenderComp(n Node) error {
	return n.Render(h)
}

// WriteString is a helper function to write a string.
func (h *Htmx) WriteString(s string) (n int, err error) {
	return h.ctx.WriteString(s)
}

// StopPolling ...
func (h *Htmx) StopPolling() error {
	return h.ctx.SendStatus(StatusStopPolling)
}

// Ctx returns the fiber context.
func (h *Htmx) Ctx() *fiber.Ctx {
	return h.ctx
}

// Redirect is a helper function to redirect the client.
func (h *Htmx) Redirect(url string) {
	h.ctx.Set(HXRedirect.String(), url)
}

// ReplaceURL is a helper function to replace the current URL.
func (h *Htmx) ReplaceURL(url string) {
	h.ctx.Set(HXReplaceUrl.String(), url)
}

// ReSwap ...
func (h *Htmx) ReSwap(target string) {
	h.ctx.Set(HXReswap.String(), target)
}

// ReTarget ...
func (h *Htmx) ReTarget(target string) {
	h.ctx.Set(HXRetarget.String(), target)
}

// ReSelect ...
func (h *Htmx) ReSelect(target string) {
	h.ctx.Set(HXReselect.String(), target)
}

// Trigger ...
func (h *Htmx) Trigger(target string) {
	h.ctx.Set(HXTrigger.String(), target)
}

// HtmxHandler ...
type HtmxHandlerFunc = func(hx *Htmx) error

// NewHtmxHandler returns a new htmx handler.
func NewHtmxHandler(handler HtmxHandlerFunc, config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		c = ContextWithHx(c)
		hx := HxFromContext(c)

		h := &Htmx{hx, c}

		err := handler(h)
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		return nil
	}
}

// NewCompFuncHandler returns a new comp handler.

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

	return cfg
}
