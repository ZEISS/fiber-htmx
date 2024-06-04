package forms

import htmx "github.com/zeiss/fiber-htmx"

// CheckboxProps represents the properties for a checkbox element.
type CheckboxProps struct {
	ClassNames htmx.ClassNames // The class names for the checkbox element.
	Name       string          // The name of the checkbox element.
	Value      string          // The value of the checkbox element.
	Checked    bool            // Whether the checkbox element is checked.
	Disabled   bool            // Whether the checkbox element is disabled.
}

// Checkbox generates a checkbox element based on the provided properties.
func Checkbox(p CheckboxProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"checkbox": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "checkbox"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// CheckboxPrimary is a component that displays a primary checkbox.
func CheckboxPrimary(p CheckboxProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"checkbox-primary": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Checkbox(p, children...)
}

// CheckboxSuccess is a component that displays a success checkbox.
func CheckboxSuccess(p CheckboxProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"checkbox-success": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Checkbox(p, children...)
}

// CheckboxWarning is a component that displays a warning checkbox.
func CheckboxWarning(p CheckboxProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"checkbox-warning": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Checkbox(p, children...)
}

// CheckboxError is a component that displays an error checkbox.
func CheckboxError(p CheckboxProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"checkbox-error": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Checkbox(p, children...)
}

// CheckboxInfo is a component that displays an info checkbox.
func CheckboxInfo(p CheckboxProps, children ...htmx.Node) htmx.Node {
	classNames := htmx.ClassNames{
		"checkbox-info": true,
	}.Merge(p.ClassNames)

	p.ClassNames = classNames

	return Checkbox(p, children...)
}
