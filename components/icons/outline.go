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
