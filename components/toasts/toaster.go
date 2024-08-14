package toasts

import (
	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

const defaultToasterID = "toaster"

// ToastDirection ...
type ToastDirection int

const (
	ToastDirectionEnd ToastDirection = iota // toast end default
	ToastDirectionTopStart
	ToastDirectionTopEnd
)

// ToastsProps ...
type ToastsProps struct {
	// ClassNames are the class names for the toast.
	ClassNames htmx.ClassNames
	// ID is the ID of the toast.
	ID string
	// Direction is the direction of the toast.
	Direction ToastDirection
}

// Toaster is the layout host for the toasts.
func Toaster(props ToastsProps, children ...htmx.Node) htmx.Node {
	if props.ID == "" {
		props.ID = defaultToasterID
	}

	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{},
			props.ClassNames,
		),
		htmx.ID(props.ID),
		htmx.Group(children...),
	)
}

// RenderToasts is the handler for rendering the toasts.
func RenderToasts(c *fiber.Ctx, children ...htmx.Node) error {
	htmx.ReSwap(c, "none") // this forces htmx not to replace the content

	return htmx.RenderComp(c, htmx.Fragment(children...))
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
		htmx.HxSwapOob("true"), // this is the key to making the toast work
		htmx.If(
			props.Direction == ToastDirectionEnd, // toast end
			ToastEnd(
				ToastProps{
					ClassNames: props.ClassNames,
				},
				children...,
			),
		),
		htmx.If(
			props.Direction == ToastDirectionTopStart,
			ToastTopStart(
				ToastProps{
					ClassNames: props.ClassNames,
				},
				children...,
			),
		),
		htmx.If(
			props.Direction == ToastDirectionTopEnd,
			ToastTopEnd(
				ToastProps{
					ClassNames: props.ClassNames,
				},
				children...,
			),
		),
	)
}
