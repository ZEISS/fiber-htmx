package link

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
)

// LinkProps represents the properties for a link element.
type LinkProps struct {
	Rel        string          // The relationship between the current document and the linked document.
	Href       string          // The URL of the linked document.
	ClassNames map[string]bool // The class names for the link element.
	Children   []htmx.Node

	ctx context.Context
}

// WithContext returns the LinkProps with the provided context.
func (p LinkProps) WithContext(ctx context.Context) LinkProps {
	p.ctx = ctx

	return p
}

// Link generates a link element based on the provided properties.
func Link(p LinkProps) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(p.Children...),
	)
}

// Primary generates a primary link element based on the provided properties.
func Primary(p LinkProps) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link":         true,
			"link-primary": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(p.Children...),
	)
}
