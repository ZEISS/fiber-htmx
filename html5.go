package htmx

import (
	"io"
	"sort"
	"strings"
)

// HTML5Props ...
type HTML5Props struct {
	Title       string
	Description string
	Language    string
	Head        []Node
	Body        []Node
}

// HTML5 ...
func HTML5(p HTML5Props) Node {
	return Doctype(
		HTML(If(p.Language != "", Lang(p.Language)),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				TitleElement(Text(p.Title)),
				If(p.Description != "", Meta(Name("description"), Content(p.Description))),
				Group(p.Head...),
			),
			Body(Group(p.Body...)),
		),
	)
}

// HxClassName ...
type HxClassName string

// String ...
func (c HxClassName) String() string {
	return string(c)
}

const (
	HxClassNameAdded     HxClassName = "htmx-added"
	HxClassNameIndicator HxClassName = "htmx-indicator"
	HxClassNameRequest   HxClassName = "htmx-request"
	HxClassNameSettling  HxClassName = "htmx-settling"
	HxClassNameSwapping  HxClassName = "htmx-swapping"
)

// ClassNames ...
type ClassNames map[string]bool

// Render ...
func (c ClassNames) Render(w io.Writer) error {
	var included []string
	for c, include := range c {
		if include {
			included = append(included, c)
		}
	}

	sort.Strings(included)

	return Class(strings.Join(included, " ")).Render(w)
}

// Type ...
func (c ClassNames) Type() NodeType {
	return AttributeType
}

// String ...
func (c ClassNames) String() string {
	var b strings.Builder

	_ = c.Render(&b)

	return b.String()
}
