package buttons

import htmx "github.com/zeiss/fiber-htmx"

// ButtonProps represents the properties for a button element.
type ButtonProps struct {
	ClassNames map[string]bool // The class names for the button element.
	Type       string          // The type of the button element.
	Disabled   bool            // Whether the button element is disabled.
}

// Button generates a button element based on the provided properties.
func Button(p ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.ClassNames{
			"btn": true,
		}.Merge(p.ClassNames),
		htmx.Attribute("type", p.Type),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// Primary generates a primary button element based on the provided properties.
func Primary(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-primary": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Neutral generates a neutral button element based on the provided properties.
func Neutral(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-neutral": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Secondary generates a secondary button element based on the provided properties.
func Secondary(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-secondary": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Accent generates an accent button element based on the provided properties.
func Accent(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-accent": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Ghost generates a ghost button element based on the provided properties.
func Ghost(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-ghost": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Link generates a link button element based on the provided properties.
func Link(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-link": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Info generates an info button element based on the provided properties.
func Info(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-info": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Success generates a success button element based on the provided properties.
func Success(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-success": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Warning generates a warning button element based on the provided properties.
func Warning(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-warning": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Error generates an error button element based on the provided properties.
func Error(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-error": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Outline generates an outline button element based on the provided properties.
func Outline(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-outline": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// OutlinePrimary generates an outline primary button element based on the provided properties.
func OutlinePrimary(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-outline": true,
		"btn-primary": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// OutlineSecondary generates an outline secondary button element based on the provided properties.
func OutlineSecondary(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-outline":   true,
		"btn-secondary": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// OutlineAccent generates an outline accent button element based on the provided properties.
func OutlineAccent(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-outline": true,
		"btn-accent":  true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// OutlineInfo generates an outline info button element based on the provided properties.
func OutlineInfo(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-outline": true,
		"btn-info":    true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// OutlineSuccess generates an outline success button element based on the provided properties.
func OutlineSuccess(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-outline": true,
		"btn-success": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// OutlineWarning generates an outline warning button element based on the provided properties.
func OutlineWarning(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-outline": true,
		"btn-warning": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// OutlineError generates an outline error button element based on the provided properties.
func OutlineError(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn-outline": true,
		"btn-error":   true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}

// Glass generates a glass button element based on the provided properties.
func Glass(p ButtonProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"btn":   true,
		"glass": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Button(p, children...)
}
