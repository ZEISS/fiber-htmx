package forms

import htmx "github.com/zeiss/fiber-htmx"

// SelectProps ...
type SelectProps struct {
	ClassNames htmx.ClassNames
}

// Select ...
func Select(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":   true,
				"w-full":   true,
				"max-w-xs": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectBordered ...
func SelectBordered(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":          true,
				"select-bordered": true,
				"w-full":          true,
				"max-w-xs":        true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectGhost ...
func SelectGhost(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":       true,
				"select-ghost": true,
				"w-full":       true,
				"max-w-xs":     true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectPrimary ...
func SelectPrimary(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":         true,
				"select-primary": true,
				"w-full":         true,
				"max-w-xs":       true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectSecondary ...
func SelectSecondary(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":           true,
				"select-secondary": true,
				"w-full":           true,
				"max-w-xs":         true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectAccent ...
func SelectAccent(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":        true,
				"select-accent": true,
				"w-full":        true,
				"max-w-xs":      true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectInfo ...
func SelectInfo(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":      true,
				"select-info": true,
				"w-full":      true,
				"max-w-xs":    true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectSuccess ...
func SelectSuccess(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":         true,
				"select-success": true,
				"w-full":         true,
				"max-w-xs":       true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectWarning ...
func SelectWarning(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":         true,
				"select-warning": true,
				"w-full":         true,
				"max-w-xs":       true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectWarning ...
func SelectDanger(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":        true,
				"select-danger": true,
				"w-full":        true,
				"max-w-xs":      true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SelectError ...
func SelectError(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Select(
		htmx.Merge(
			htmx.ClassNames{
				"select":       true,
				"select-error": true,
				"w-full":       true,
				"max-w-xs":     true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// OptionProps ...
type OptionProps struct {
	ClassNames htmx.ClassNames
	Value      string
	Selected   bool
	Disabled   bool
}

// Option ...
func Option(p OptionProps, children ...htmx.Node) htmx.Node {
	return htmx.Option(
		htmx.Merge(p.ClassNames),
		htmx.If(p.Selected, htmx.Selected()),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.If(p.Value != "", htmx.Attribute("value", p.Value)),
		htmx.Group(children...),
	)
}
