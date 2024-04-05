// Package icons provides a set of icons that can be used in the application.
// The icons are implemented as functions that return htmx.Node.
//
// https://heroicons.com
package icons

import (
	htmx "github.com/zeiss/fiber-htmx"
)

// IconProps ...
type IconProps struct {
	ClassNames htmx.ClassNames
}

// ChevronUpDownOutline ...
func ChevronUpDownOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M8.25 15 12 18.75 15.75 15m-7.5-6L12 5.25 15.75 9"),
		),
	)
}

// ChevronUpOutline ...
func ChevronUpOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M19.5 15 12 7.5 4.5 15"),
		),
	)
}

// ChevronDownOutline ...
func ChevronDownOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M19.5 8.25 12 15.75 4.5 8.25"),
		),
	)
}

// ChevronLeftOutline ...
func ChevronLeftOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M15 19.5 7.5 12 15 4.5"),
		),
	)
}

// ChevronRightOutline ...
func ChevronRightOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.ClassNames{
			"w-6 h-6": true,
		}.Merge(p.ClassNames),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M9 19.5 16.5 12 9 4.5"),
		),
	)
}

// SearchOutline ...
func SearchOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"),
		),
	)
}

// BoltOutline ...
func BoltOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m3.75 13.5 10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75Z"),
		),
	)
}

// BoltSlashOutline ...
func BoltSlashOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M11.412 15.655 9.75 21.75l3.745-4.012M9.257 13.5H3.75l2.659-2.849m2.048-2.194L14.25 2.25 12 10.5h8.25l-4.707 5.043M8.457 8.457 3 3m5.457 5.457 7.086 7.086m0 0L21 21"),
		),
	)
}

// HeartOutline ...
func HeartOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z"),
		),
	)
}

// Bars2Outline ...
func Bars2Outline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M3.75 9h16.5m-16.5 6.75h16.5"),
		),
	)
}

// Bars3Outline ...
func Bars3Outline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M3.75 5.25h16.5m-16.5 6.75h16.5m-16.5 6.75h16.5"),
		),
	)
}

// UserOutline ...
func UserOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z"),
		),
	)
}

// TrashOutline ...
func TrashOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0"),
		),
	)
}

// BriefcaseOutline ...
func BriefcaseOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M20.25 14.15v4.25c0 1.094-.787 2.036-1.872 2.18-2.087.277-4.216.42-6.378.42s-4.291-.143-6.378-.42c-1.085-.144-1.872-1.086-1.872-2.18v-4.25m16.5 0a2.18 2.18 0 0 0 .75-1.661V8.706c0-1.081-.768-2.015-1.837-2.175a48.114 48.114 0 0 0-3.413-.387m4.5 8.006c-.194.165-.42.295-.673.38A23.978 23.978 0 0 1 12 15.75c-2.648 0-5.195-.429-7.577-1.22a2.016 2.016 0 0 1-.673-.38m0 0A2.18 2.18 0 0 1 3 12.489V8.706c0-1.081.768-2.015 1.837-2.175a48.111 48.111 0 0 1 3.413-.387m7.5 0V5.25A2.25 2.25 0 0 0 13.5 3h-3a2.25 2.25 0 0 0-2.25 2.25v.894m7.5 0a48.667 48.667 0 0 0-7.5 0M12 12.75h.008v.008H12v-.008Z"),
		),
	)
}

// DocumentMagnifyingGlassOutline ...
func DocumentMagnifyingGlassOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m5.231 13.481L15 17.25m-4.5-15H5.625c-.621 0-1.125.504-1.125 1.125v16.5c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9Zm3.75 11.625a2.625 2.625 0 1 1-5.25 0 2.625 2.625 0 0 1 5.25 0Z"),
		),
	)
}

// MagnifyingGlassOutline ...
func MagnifyingGlassOutline(p IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("fill", "none"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Attribute("stroke-width", "1.5"),
		htmx.Attribute("stroke", "currentColor"),
		htmx.Merge(
			htmx.ClassNames{
				"w-6 h-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"),
		),
	)
}
