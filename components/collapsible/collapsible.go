package collapsible

import htmx "github.com/zeiss/fiber-htmx"

// CollapseProps is a component that can be expanded and collapsed.
type CollapseProps struct {
	ClassNames htmx.ClassNames
	TabIndex   int
}

// Collapse is a component that can be expanded and collapsed.
func Collapse(props CollapseProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Attribute("tabindex", htmx.IntAsString(props.TabIndex)),
		htmx.ClassNames{
			"collapse":    true,
			"bg-base-200": true,
		}.Merge(props.ClassNames),
		htmx.Group(children...),
	)
}

// CollapseArrow is a component that can be expanded and collapsed.
func CollapseArrow(p CollapseProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Attribute("tabindex", htmx.IntAsString(p.TabIndex)),
		htmx.Merge(
			htmx.ClassNames{
				"collapse":        true,
				"collapse-arrow":  true,
				"bg-base-200":     true,
				"border":          true,
				"border-base-300": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// CollapseTitleProps is a component that can be expanded and collapsed.
type CollapseTitleProps struct {
	ClassNames htmx.ClassNames
}

// CollapseTitle is a component that can be expanded and collapsed.
func CollapseTitle(p CollapseTitleProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse-title": true,
				"title-md":       true,
				"font-medium":    true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// CollapseContentProps is a component that can be expanded and collapsed.
type CollapseContentProps struct {
	ClassNames htmx.ClassNames
}

// CollapseContent is a component that can be expanded and collapsed.
func CollapseContent(props CollapseContentProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"collapse-content": true,
		}.Merge(props.ClassNames),
		htmx.Group(children...),
	)
}

// CollapseCheckboxProps is a component that can be expanded and collapsed.
type CollapseCheckboxProps struct {
	ClassNames htmx.ClassNames
	Checked    bool
}

// CollapseCheckbox is a component that can be expanded and collapsed.
func CollapseCheckbox(props CollapseCheckboxProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("checkbox"),
	)
}
