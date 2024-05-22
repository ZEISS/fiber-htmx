package paginations

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
)

var DefaultLimits = []int{10, 25, 50, 100}

// PaginationProps is a struct that contains the properties of a pagination
type PaginationProps struct {
	ClassNames htmx.ClassNames
	Total      int
	Offset     int
	Limit      int
	URL        string
}

// Pagination ...
func Pagination(p PaginationProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Prev ...
func Prev(p PaginationProps) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"join-item":   true,
			"btn":         true,
			"btn-outline": true,
			"disabled":    p.Offset-p.Limit < 0,
		},
		htmx.HxGet(fmt.Sprintf("%s?offset=%d&limit=%d", p.URL, p.Offset-p.Limit, p.Limit)),
		htmx.HxSwap("innerHTML"),
		htmx.HxTarget("#data-table"),
		htmx.Text("Prev"),
	)
}

// Next ...
func Next(p PaginationProps) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"join-item":   true,
			"btn":         true,
			"btn-outline": true,
		},
		htmx.HxGet(fmt.Sprintf("%s?offset=%d&limit=%d", p.URL, p.Offset+p.Limit, p.Limit)),
		htmx.HxSwap("innerHTML"),
		htmx.HxTarget("#data-table"),
		htmx.Text("Next"),
	)
}

// SelectProps ...
type SelectProps struct {
	ClassNames htmx.ClassNames
	Limits     []int
}

// Select ...
func Select(p SelectProps, children ...htmx.Node) htmx.Node {
	options := []htmx.Node{}
	for _, limit := range p.Limits {
		options = append(options, forms.Option(
			forms.OptionProps{},
			htmx.Text(fmt.Sprintf("%d", limit)),
		))
	}

	return forms.Select(
		forms.SelectProps{
			ClassNames: htmx.Merge(p.ClassNames),
		},
		options...,
	)
}
