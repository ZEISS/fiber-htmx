package tabs

import (
	htmx "github.com/zeiss/fiber-htmx"
)

// TabsProps is a struct that contains the properties of the Tabs component
type TabsProps struct {
	ClassNames htmx.ClassNames
}

// Tabs is a component that renders a list of tabs
func Tabs(props TabsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tabs": true,
			},
			props.ClassNames,
		),
		htmx.Role("tablist"),
		htmx.Group(children...),
	)
}

// TabsBoxed is a component that renders a list of tabs in a box
func TabsBoxed(props TabsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tabs":       true,
				"tabs-boxed": true,
			},
			props.ClassNames,
		),
		htmx.Role("tablist"),
		htmx.Group(children...),
	)
}

// TabProps is a struct that contains the properties of the Tab component
type TabProps struct {
	ClassNames htmx.ClassNames
	Active     bool
}

// Tab is a component that renders a tab.
func Tab(props TabProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"tab":        true,
				"tab-active": props.Active,
			},
			props.ClassNames,
		),
		htmx.Role("tab"),
		htmx.Group(children...),
	)
}
