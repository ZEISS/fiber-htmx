package toasts

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
)

// ToastProps contains the properties for the toast component.
type ToastProps struct {
	// ClassNames contains the class names for the toast component.
	ClassNames htmx.ClassNames
	// DisableCloseButton is enabled when the close button is disabled.
	DisableCloseButton bool
}

// Toast is a component for the htmx toast extension.
func Toast(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":     true,
				"z-[99999]": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastTopToastStart is a component for the htmx toast extension.
func ToastTopStart(props ToastProps, children ...htmx.Node) htmx.Node {
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
func ToastTopEnd(props ToastProps, children ...htmx.Node) htmx.Node {
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
func ToastEnd(p ToastProps, children ...htmx.Node) htmx.Node {
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

// ToastAlertInfo is a component for the htmx toast extension.
func ToastAlertInfo(props ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"flex":            true,
			"alert":           true,
			"alert-info":      true,
			"justify-between": true,
		},
		htmx.Role("alert"),
		htmx.Group(children...),
		htmx.HyperScript(`on load wait 5s then remove me`),
		htmx.If(
			!props.DisableCloseButton,
			buttons.Outline(
				buttons.ButtonProps{
					ClassNames: htmx.ClassNames{
						"btn-sm": true,
					},
				},
				htmx.HyperScript("on click remove closest .alert"),
				htmx.Text("Close"),
			),
		),
	)
}

// ToastAlertSuccess is a component for the htmx toast extension.
func ToastAlertSuccess(props ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"flex":            true,
			"alert":           true,
			"alert-success":   true,
			"justify-between": true,
		},
		htmx.Role("alert"),
		htmx.HyperScript(`on load wait 5s then remove me`),
		htmx.Group(children...),
		htmx.If(
			!props.DisableCloseButton,
			buttons.Outline(
				buttons.ButtonProps{
					ClassNames: htmx.ClassNames{
						"btn-sm": true,
					},
				},
				htmx.HyperScript("on click remove closest .alert"),
				htmx.Text("Close"),
			),
		),
	)
}

// ToastAlertError is a component for the htmx toast extension.
func ToastAlertError(props ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"flex":            true,
			"alert":           true,
			"alert-error":     true,
			"justify-between": true,
		},
		htmx.Role("alert"),
		htmx.HyperScript(`on load wait 5s then remove me`),
		htmx.Group(children...),
		htmx.If(
			!props.DisableCloseButton,
			buttons.Outline(
				buttons.ButtonProps{
					ClassNames: htmx.ClassNames{
						"btn-sm": true,
					},
				},
				htmx.HyperScript("on click remove closest .alert"),
				htmx.Text("Close"),
			),
		),
	)
}
