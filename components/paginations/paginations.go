package paginations

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/selects"
)

var DefaultLimits = []int{10, 25, 50, 100}

// PaginationProps is a struct that contains the properties of a pagination
type PaginationProps struct {
	ClassName htmx.ClassNames
	Total     int
	Offset    int
	Limit     int
	URL       string
}

// Pagination ...
func Pagination(p PaginationProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"join": true,
		}.Merge(p.ClassName),
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
		htmx.HxGet(fmt.Sprintf("%s?offset=%d&limit=%d", p.URL, p.Offset-p.Limit, p.Limit)),
		htmx.HxSwap("innerHTML"),
		htmx.HxTarget("#data-table"),
		htmx.Text("Next"),
	)
}

// SelectProps ...
type SelectProps struct {
	ClassName htmx.ClassNames
	Limits    []int
}

// Select ...
func Select(p SelectProps, children ...htmx.Node) htmx.Node {
	options := []htmx.Node{}
	for _, limit := range p.Limits {
		options = append(options, selects.Option(
			selects.OptionProps{},
			htmx.Text(fmt.Sprintf("%d", limit)),
		))
	}

	return selects.Select(
		selects.SelectProps{
			ClassName: htmx.ClassNames{}.Merge(p.ClassName),
		},
		options...,
	)
}
