package tables

import (
	"fmt"

	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/utils"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/zeiss/fiber-htmx"
)

// DefaultLimits is a list of default limits
var DefaultLimits = []int{5, 10, 25, 50}

// PaginationProps is a struct that contains the properties of a pagination
type PaginationProps struct {
	ClassName htmx.ClassNames
	Limit     int
	Offset    int
	Target    string
	Total     int
	URL       string
}

// Pagination ...
func Pagination(p PaginationProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join": true,
			},
			p.ClassName,
		),
		htmx.Group(children...),
	)
}

// Prev ...
func Prev(p PaginationProps) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(p.URL),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("offset"),
			htmx.Value(fmt.Sprintf("%d", p.Offset-p.Limit)),
		),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("limit"),
			htmx.Value(fmt.Sprintf("%d", p.Limit)),
		),
		htmx.HxBoost(true),
		buttons.Button(
			buttons.ButtonProps{
				ClassNames: htmx.ClassNames{
					"join-item":      true,
					"btn":            true,
					"btn-outline":    true,
					"input-bordered": true,
				},
				Type: "submit",
			},
			htmx.If(p.Offset-p.Limit < 0, htmx.Disabled()),
			htmx.Text("Prev"),
		),
	)
}

// Next ...
func Next(p PaginationProps) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(p.URL),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("offset"),
			htmx.Value(fmt.Sprintf("%d", p.Offset+p.Limit)),
		),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("limit"),
			htmx.Value(fmt.Sprintf("%d", p.Limit)),
		),
		htmx.HxBoost(true),
		buttons.Button(
			buttons.ButtonProps{
				ClassNames: htmx.ClassNames{
					"join-item":      true,
					"btn":            true,
					"btn-outline":    true,
					"input-bordered": true,
				},
				Type: "submit",
			},
			htmx.If(p.Offset+p.Limit > p.Total, htmx.Disabled()),
			htmx.Text("Next"),
		),
	)
}

// SelectProps ...
type SelectProps struct {
	ClassNames htmx.ClassNames
	Limit      int
	Limits     []int
	Offset     int
	Target     string
	Total      int
	URL        string
}

// Select ...
func Select(p SelectProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(p.URL),
		htmx.HxTrigger("change from:#select-table-options"),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("offset"),
			htmx.Value(fmt.Sprintf("%d", p.Offset)),
		),
		htmx.HxBoost(true),
		forms.Select(
			forms.SelectProps{
				ClassName: htmx.Merge(
					htmx.ClassNames{
						"join-item":      true,
						"input-bordered": true,
					},
					p.ClassNames,
				),
			},
			htmx.ID("select-table-options"),
			htmx.Attribute("name", "limit"),
			utils.Map(func(limit int) htmx.Node {
				return forms.Option(
					forms.OptionProps{
						Selected: limit == p.Limit,
					},
					htmx.Text(fmt.Sprintf("%d", limit)),
					htmx.Value(fmt.Sprintf("%d", limit)),
				)
			}, p.Limits...),
		),
	)
}

// TableToolbarProps is a struct that contains the properties of a table toolbar
type TableToolbarProps[R comparable] struct {
	ClassName htmx.ClassNames
}

// TableToolbar is a component that renders a table toolbar
func TableToolbar[R comparable](p TableToolbarProps[R], children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"table-toolbar": true,
			},
			p.ClassName,
		),
		htmx.Group(children...),
	)
}

// TablePaginationProps is a struct that contains the properties of a table pagination
type TablePaginationProps[R comparable] struct {
	ClassName  htmx.ClassNames
	Pagination htmx.Node
}

// TablePagination is a component that renders a table pagination
func TablePagination[R comparable](p TablePaginationProps[R], children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"flex":            true,
				"items-center":    true,
				"justify-between": true,
				"px-2":            true,
			},
			p.ClassName,
		),
		htmx.Div(
			htmx.Merge(
				htmx.ClassNames{
					"flex":         true,
					"items-center": true,
					"space-x-6":    true,
					"lg:space-x-8": true,
				},
			),
			htmx.P(
				htmx.Merge(
					htmx.ClassNames{
						"text-sm":     true,
						"font-medium": true,
					},
				),
				htmx.Text("Rows per page"),
			),
			p.Pagination,
		),
	)
}

// TableProps is a struct that contains the properties of a table
type TableProps[R comparable] struct {
	ClassName  htmx.ClassNames
	Columns    Columns[R]
	ID         string
	Pagination htmx.Node
	Rows       Rows[R]
	Toolbar    htmx.Node
}

// Rows is a struct that contains the data of a table
type Rows[R comparable] struct {
	Data []R
}

// NewRows returns a new Rows object.
func NewRows[R comparable](data []R) Rows[R] {
	return Rows[R]{
		Data: data,
	}
}

// Insert inserts a new row.
func (r *Rows[R]) Insert(data R) {
	r.Data = append(r.Data, data)
}

// ValueByIndex is a helper function that returns the value of a row based on the provided index.
func (r *Rows[R]) ValueByIndex(index int) R {
	if index >= len(r.Data) {
		panic("Index out of range")
	}

	return r.Data[index]
}

// GetAll returns all the rows.
func (r *Rows[T]) GetAll() []T {
	return r.Data
}

// Columns returns a new column definition.
type Columns[R comparable] []ColumnDef[R]

// ColumnDef returns a new column definition.
type ColumnDef[R comparable] struct {
	ID              string
	AccessorKey     string
	Header          func(p TableProps[R]) htmx.Node
	Cell            func(p TableProps[R], row R) htmx.Node
	EnableSorting   bool
	EnableFiltering bool
}

// Table is a struct that contains the properties of a table
func Table[R comparable](p TableProps[R], children ...htmx.Node) htmx.Node {
	headers := []htmx.Node{}
	for _, column := range p.Columns {
		headers = append(headers, column.Header(p))
	}

	rows := []htmx.Node{}
	for _, row := range p.Rows.Data {
		cells := []htmx.Node{}
		for _, column := range p.Columns {
			cells = append(cells, column.Cell(p, row))
		}
		rows = append(rows, htmx.Tr(cells...))

	}

	return htmx.Div(
		htmx.ID(p.ID),
		htmx.Merge(
			htmx.ClassNames{
				"space-y-4": true,
			},
		),
		htmx.Div(
			htmx.Merge(),
			htmx.Table(
				htmx.Merge(
					htmx.ClassNames{
						"table": true,
					},
					p.ClassName,
				),
				htmx.Group(children...),
				htmx.THead(
					htmx.Tr(
						headers...,
					),
				),
				htmx.TBody(
					rows...,
				),
			),
		),
		p.Pagination,
	)
}

// PaginationFromContext returns a new Pagination object based on the provided context.
func PaginationPropsFromContext(c *fiber.Ctx) (limit int, offset int) {
	limit = c.QueryInt("limit", 10)
	offset = c.QueryInt("offset", 0)

	return
}
