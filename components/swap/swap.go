package swap

import htmx "github.com/zeiss/fiber-htmx"

// SwapProps contains the properties for the swap component.
type SwapProps struct {
	ClassNames htmx.ClassNames
}

// SwapOn is a component for the on state of the swap component.
func SwapOn(p SwapProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"swap-on": true,
			},
			p.ClassNames,
		),
	)
}

// SwapOff is a component for the off state of the swap component.
func SwapOff(p SwapProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"swap-off": true,
			},
			p.ClassNames,
		),
	)
}

// Swap is a component for the htmx swap extension.
func Swap(p SwapProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"swap": true,
			},
			p.ClassNames,
		),
		htmx.Input(
			htmx.Attribute("type", "checkbox"),
		),
		htmx.Group(children...),
	)
}
