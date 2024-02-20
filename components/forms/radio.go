package forms

import htmx "github.com/zeiss/fiber-htmx"

// RadioProps represents the properties for a radio element.
type RadioProps struct {
	ClassNames map[string]bool // The class names for the radio element.
	Name       string          // The name of the radio element.
	Value      string          // The value of the radio element.
	Checked    bool            // Whether the radio element is checked.
}

// Radio generates a radio element based on the provided properties.
func Radio(p RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.ClassNames{
			"radio": true,
		}.Merge(p.ClassNames),
		htmx.Attribute("type", "radio"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.If(p.Checked, htmx.Checked()),
		htmx.Group(children...),
	)
}
