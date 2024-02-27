package selects

import htmx "github.com/zeiss/fiber-htmx"

// SelectProps ...
type SelectProps struct {
	ClassName htmx.ClassNames
}

// Select ...
func Select(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.ClassNames{
			"select":   true,
			"w-full":   true,
			"max-w-xs": true,
		}.Merge(p.ClassName),
		htmx.Group(children...),
	)
}

// OptionProps ...
type OptionProps struct {
	ClassNames htmx.ClassNames
	Selected   bool
	Disabled   bool
}

// Option ...
func Option(p OptionProps, children ...htmx.Node) htmx.Node {
	return htmx.Option(
		htmx.ClassNames{}.Merge(p.ClassNames),
		htmx.If(p.Selected, htmx.Selected()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}
