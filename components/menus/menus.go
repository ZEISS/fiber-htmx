package menus

import htmx "github.com/zeiss/fiber-htmx"

// MenuProps is the struct for the menu props
type MenuProps struct {
	ClassNames htmx.ClassNames
}

// Menu is the component for the menu
func Menu(p MenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.Merge(
			htmx.ClassNames{
				"menu":        true,
				"bg-base-200": true,
				"w-56":        true,
				"rounded-box": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// MenuItemProps is the struct for the menu item props
type MenuItemProps struct {
	ClassNames htmx.ClassNames
}

// MenuItem is the component for the menu item
func MenuItem(p MenuItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{
				"menu-item":         true,
				"hover:bg-base-300": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
