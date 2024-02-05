package htmx

import (
	"io"
)

// Doctype ...
func Doctype(sibling Node) Node {
	return NodeFunc(func(w io.Writer) error {
		if _, err := w.Write([]byte("<!doctype html>")); err != nil {
			return err
		}

		return sibling.Render(w)
	})
}

// A ...
func A(children ...Node) Node {
	return Element("a", children...)
}

// Address ...
func Address(children ...Node) Node {
	return Element("address", children...)
}

// Area ...
func Area(children ...Node) Node {
	return Element("area", children...)
}

// Article ...
func Article(children ...Node) Node {
	return Element("article", children...)
}

// Aside
func Aside(children ...Node) Node {
	return Element("aside", children...)
}

// Audio ...
func Audio(children ...Node) Node {
	return Element("audio", children...)
}

// Base ...
func Base(children ...Node) Node {
	return Element("base", children...)
}

// BlockQuote ...
func BlockQuote(children ...Node) Node {
	return Element("blockquote", children...)
}

// Body ...
func Body(children ...Node) Node {
	return Element("body", children...)
}

// Br ...
func Br(children ...Node) Node {
	return Element("br", children...)
}

// Button ...
func Button(children ...Node) Node {
	return Element("button", children...)
}

// Canvas ...
func Canvas(children ...Node) Node {
	return Element("canvas", children...)
}

// Cite ...
func Cite(children ...Node) Node {
	return Element("cite", children...)
}

// Code ...
func Code(children ...Node) Node {
	return Element("code", children...)
}

// Col ...
func Col(children ...Node) Node {
	return Element("col", children...)
}

// ColGroup ...
func ColGroup(children ...Node) Node {
	return Element("colgroup", children...)
}

// Data ...
func DataElement(children ...Node) Node {
	return Element("data", children...)
}

// DataList ...
func DataList(children ...Node) Node {
	return Element("datalist", children...)
}

// Details ...
func Details(children ...Node) Node {
	return Element("details", children...)
}

// Dialog ...
func Dialog(children ...Node) Node {
	return Element("dialog", children...)
}

// Div ...
func Div(children ...Node) Node {
	return Element("div", children...)
}

// Dl ...
func Dl(children ...Node) Node {
	return Element("dl", children...)
}

// Embed ...
func Embed(children ...Node) Node {
	return Element("embed", children...)
}

func FormElement(children ...Node) Node {
	return Element("form", children...)
}

func FieldSet(children ...Node) Node {
	return Element("fieldset", children...)
}

func Figure(children ...Node) Node {
	return Element("figure", children...)
}

func Footer(children ...Node) Node {
	return Element("footer", children...)
}

func Head(children ...Node) Node {
	return Element("head", children...)
}

func Header(children ...Node) Node {
	return Element("header", children...)
}

func HGroup(children ...Node) Node {
	return Element("hgroup", children...)
}

func Hr(children ...Node) Node {
	return Element("hr", children...)
}

func HTML(children ...Node) Node {
	return Element("html", children...)
}

func IFrame(children ...Node) Node {
	return Element("iframe", children...)
}

func Img(children ...Node) Node {
	return Element("img", children...)
}

func Input(children ...Node) Node {
	return Element("input", children...)
}

func LabElement(children ...Node) Node {
	return Element("label", children...)
}

func Legend(children ...Node) Node {
	return Element("legend", children...)
}

func Li(children ...Node) Node {
	return Element("li", children...)
}

func Link(children ...Node) Node {
	return Element("link", children...)
}

func Main(children ...Node) Node {
	return Element("main", children...)
}

func Menu(children ...Node) Node {
	return Element("menu", children...)
}

func Meta(children ...Node) Node {
	return Element("meta", children...)
}

func Meter(children ...Node) Node {
	return Element("meter", children...)
}

func Nav(children ...Node) Node {
	return Element("nav", children...)
}

func NoScript(children ...Node) Node {
	return Element("noscript", children...)
}

func Object(children ...Node) Node {
	return Element("object", children...)
}

func Ol(children ...Node) Node {
	return Element("ol", children...)
}

func OptGroup(children ...Node) Node {
	return Element("optgroup", children...)
}

func Option(children ...Node) Node {
	return Element("option", children...)
}

func P(children ...Node) Node {
	return Element("p", children...)
}

func Param(children ...Node) Node {
	return Element("param", children...)
}

func Picture(children ...Node) Node {
	return Element("picture", children...)
}

func Pre(children ...Node) Node {
	return Element("pre", children...)
}

func Progress(children ...Node) Node {
	return Element("progress", children...)
}

func Script(children ...Node) Node {
	return Element("script", children...)
}

func Section(children ...Node) Node {
	return Element("section", children...)
}

func Select(children ...Node) Node {
	return Element("select", children...)
}

func Source(children ...Node) Node {
	return Element("source", children...)
}

func Span(children ...Node) Node {
	return Element("span", children...)
}

func StyleElement(children ...Node) Node {
	return Element("style", children...)
}

func Summary(children ...Node) Node {
	return Element("summary", children...)
}

func SVG(children ...Node) Node {
	return Element("svg", children...)
}

func Table(children ...Node) Node {
	return Element("table", children...)
}

func TBody(children ...Node) Node {
	return Element("tbody", children...)
}

func Td(children ...Node) Node {
	return Element("td", children...)
}

func Textarea(children ...Node) Node {
	return Element("textarea", children...)
}

func TFoot(children ...Node) Node {
	return Element("tfoot", children...)
}

func Th(children ...Node) Node {
	return Element("th", children...)
}

func THead(children ...Node) Node {
	return Element("thead", children...)
}

func Tr(children ...Node) Node {
	return Element("tr", children...)
}

func Ul(children ...Node) Node {
	return Element("ul", children...)
}

func Wbr(children ...Node) Node {
	return Element("wbr", children...)
}

func Abbr(children ...Node) Node {
	return Element("abbr", Group(children...))
}

func B(children ...Node) Node {
	return Element("b", Group(children...))
}

func Caption(children ...Node) Node {
	return Element("caption", Group(children...))
}

func Dd(children ...Node) Node {
	return Element("dd", Group(children...))
}

func DElement(children ...Node) Node {
	return Element("del", Group(children...))
}

func Dfn(children ...Node) Node {
	return Element("dfn", Group(children...))
}

func Dt(children ...Node) Node {
	return Element("dt", Group(children...))
}

func Em(children ...Node) Node {
	return Element("em", Group(children...))
}

func FigCaption(children ...Node) Node {
	return Element("figcaption", Group(children...))
}

func H1(children ...Node) Node {
	return Element("h1", Group(children...))
}

func H2(children ...Node) Node {
	return Element("h2", Group(children...))
}

func H3(children ...Node) Node {
	return Element("h3", Group(children...))
}

func H4(children ...Node) Node {
	return Element("h4", Group(children...))
}

func H5(children ...Node) Node {
	return Element("h5", Group(children...))
}

func H6(children ...Node) Node {
	return Element("h6", Group(children...))
}

func I(children ...Node) Node {
	return Element("i", Group(children...))
}

func Ins(children ...Node) Node {
	return Element("ins", Group(children...))
}

func Kbd(children ...Node) Node {
	return Element("kbd", Group(children...))
}

func Mark(children ...Node) Node {
	return Element("mark", Group(children...))
}

func Q(children ...Node) Node {
	return Element("q", Group(children...))
}

func S(children ...Node) Node {
	return Element("s", Group(children...))
}

func Samp(children ...Node) Node {
	return Element("samp", Group(children...))
}

func Small(children ...Node) Node {
	return Element("small", Group(children...))
}

func Strong(children ...Node) Node {
	return Element("strong", Group(children...))
}

func Sub(children ...Node) Node {
	return Element("sub", Group(children...))
}

func Sup(children ...Node) Node {
	return Element("sup", Group(children...))
}

func Time(children ...Node) Node {
	return Element("time", Group(children...))
}

func TitleElement(children ...Node) Node {
	return Element("title", Group(children...))
}

func U(children ...Node) Node {
	return Element("u", Group(children...))
}

func Var(children ...Node) Node {
	return Element("var", Group(children...))
}

func Video(children ...Node) Node {
	return Element("video", Group(children...))
}
