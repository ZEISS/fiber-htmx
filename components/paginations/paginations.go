package paginations

import (
	"net/url"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/pkg/urlx"
	"github.com/zeiss/pkg/utilx"
)

// DefaultLimits is a slice of integers.
var DefaultLimits = []int{10, 25, 50, 100}

// PaginationProps is a struct that contains the properties of a pagination
type PaginationProps struct {
	// ClassNames is a struct that contains the class names of a pagination.
	ClassNames htmx.ClassNames
	// ID is the id of the pagination.
	ID string
	// Total is the total number of items.
	Total int
	// Offset is the number of items to skip.
	Offset int
	// Limit is the number of items to return.
	Limit int
	// URL is the URL of the pagination.
	URL string
}

// Pagination is a component that renders a pagination.
func Pagination(props PaginationProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Prev is a component that renders a previous button.
func Prev(props PaginationProps) htmx.Node {
	values := url.Values{
		"offset": {conv.String(props.Offset - props.Limit)},
		"limit":  {conv.String(props.Limit)},
	}

	props.URL = urlx.MustCopyValues(props.URL, values)

	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"join-item":   true,
				"btn":         true,
				"btn-outline": true,
				"disabled":    props.Offset-props.Limit < 0,
			},
			props.ClassNames,
		),
		htmx.HxGet(props.URL),
		htmx.IfElse(utilx.NotEmpty(props.ID), htmx.HxTarget(props.ID), htmx.HxTarget("#data-table")),
		htmx.HxSwap("innerHTML"),
		htmx.Text("Prev"),
	)
}

// Next is a component that renders a next button.
func Next(props PaginationProps) htmx.Node {
	values := url.Values{
		"offset": {conv.String(props.Offset + props.Limit)},
		"limit":  {conv.String(props.Limit)},
	}

	props.URL = urlx.MustCopyValues(props.URL, values)

	return htmx.A(
		htmx.Merge(
			htmx.ClassNames{
				"btn-outline": true,
				"btn":         true,
				"join-item":   true,
			},
			props.ClassNames,
		),
		htmx.HxGet(props.URL),
		htmx.HxSwap("innerHTML"),
		htmx.IfElse(utilx.NotEmpty(props.ID), htmx.HxTarget(props.ID), htmx.HxTarget("#data-table")),
		htmx.Text("Next"),
	)
}

// SelectProps is a component that contains the properties of a select.
type SelectProps struct {
	// ClassNames is a struct that contains the class names of a select.
	ClassNames htmx.ClassNames
	// Limits is a slice of integers to select from.
	Limits []int
}

// Select ...
func Select(props SelectProps, children ...htmx.Node) htmx.Node {
	return forms.Select(
		forms.SelectProps{
			ClassNames: htmx.Merge(props.ClassNames),
		},
		htmx.ForEach(props.Limits, func(limit int, _ int) htmx.Node {
			return forms.Option(
				forms.OptionProps{},
				htmx.Text(conv.String(limit)),
			)
		})...,
	)
}
