package dropdowns

import htmx "github.com/zeiss/fiber-htmx"

// DropdownProps represents the properties for a dropdown element.
type DropdownProps struct {
	ClassNames htmx.ClassNames // The class names for the dropdown element.
}

// Dropdown generates a dropdown element based on the provided properties.
func Dropdown(p DropdownProps, children ...htmx.Node) htmx.Node {
	return htmx.Details(
		htmx.ClassNames{
			"dropdown": true,
		}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}

// DropdownButtonProps represents the properties for a dropdown summary element.
type DropdownButtonProps struct {
	ClassNames htmx.ClassNames // The class names for the dropdown summary element.
}

// DropdownButton generates a dropdown summary element based on the provided properties.
func DropdownButton(p DropdownButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Summary(
		htmx.ClassNames{
			"m-1": true,
			"btn": true,
		}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}

// DropdownMenuItemsProps represents the properties for a dropdown menu items element.
type DropdownMenuItemsProps struct {
	ClassNames htmx.ClassNames // The class names for the dropdown menu items element.
}

// DropdownMenuItems generates a dropdown menu items element based on the provided properties.
func DropdownMenuItems(p DropdownMenuItemsProps, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.ClassNames{
			"bg-base-100":      true,
			"dropdown-content": true,
			"menu":             true,
			"p-2":              true,
			"rounded-box":      true,
			"shadow":           true,
			"w-52":             true,
			"z-[1]":            true,
		}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}

// DropdownMenuItem represents the properties for a dropdown items element.
type DropdownMenuItemProps struct {
	ClassNames htmx.ClassNames // The class names for the dropdown items element.
}

// DropdownMenuItem generates a dropdown items element based on the provided properties.
func DropdownMenuItem(p DropdownMenuItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.ClassNames{}.Merge(p.ClassNames),
		htmx.Group(children...),
	)
}
