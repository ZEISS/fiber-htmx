package toasts

import htmx "github.com/zeiss/fiber-htmx"

// ToastProps contains the properties for the toast component.
type ToastProps struct {
	ClassNames htmx.ClassNames
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
func ToastTopStart(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":       true,
				"toast-top":   true,
				"toast-start": true,
				"z-[99999]":   true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastTopToastEnd is a component for the htmx toast extension.
func ToastTopEnd(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":     true,
				"toast-top": true,
				"toast-end": true,
				"z-[99999]": true,
			},
			p.ClassNames,
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
func ToastAlertInfo(children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"alert":      true,
			"alert-info": true,
		},
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}

// ToastAlertSuccess is a component for the htmx toast extension.
func ToastAlertSuccess(children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"alert":         true,
			"alert-success": true,
		},
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}

// ToastAlertError is a component for the htmx toast extension.
func ToastAlertError(children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"alert":       true,
			"alert-error": true,
		},
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}
