package htmx

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var _ Ctx = (*Htmx)(nil)

// Ctx is the component context.
type Ctx interface {
	// Values is a helper function to get the values from the context.
	Values(key any, value ...any) (val any)
	// ValuesString is a helper function to get the values as a string from the context.
	ValuesString(key any, value ...any) (val string)
	// ValuesInt is a helper function to get the values as an int from the context.
	ValuesInt(key any, value ...any) (val int)
	// ValuesBool is a helper function to get the values as a bool from the context.
	ValuesBool(key any, value ...any) (val bool)
	// Path is a helper function to get the path from the context.
	Path() string
}

// UnimplementedContext is a helper type for unimplemented contexts.
type UnimplementedContext struct{}

// Values is a helper function to get the values from the context.
func (u *UnimplementedContext) Values(key any, value ...any) (val any) {
	return nil
}

// ValuesString is a helper function to get the values as a string from the context.
func (u *UnimplementedContext) ValuesString(key any, value ...any) (val string) {
	return ""
}

// ValuesInt is a helper function to get the values as an int from the context.
func (u *UnimplementedContext) ValuesInt(key any, value ...any) (val int) {
	return 0
}

// ValuesBool is a helper function to get the values as a bool from the context.
func (u *UnimplementedContext) ValuesBool(key any, value ...any) (val bool) {
	return false
}

// Path is a helper function to get the path from the context.
func (u *UnimplementedContext) Path() string {
	return ""
}

const (
	// StatusStopPolling is a helper status code to stop polling.
	StatusStopPolling = 286
)

// ResolveFunc is a function that resolves locals for the context.
type ResolveFunc func(c *fiber.Ctx) (interface{}, interface{}, error)

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

	// Resolvers is a list of resolvers that resolve locals for the context.
	Resolvers []ResolveFunc

	// ErrorHandler is executed when an error is returned from fiber.Handler.
	//
	// Optional. Default: DefaultErrorHandler
	ErrorHandler fiber.ErrorHandler
}

// ConfigDefault is the default config.
var ConfigDefault = Config{
	ErrorHandler: defaultErrorHandler,
	Resolvers:    []ResolveFunc{},
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
	localValues map[any]any

	request *Hx
	ctx     *fiber.Ctx

	sync.RWMutex
}

// Resolve is a method that resolves locals for the context.
func (h *Htmx) Resolve(ctx *fiber.Ctx, funcs ...ResolveFunc) error {
	var wg sync.WaitGroup
	var errOnce sync.Once
	var err error

	for _, f := range funcs {
		f := f

		wg.Add(1)
		go func() {
			defer wg.Done()

			k, v, errr := f(ctx)
			if errr != nil {
				errOnce.Do(func() {
					err = errr
				})
			}

			h.Locals(k, v)
		}()
	}

	wg.Wait()

	return err
}

// Locals is a method that returns the local values.
func (h *Htmx) Locals(key any, value ...any) (val any) {
	h.Lock()
	defer h.Unlock()

	if len(value) == 0 {
		return h.localValues[key]
	}

	h.localValues[key] = value[0]

	return value[0]
}

// Values is a method that returns the local values.
func (h *Htmx) Values(key any, value ...any) (val any) {
	return h.Locals(key, value...)
}

// ValuesString is a method that returns the local values.
func (h *Htmx) ValuesString(key any, value ...any) (val string) {
	return h.Locals(key, value...).(string)
}

// ValuesInt is a method that returns the local values.
func (h *Htmx) ValuesInt(key any, value ...any) (val int) {
	return h.Locals(key, value...).(int)
}

// ValuesBool is a method that returns the local values.
func (h *Htmx) ValuesBool(key any, value ...any) (val bool) {
	return h.Locals(key, value...).(bool)
}

// Reset is a method that resets the local values.
func (h *Htmx) Reset() {
	h.Lock()
	defer h.Unlock()

	h.localValues = make(map[any]any)
	h.ctx = nil
}

// Context is a method that returns the fiber context.
func (h *Htmx) Context() *fiber.Ctx {
	return h.ctx
}

// Path is a method that returns the path.
func (h *Htmx) Path() string {
	return h.ctx.Path()
}

// HtmxFromContext is a helper function to get the htmx from the context.
func HtmxFromContext(c *fiber.Ctx) *Htmx {
	return &Htmx{
		localValues: make(map[any]any),
		ctx:         c,
	}
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

// RenderPartial returns true if the request is an htmx request.
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

	h.ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

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

		h := &Htmx{
			localValues: make(map[any]any),
			request:     hx,
			ctx:         c,
		}

		err := h.Resolve(c, cfg.Resolvers...)
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		err = handler(h)
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		return nil
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

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		c = ContextWithHx(c)
		hx := HxFromContext(c)

		h := &Htmx{
			localValues: make(map[any]any),
			request:     hx,
			ctx:         c,
		}

		err := h.Resolve(c, cfg.Resolvers...)
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		err = ctrl.Init(h)
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		err = ctrl.Prepare()
		if err != nil {
			return cfg.ErrorHandler(c, err)
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
		default:
			err = fiber.ErrMethodNotAllowed
		}

		if err != nil {
			return cfg.ErrorHandler(c, err)
		}

		err = ctrl.Finalize()
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

	if cfg.Resolvers == nil {
		cfg.Resolvers = ConfigDefault.Resolvers
	}

	return cfg
}
