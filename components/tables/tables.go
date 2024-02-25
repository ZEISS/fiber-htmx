package tables

import htmx "github.com/zeiss/fiber-htmx"

// TableProps is a struct that contains the properties of a table
type TableProps struct {
	ClassName htmx.ClassNames
	Columns   Columns
}

// Columns ...
type Columns []ColumnDef

// ColumnDef ...
type ColumnDef struct {
	ID              string
	AccessorKey     string
	Header          func(p TableProps) htmx.Node
	Cell            func(p TableProps) htmx.Node
	EnableSorting   bool
	EnableFiltering bool
}

// Table is a struct that contains the properties of a table
func Table(p TableProps, children ...htmx.Node) htmx.Node {
	headers := []htmx.Node{}
	for _, column := range p.Columns {
		headers = append(headers, column.Header(p))
	}

	return htmx.Table(
		htmx.ClassNames{
			"table": true,
		}.Merge(p.ClassName),
		htmx.THead(
			htmx.Tr(
				headers...,
			),
		),
	)
}
