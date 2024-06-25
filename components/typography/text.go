package typography

import htmx "github.com/zeiss/fiber-htmx"

// Error is a component that displays an error text.
func Error(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-error": true,
			},
		),
		htmx.Group(children...),
	)
}

// Info is a component that displays an info text.
func Info(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-info": true,
			},
		),
		htmx.Group(children...),
	)
}

// Success is a component that displays a success text.
func Success(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-success": true,
			},
		),
		htmx.Group(children...),
	)
}

// Warning is a component that displays a warning text.
func Warning(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-warning": true,
			},
		),
		htmx.Group(children...),
	)
}

// Primary is a component that displays a primary text.
func Primary(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-primary": true,
			},
		),
		htmx.Group(children...),
	)
}

// Secondary is a component that displays a secondary text.
func Secondary(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-secondary": true,
			},
		),
		htmx.Group(children...),
	)
}

// Accent is a component that displays an accent text.
func Accent(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-accent": true,
			},
		),
		htmx.Group(children...),
	)
}

// Neutral is a component that displays a neutral text.
func Neutral(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-neutral": true,
			},
		),
		htmx.Group(children...),
	)
}

// Base100 is a component that displays a base100 text.
func Base100(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-base-100": true,
			},
		),
		htmx.Group(children...),
	)
}

// Base200 is a component that displays a base200 text.
func Base200(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-base-200": true,
			},
		),
		htmx.Group(children...),
	)
}

// Base300 is a component that displays a base300 text.
func Base300(children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"text-base-300": true,
			},
		),
		htmx.Group(children...),
	)
}
