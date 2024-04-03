package utils

import htmx "github.com/zeiss/fiber-htmx"

// Transformer is a function that transforms a node.
type Transformer[T any] func(el T) htmx.Node

// ForEach loops over the content.
func ForEach[T any](transformer Transformer[T], elems ...T) htmx.Node {
	nodes := []htmx.Node{}

	for _, n := range elems {
		nodes = append(nodes, transformer(n))
	}

	return htmx.Group(nodes...)
}
