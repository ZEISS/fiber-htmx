package htmx

import (
	"io"
	"sort"
	"strings"
)

// HTML5Props represents the properties for an HTML5 document.
type HTML5Props struct {
	Title       string // The title of the HTML document.
	Description string // The description of the HTML document.
	Language    string // The language of the HTML document.
	Head        []Node // The nodes to be included in the head section of the HTML document.
	Body        []Node // The nodes to be included in the body section of the HTML document.
}

// HTML5 generates an HTML5 document based on the provided properties.
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

// HxClassName represents a class name for htmx elements.
type HxClassName string

// String returns the string representation of the HxClassName.
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

// ClassNames represents a set of class names.
type ClassNames map[string]bool

// Render writes the class names to the provided writer.
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

// Type returns the node type of the ClassNames.
func (c ClassNames) Type() NodeType {
	return AttributeType
}

// String returns the string representation of the ClassNames.
func (c ClassNames) String() string {
	var b strings.Builder

	_ = c.Render(&b)

	return b.String()
}
