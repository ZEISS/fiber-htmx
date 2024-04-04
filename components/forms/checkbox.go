package forms

import htmx "github.com/zeiss/fiber-htmx"

// FormControlProps represents the properties for a form control element.
type FormControlProps struct {
	ClassNames htmx.ClassNames // The class names for the form control element.
}

// FormControl generates a form control element based on the provided properties.
func FormControl(p FormControlProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"form-control": true,
				"w-full":       true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FormControlLabelProps represents the properties for a form control label element.
type FormControlLabelProps struct {
	ClassNames htmx.ClassNames // The class names for the form control label element.
}

// FormControlLabel generates a form control label element based on the provided properties.
func FormControlLabel(props FormControlLabelProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"label": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FormControlLabelTextProps represents the properties for a form control label text element.
type FormControlLabelTextProps struct {
	ClassNames htmx.ClassNames // The class names for the form control label text element.
}

// FormControlLabelText generates a form control label text element based on the provided properties.
func FormControlLabelText(props FormControlLabelTextProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"label-text": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FormControlLabelAltTextProps represents the properties for a form control label alt text element.
type FormControlLabelAltTextProps struct {
	ClassNames htmx.ClassNames // The class names for the form control label alt text element.
}

// FormControlLabelAltText generates a form control label alt text element based on the provided properties.
func FormControlLabelAltText(props FormControlLabelAltTextProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"label-text-all": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

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
