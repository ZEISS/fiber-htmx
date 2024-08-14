package loading

import htmx "github.com/zeiss/fiber-htmx"

// SpinnerProps is the type of the props for the Spinner component
type SpinnerProps struct {
	ClassNames htmx.ClassNames
}

// Spinner is a component that displays a spinner.
func Spinner(p SpinnerProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"loading":         true,
				"loading-spinner": true,
				"loading-sm":      true,
			},
			p.ClassNames,
		),
	)
}

// SpinnerSmall is a component that displays a small spinner.
func SpinnerSmall(p SpinnerProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"loading":         true,
				"loading-spinner": true,
				"loading-sm":      true,
			},
			p.ClassNames,
		),
	)
}

// SpinnerXtraSmall is a component that displays an extra small spinner.
func SpinnerXtraSmall(p SpinnerProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"loading":         true,
				"loading-spinner": true,
				"loading-xs":      true,
			},
			p.ClassNames,
		),
	)
}

// SpinnerLarge is a component that displays a large spinner.
func SpinnerLarge(p SpinnerProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"loading":         true,
				"loading-spinner": true,
				"loading-lg":      true,
			},
			p.ClassNames,
		),
	)
}

// SpinnerMedium is a component that displays a medium spinner.
func SpinnerMedium(p SpinnerProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"loading":         true,
				"loading-spinner": true,
				"loading-md":      true,
			},
			p.ClassNames,
		),
	)
}
