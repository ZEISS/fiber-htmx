package alerts

import htmx "github.com/zeiss/fiber-htmx"

// AlertProps is the type of the props for the Alert component
type AlertProps struct {
	ClassNames htmx.ClassNames
}

// Alert is a component that displays an alert.
func Alert(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"alert": true,
		}.Merge(p.ClassNames),
		htmx.SVG(
			htmx.ClassNames{
				"stroke-info": true,
				"shrink-0":    true,
				"w-6":         true,
				"h-6":         true,
			},
			htmx.SVG(
				htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				htmx.Attribute("fill", "none"),
				htmx.Attribute("viewBox", "0 0 24 24"),
				htmx.Path(
					htmx.Attribute("stroke-linecap", "round"),
					htmx.Attribute("stroke-linejoin", "round"),
					htmx.Attribute("stroke-width", "2"),
					htmx.Attribute("d", "M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"),
				),
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}

// Info is a component that displays an info alert.
func Info(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"alert":      true,
			"alert-info": true,
		}.Merge(p.ClassNames),
		htmx.SVG(
			htmx.ClassNames{
				"stroke-info": true,
				"shrink-0":    true,
				"w-6":         true,
				"h-6":         true,
			},
			htmx.SVG(
				htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				htmx.Attribute("fill", "none"),
				htmx.Attribute("viewBox", "0 0 24 24"),
				htmx.Path(
					htmx.Attribute("stroke-linecap", "round"),
					htmx.Attribute("stroke-linejoin", "round"),
					htmx.Attribute("stroke-width", "2"),
					htmx.Attribute("d", "M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"),
				),
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}

// Success is a component that displays a success alert.
func Success(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"alert":         true,
			"alert-success": true,
		}.Merge(p.ClassNames),
		htmx.SVG(
			htmx.ClassNames{
				"stroke-success": true,
				"shrink-0":       true,
				"w-6":            true,
				"h-6":            true,
			},
			htmx.SVG(
				htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				htmx.Attribute("fill", "none"),
				htmx.Attribute("viewBox", "0 0 24 24"),
				htmx.Path(
					htmx.Attribute("stroke-linecap", "round"),
					htmx.Attribute("stroke-linejoin", "round"),
					htmx.Attribute("stroke-width", "2"),
					htmx.Attribute("d", "M5 13l4 4L19 7"),
				),
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}

// Warning is a component that displays a warning alert.
func Warning(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"alert":         true,
			"alert-warning": true,
		}.Merge(p.ClassNames),
		htmx.SVG(
			htmx.ClassNames{
				"stroke-warning": true,
				"shrink-0":       true,
				"w-6":            true,
				"h-6":            true,
			},
			htmx.SVG(
				htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				htmx.Attribute("fill", "none"),
				htmx.Attribute("viewBox", "0 0 24 24"),
				htmx.Path(
					htmx.Attribute("stroke-linecap", "round"),
					htmx.Attribute("stroke-linejoin", "round"),
					htmx.Attribute("stroke-width", "2"),
					htmx.Attribute("d", "M19 5L5 19"),
				),
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}

// Error is a component that displays an error alert.
func Error(p AlertProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"alert":       true,
			"alert-error": true,
		}.Merge(p.ClassNames),
		htmx.SVG(
			htmx.ClassNames{
				"stroke-error": true,
				"shrink-0":     true,
				"w-6":          true,
				"h-6":          true,
			},
			htmx.SVG(
				htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				htmx.Attribute("fill", "none"),
				htmx.Attribute("viewBox", "0 0 24 24"),
				htmx.Path(
					htmx.Attribute("stroke-linecap", "round"),
					htmx.Attribute("stroke-linejoin", "round"),
					htmx.Attribute("stroke-width", "2"),
					htmx.Attribute("d", "M6 18L18 6M6 6l12 12"),
				),
			),
		),
		htmx.Span(
			htmx.Group(children...),
		),
	)
}
