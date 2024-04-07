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
	Attributes  []Node // The attributes to be included in the HTML document.
}

// HTML5 generates an HTML5 document based on the provided properties.
func HTML5(ctx Ctx, props HTML5Props, body ...Node) Node {
	return Doctype(
		HTML(
			If(props.Language != "", Lang(props.Language)),
			Group(props.Attributes...),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				TitleElement(Text(props.Title)),
				If(props.Description != "", Meta(Name("description"), Content(props.Description))),
				Group(props.Head...),
			),
			Body(body...),
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

// Merge merges the provided class names into the class names.
func (c ClassNames) Merge(classNames ClassNames) ClassNames {
	for className, include := range classNames {
		c[className] = include
	}

	return c
}

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
