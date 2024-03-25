package htmx

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

// Node is a node in the HTML tree.
type Node interface {
	Render(w io.Writer) error
}

// NodeType is the type of a node.
type NodeType int

const (
	ElementType = NodeType(iota)
	AttributeType
)

// NodeTypeDescriptor is a node that has a type.
type NodeTypeDescriptor interface {
	Type() NodeType
}

// NodeFunc is a function that renders a node.
type NodeFunc func(io.Writer) error

// Render renders the node.
func (n NodeFunc) Render(w io.Writer) error {
	return n(w)
}

// Type returns the node type.
func (n NodeFunc) Type() NodeType {
	return ElementType
}

// String returns the node as a string.
func (n NodeFunc) String() string {
	var b strings.Builder

	_ = n.Render(&b)

	return b.String()
}

// Element is a node that renders an HTML element.
func Element(name string, children ...Node) NodeFunc {
	return NodeFunc(func(w2 io.Writer) error {
		w := &statefulWriter{w: w2}

		w.Write([]byte("<" + name))

		for _, c := range children {
			renderChild(w, c, AttributeType)
		}

		w.Write([]byte(">"))

		if isVoidElement(name) {
			return w.err
		}

		for _, c := range children {
			renderChild(w, c, ElementType)
		}

		w.Write([]byte("</" + name + ">"))

		return w.err
	})
}

//nolint:gocyclo
func renderChild(w *statefulWriter, c Node, t NodeType) {
	if w.err != nil || c == nil {
		return
	}

	if g, ok := c.(group); ok {
		for _, groupC := range g.children {
			renderChild(w, groupC, t)
		}
		return
	}

	switch t {
	case ElementType:
		if p, ok := c.(NodeTypeDescriptor); !ok || p.Type() == ElementType {
			w.err = c.Render(w.w)
		}
	case AttributeType:
		if p, ok := c.(NodeTypeDescriptor); ok && p.Type() == AttributeType {
			w.err = c.Render(w.w)
		}
	}
}

type statefulWriter struct {
	w   io.Writer
	err error
}

// Write is a node that writes to the stateful writer.
func (w *statefulWriter) Write(p []byte) {
	if w.err != nil {
		return
	}
	_, w.err = w.w.Write(p)
}

var voidElements = map[string]struct{}{
	"area":    {},
	"base":    {},
	"br":      {},
	"col":     {},
	"command": {},
	"embed":   {},
	"hr":      {},
	"img":     {},
	"input":   {},
	"keygen":  {},
	"link":    {},
	"meta":    {},
	"param":   {},
	"source":  {},
	"track":   {},
	"wbr":     {},
}

func isVoidElement(name string) bool {
	_, ok := voidElements[name]
	return ok
}

// Attribute is a node that renders an HTML attribute.
func Attribute(name string, value ...string) Node {
	switch len(value) {
	case 0:
		return &attr{name: name}
	case 1:
		return &attr{name: name, value: &value[0]}
	default:
		panic("attribute must be just name or name and value pair")
	}
}

type attr struct {
	name  string
	value *string
}

// Render is a node that renders an attribute.
func (a *attr) Render(w io.Writer) error {
	if a.value == nil {
		_, err := w.Write([]byte(" " + a.name))
		return err
	}

	_, err := w.Write([]byte(" " + a.name + `="` + template.HTMLEscapeString(*a.value) + `"`))

	return err
}

// Type is a node that returns the type of an attribute.
func (a *attr) Type() NodeType {
	return AttributeType
}

// String is a node that returns the attribute as a string.
func (a *attr) String() string {
	var b strings.Builder

	_ = a.Render(&b)

	return b.String()
}

// Text is a node that renders a text.
func Text(t string) Node {
	return NodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(template.HTMLEscapeString(t)))

		return err
	})
}

// Textf is a node that renders a formatted text.
func Textf(format string, a ...interface{}) Node {
	return NodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(template.HTMLEscapeString(fmt.Sprintf(format, a...))))

		return err
	})
}

// Raw is a node that renders raw HTML.
func Raw(t string) Node {
	return NodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(t))

		return err
	})
}

// Rawf is a node that renders a formatted raw HTML.
func Rawf(format string, a ...interface{}) Node {
	return NodeFunc(func(w io.Writer) error {
		_, err := fmt.Fprintf(w, format, a...)

		return err
	})
}

type group struct {
	children []Node
}

// String is a node that renders a group of nodes.
func (c group) String() string {
	panic("cannot render group directly")
}

// Render is a node that renders a group of nodes.
func (c group) Render(io.Writer) error {
	panic("cannot render children directly")
}

// Group is a node that groups children nodes.
func Group(children ...Node) Node {
	return group{children: children}
}

// If is a node that renders a child node if a condition is true.
func If(condition bool, n Node) Node {
	if condition {
		return n
	}

	return nil
}
