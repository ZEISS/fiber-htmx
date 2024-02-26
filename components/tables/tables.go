package tables

import htmx "github.com/zeiss/fiber-htmx"

// TableProps is a struct that contains the properties of a table
type TableProps[R comparable] struct {
	ClassName htmx.ClassNames
	Columns   Columns[R]
	Rows      Rows[R]
}

// Rows ...
type Rows[R comparable] struct {
	Data []R
}

// NewRows ...
func NewRows[R comparable](data []R) Rows[R] {
	return Rows[R]{
		Data: data,
	}
}

// Insert ...
func (r *Rows[R]) Insert(data R) {
	r.Data = append(r.Data, data)
}

// ValueByIndex ...
func (r *Rows[R]) ValueByIndex(index int) R {
	if index >= len(r.Data) {
		panic("Index out of range")
	}

	return r.Data[index]
}

// GetAll ...
func (r *Rows[T]) GetAll() []T {
	return r.Data
}

// Columns ...
type Columns[R comparable] []ColumnDef[R]

// ColumnDef ...
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

	return htmx.Table(
		htmx.ClassNames{
			"table": true,
		}.Merge(p.ClassName),
		htmx.THead(
			htmx.Tr(
				headers...,
			),
		),
		htmx.TBody(
			rows...,
		),
	)
}
