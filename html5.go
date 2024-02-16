package htmx

import (
	"io"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var _ Props[HTML5Props] = (*HTML5Props)(nil)

// HTML5Props represents the properties for an HTML5 document.
type HTML5Props struct {
	Title       string // The title of the HTML document.
	Description string // The description of the HTML document.
	Language    string // The language of the HTML document.
	Head        []Node // The nodes to be included in the head section of the HTML document.
	Body        []Node // The nodes to be included in the body section of the HTML document.

	ctx *fiber.Ctx
}

// WithContext returns the HTML5Props with the provided context.
func (p HTML5Props) WithContext(ctx *fiber.Ctx) HTML5Props {
	p.ctx = ctx

	return p
}

// Context returns the context of the provided props.
func (p HTML5Props) Context() *fiber.Ctx {
	return p.ctx
}

// PropsWithContext returns the HTML5Props with the provided context.
type PropsWithContext[P any] interface {
	WithContext(ctx *fiber.Ctx) P
}

// PropsContext returns the context of the provided props.
type PropContext interface {
	Context() *fiber.Ctx
}

// Props is the interface for components that have properties.
type Props[P any] interface {
	PropContext
	PropsWithContext[P]
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
			Body(p.Body...),
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
