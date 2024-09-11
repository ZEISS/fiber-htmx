package toasts

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
)

const defaultToasterID = "toaster"

// ToastDirection ...
type ToastDirection int

const (
	ToastDirectionEnd ToastDirection = iota // toast end default
	ToastDirectionTopStart
	ToastDirectionTopEnd
)

const (
	INFO    = "info"
	SUCCESS = "success"
	ERROR   = "error"
)

// Toast is a message to display to the user.
type Toast struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

// New returns a new Toast.
func New(level string, message string) Toast {
	return Toast{level, message}
}

// Error returns the error message.
func (t Toast) Error() string {
	return t.Message
}

// Info returns an info message.
func Info(message string) Toast {
	return New(INFO, message)
}

// Success returns a success message.
func Success(c *fiber.Ctx, message string) {
	New(SUCCESS, message).SetHXTriggerHeader(c)
}

// Error returns an error message.
func Error(message string) Toast {
	return New(ERROR, message)
}

// ToJson returns the JSON representation of the toast.
func (t Toast) ToJson() (string, error) {
	t.Message = t.Error()

	eventMap := map[string]Toast{}
	eventMap["notify"] = t

	jsonData, err := json.Marshal(eventMap)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// SetHXTriggerHeader sets the HTMX trigger header.
func (t Toast) SetHXTriggerHeader(c *fiber.Ctx) error {
	jsonData, err := t.ToJson()
	if err != nil {
		return err
	}

	htmx.Trigger(c, jsonData)

	return nil
}

// ToasterProps is the properties for the Toaster component.
type ToasterProps struct {
	// ClassNames are the class names for the toast.
	ClassNames htmx.ClassNames
}

// Toaster is the layout host for the toasts.
func Toaster() htmx.Node {
	return Toasts(
		ToastsProps{},
		alpine.XData(`{
    notifications: [],
    add(e) {
      this.notifications.push({
        id: e.timeStamp,
        level: e.detail.level,
        message: e.detail.message,
      })
    },
    remove(notification) {
      this.notifications = this.notifications.filter(i => i.id !== notification.id)
    },
  }`),
		htmx.Role("status"),
		htmx.Attribute("aria-live", "polite"),
		alpine.XOn("notify.window", "add($event)"),
		htmx.Template(
			alpine.XFor("notification in notifications"),
			htmx.Attribute(":key", "notification.id"),
			htmx.Div(
				htmx.ClassNames{
					"toast":     true,
					"z-[99999]": true,
				},
				alpine.XData(`{
        show: false,
        init() {
          this.$nextTick(() => this.show = true)

          setTimeout(() => this.transitionOut(), 3000)
        },
        transitionOut() {
          this.show = false

          setTimeout(() => this.remove(this.notification), 500)
        },
      }`),
				alpine.XShow("show"),
				alpine.XTransition("duration.500ms"),
				htmx.Div(
					htmx.ClassNames{
						"alert":       true,
						"alert-error": true,
					},
					alpine.XShow("notification.level === 'error'"),
					htmx.Div(
						alpine.XText("notification.message"),
					),
					buttons.Outline(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								"btn-sm": true,
							},
						},
						alpine.XOn("click", "transitionOut()"),
						htmx.Text("Close"),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"alert":         true,
						"alert-success": true,
					},
					alpine.XShow("notification.level === 'success'"),
					htmx.Div(
						alpine.XText("notification.message"),
					),
					buttons.Outline(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								"btn-sm": true,
							},
						},
						alpine.XOn("click", "transitionOut()"),
						htmx.Text("Close"),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"alert":      true,
						"alert-info": true,
					},
					alpine.XShow("notification.level === 'info'"),
					htmx.Div(
						alpine.XText("notification.message"),
					),
					buttons.Outline(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								"btn-sm": true,
							},
						},
						alpine.XOn("click", "transitionOut()"),
						htmx.Text("Close"),
					),
				),
			),
		),
	)
}

// ToastsProps ...
type ToastsProps struct {
	// ClassNames are the class names for the toast.
	ClassNames htmx.ClassNames
	// ID is the ID of the toast.
	ID string
	// Direction is the direction of the toast.
	Direction ToastDirection
}

// RenderToas is the handler for rendering the toasts.
func RenderToasts(c *fiber.Ctx, err error) error {
	te, ok := err.(Toast)

	if !ok {
		te = Error("there has been an unexpected error")
	}

	if te.Level != SUCCESS {
		htmx.ReSwap(c, "none")
	}

	return te.SetHXTriggerHeader(c)
}

// Notify is the container for the toast.
func Notify(c *fiber.Ctx, children ...htmx.Node) htmx.Node {
	htmx.ReSwap(c, "none")

	return htmx.Fragment(children...)
}

// Toasts are messsage to toast.
func Toasts(props ToastsProps, children ...htmx.Node) htmx.Node {
	if props.ID == "" {
		props.ID = defaultToasterID
	}

	return htmx.Div(
		htmx.ID(props.ID),
		htmx.If(
			props.Direction == ToastDirectionEnd, // toast end
			ToastEnd(
				ToastsProps{
					ClassNames: props.ClassNames,
				},
				children...,
			),
		),
		htmx.If(
			props.Direction == ToastDirectionTopStart,
			ToastTopStart(
				ToastsProps{
					ClassNames: props.ClassNames,
				},
				children...,
			),
		),
		htmx.If(
			props.Direction == ToastDirectionTopEnd,
			ToastTopEnd(
				ToastsProps{
					ClassNames: props.ClassNames,
				},
				children...,
			),
		),
	)
}

// ToastTopToastStart is a component for the htmx toast extension.
func ToastTopStart(props ToastsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":       true,
				"toast-top":   true,
				"toast-start": true,
				"z-[99999]":   true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastTopToastEnd is a component for the htmx toast extension.
func ToastTopEnd(props ToastsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":     true,
				"toast-top": true,
				"toast-end": true,
				"z-[99999]": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastEnd is a component for the htmx toast extension.
func ToastEnd(p ToastsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":     true,
				"toast-end": true,
				"z-[99999]": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
