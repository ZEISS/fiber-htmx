package buttons

import htmx "github.com/zeiss/fiber-htmx"

// ButtonProps represents the properties for a button element.
type ButtonProps struct {
	ClassNames map[string]bool // The class names for the button element.
	Type       string          // The type of the button element.
	Disabled   bool            // Whether the button element is disabled.
}

// Button generates a button element based on the provided properties.
func Button(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Primary generates a primary button element based on the provided properties.
func Primary(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":         true,
				"btn-primary": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Neutral generates a neutral button element based on the provided properties.
func Neutral(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":         true,
				"btn-neutral": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Secondary generates a secondary button element based on the provided properties.
func Secondary(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":           true,
				"btn-secondary": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Accent generates an accent button element based on the provided properties.
func Accent(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":        true,
				"btn-accent": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Ghost generates a ghost button element based on the provided properties.
func Ghost(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":       true,
				"btn-ghost": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Link generates a link button element based on the provided properties.
func Link(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":      true,
				"btn-link": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Info generates an info button element based on the provided properties.
func Info(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":      true,
				"btn-info": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Success generates a success button element based on the provided properties.
func Success(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":         true,
				"btn-success": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Warning generates a warning button element based on the provided properties.
func Warning(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":         true,
				"btn-warning": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Error generates an error button element based on the provided properties.
func Error(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":       true,
				"btn-error": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Outline generates an outline button element based on the provided properties.
func Outline(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn":         true,
				"btn-outline": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// OutlinePrimary generates an outline primary button element based on the provided properties.
func OutlinePrimary(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn-outline": true,
				"btn-primary": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// OutlineSecondary generates an outline secondary button element based on the provided properties.
func OutlineSecondary(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn-outline":   true,
				"btn-secondary": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// OutlineAccent generates an outline accent button element based on the provided properties.
func OutlineAccent(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn-outline": true,
				"btn-accent":  true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// OutlineInfo generates an outline info button element based on the provided properties.
func OutlineInfo(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn-outline": true,
				"btn-info":    true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// OutlineSuccess generates an outline success button element based on the provided properties.
func OutlineSuccess(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn-outline": true,
				"btn-success": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// OutlineWarning generates an outline warning button element based on the provided properties.
func OutlineWarning(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn-outline": true,
				"btn-warning": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// OutlineError generates an outline error button element based on the provided properties.
func OutlineError(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn-outline": true,
				"btn-error":   true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Glass generates a glass button element based on the provided properties.
func Glass(props ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn-outline": true,
				"glass":       true,
			},
			props.ClassNames,
		),
		htmx.Attribute("type", props.Type),
		htmx.If(props.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}
