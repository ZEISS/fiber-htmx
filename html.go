package htmx

import (
	"io"
)

// Doctype represents the HTML doctype declaration.
func Doctype(sibling Node) Node {
	return NodeFunc(func(w io.Writer) error {
		if _, err := w.Write([]byte("<!DOCTYPE html>")); err != nil {
			return err
		}

		return sibling.Render(w)
	})
}

// A represents an HTML anchor element.
func A(children ...Node) Node {
	return Element("a", children...)
}

// Address represents an HTML address element.
func Address(children ...Node) Node {
	return Element("address", children...)
}

// Area represents an HTML area element.
func Area(children ...Node) Node {
	return Element("area", children...)
}

// Article represents an HTML article element.
func Article(children ...Node) Node {
	return Element("article", children...)
}

// Aside represents an HTML aside element.
func Aside(children ...Node) Node {
	return Element("aside", children...)
}

// Audio represents an HTML audio element.
func Audio(children ...Node) Node {
	return Element("audio", children...)
}

// Base represents an HTML base element.
func Base(children ...Node) Node {
	return Element("base", children...)
}

// BlockQuote represents an HTML blockquote element.
func BlockQuote(children ...Node) Node {
	return Element("blockquote", children...)
}

// Body represents an HTML body element.
func Body(children ...Node) Node {
	return Element("body", children...)
}

// Br represents an HTML line break element.
func Br(children ...Node) Node {
	return Element("br", children...)
}

// Button represents an HTML button element.
func Button(children ...Node) Node {
	return Element("button", children...)
}

// Canvas represents an HTML canvas element.
func Canvas(children ...Node) Node {
	return Element("canvas", children...)
}

// Cite represents an HTML cite element.
func Cite(children ...Node) Node {
	return Element("cite", children...)
}

// Code represents an HTML code element.
func Code(children ...Node) Node {
	return Element("code", children...)
}

// Col represents an HTML col element.
func Col(children ...Node) Node {
	return Element("col", children...)
}

// ColGroup represents an HTML colgroup element.
func ColGroup(children ...Node) Node {
	return Element("colgroup", children...)
}

// DataElement represents an HTML data element.
func DataElement(children ...Node) Node {
	return Element("data", children...)
}

// DataList represents an HTML datalist element.
func DataList(children ...Node) Node {
	return Element("datalist", children...)
}

// Details represents an HTML details element.
func Details(children ...Node) Node {
	return Element("details", children...)
}

// Dialog represents an HTML dialog element.
func Dialog(children ...Node) Node {
	return Element("dialog", children...)
}

// Div represents an HTML div element.
func Div(children ...Node) Node {
	return Element("div", children...)
}

// Dl represents an HTML dl element.
func Dl(children ...Node) Node {
	return Element("dl", children...)
}

// Embed represents an HTML embed element.
func Embed(children ...Node) Node {
	return Element("embed", children...)
}

// FormElement represents an HTML form element.
func FormElement(children ...Node) Node {
	return Element("form", children...)
}

// Form represents an HTML form element.
func Form(children ...Node) Node {
	return Element("form", children...)
}

// FieldSet represents an HTML fieldset element.
func FieldSet(children ...Node) Node {
	return Element("fieldset", children...)
}

// Figure represents an HTML figure element.
func Figure(children ...Node) Node {
	return Element("figure", children...)
}

// Footer represents an HTML footer element.
func Footer(children ...Node) Node {
	return Element("footer", children...)
}

// Head represents an HTML head element.
func Head(children ...Node) Node {
	return Element("head", children...)
}

// Header represents an HTML header element.
func Header(children ...Node) Node {
	return Element("header", children...)
}

// HGroup represents an HTML hgroup element.
func HGroup(children ...Node) Node {
	return Element("hgroup", children...)
}

// Hr represents an HTML horizontal rule element.
func Hr(children ...Node) Node {
	return Element("hr", children...)
}

// HTML represents an HTML html element.
func HTML(children ...Node) Node {
	return Element("html", children...)
}

// IFrame represents an HTML iframe element.
func IFrame(children ...Node) Node {
	return Element("iframe", children...)
}

// Img represents an HTML img element.
func Img(children ...Node) Node {
	return Element("img", children...)
}

// Input represents an HTML input element.
func Input(children ...Node) Node {
	return Element("input", children...)
}

// Label represents an HTML label element.
func Label(children ...Node) Node {
	return Element("label", children...)
}

// Legend represents an HTML legend element.
func Legend(children ...Node) Node {
	return Element("legend", children...)
}

// Li represents an HTML li element.
func Li(children ...Node) Node {
	return Element("li", children...)
}

// Link represents an HTML link element.
func Link(children ...Node) Node {
	return Element("link", children...)
}

// Main represents an HTML main element.
func Main(children ...Node) Node {
	return Element("main", children...)
}

// Meta represents an HTML meta element.
func Meta(children ...Node) Node {
	return Element("meta", children...)
}

// Meter represents an HTML meter element.
func Meter(children ...Node) Node {
	return Element("meter", children...)
}

// Nav represents an HTML nav element.
func Nav(children ...Node) Node {
	return Element("nav", children...)
}

// NoScript represents an HTML noscript element.
func NoScript(children ...Node) Node {
	return Element("noscript", children...)
}

// Object represents an HTML object element.
func Object(children ...Node) Node {
	return Element("object", children...)
}

// Ol represents an HTML ol element.
func Ol(children ...Node) Node {
	return Element("ol", children...)
}

// OptGroup represents an HTML optgroup element.
func OptGroup(children ...Node) Node {
	return Element("optgroup", children...)
}

// Option represents an HTML option element.
func Option(children ...Node) Node {
	return Element("option", children...)
}

// P represents an HTML p element.
func P(children ...Node) Node {
	return Element("p", children...)
}

// Param represents an HTML param element.
func Param(children ...Node) Node {
	return Element("param", children...)
}

// Picture represents an HTML picture element.
func Picture(children ...Node) Node {
	return Element("picture", children...)
}

// Pre represents an HTML pre element.
func Pre(children ...Node) Node {
	return Element("pre", children...)
}

// Progress represents an HTML progress element.
func Progress(children ...Node) Node {
	return Element("progress", children...)
}

// Script represents an HTML script element.
func Script(children ...Node) Node {
	return Element("script", children...)
}

// Section represents an HTML section element.
func Section(children ...Node) Node {
	return Element("section", children...)
}

// Select represents an HTML select element.
func Select(children ...Node) Node {
	return Element("select", children...)
}

// Source represents an HTML source element.
func Source(children ...Node) Node {
	return Element("source", children...)
}

// Span represents an HTML span element.
func Span(children ...Node) Node {
	return Element("span", children...)
}

// StyleElement represents an HTML style element.
func StyleElement(children ...Node) Node {
	return Element("style", children...)
}

// Summary represents an HTML summary element.
func Summary(children ...Node) Node {
	return Element("summary", children...)
}

// Table represents an HTML table element.
func Table(children ...Node) Node {
	return Element("table", children...)
}

// TBody represents an HTML tbody element.
func TBody(children ...Node) Node {
	return Element("tbody", children...)
}

// Td represents an HTML td element.
func Td(children ...Node) Node {
	return Element("td", children...)
}

// Textarea represents an HTML textarea element.
func Textarea(children ...Node) Node {
	return Element("textarea", children...)
}

// TFoot represents an HTML tfoot element.
func TFoot(children ...Node) Node {
	return Element("tfoot", children...)
}

// Th represents an HTML th element.
func Th(children ...Node) Node {
	return Element("th", children...)
}

// THead represents an HTML thead element.
func THead(children ...Node) Node {
	return Element("thead", children...)
}

// Tr represents an HTML tr element.
func Tr(children ...Node) Node {
	return Element("tr", children...)
}

// Ul represents an HTML ul element.
func Ul(children ...Node) Node {
	return Element("ul", children...)
}

// Wbr represents an HTML wbr element.
func Wbr(children ...Node) Node {
	return Element("wbr", children...)
}

// Abbr represents an HTML abbr element.
func Abbr(children ...Node) Node {
	return Element("abbr", Group(children...))
}

// B represents an HTML b element.
func B(children ...Node) Node {
	return Element("b", Group(children...))
}

// Caption represents an HTML caption element.
func Caption(children ...Node) Node {
	return Element("caption", Group(children...))
}

// Dd represents an HTML dd element.
func Dd(children ...Node) Node {
	return Element("dd", Group(children...))
}

// DElement represents an HTML del element.
func DElement(children ...Node) Node {
	return Element("del", Group(children...))
}

// Dfn represents an HTML dfn element.
func Dfn(children ...Node) Node {
	return Element("dfn", Group(children...))
}

// Dt represents an HTML dt element.
func Dt(children ...Node) Node {
	return Element("dt", Group(children...))
}

// Em represents an HTML em element.
func Em(children ...Node) Node {
	return Element("em", Group(children...))
}

// FigCaption represents an HTML figcaption element.
func FigCaption(children ...Node) Node {
	return Element("figcaption", Group(children...))
}

// H1 represents an HTML h1 element.
func H1(children ...Node) Node {
	return Element("h1", Group(children...))
}

// H2 represents an HTML h2 element.
func H2(children ...Node) Node {
	return Element("h2", Group(children...))
}

// H3 represents an HTML h3 element.
func H3(children ...Node) Node {
	return Element("h3", Group(children...))
}

// H4 represents an HTML h4 element.
func H4(children ...Node) Node {
	return Element("h4", Group(children...))
}

// H5 represents an HTML h5 element.
func H5(children ...Node) Node {
	return Element("h5", Group(children...))
}

// H6 represents an HTML h6 element.
func H6(children ...Node) Node {
	return Element("h6", Group(children...))
}

// I represents an HTML i element.
func I(children ...Node) Node {
	return Element("i", Group(children...))
}

// Ins represents an HTML ins element.
func Ins(children ...Node) Node {
	return Element("ins", Group(children...))
}

// Kbd represents an HTML kbd element.
func Kbd(children ...Node) Node {
	return Element("kbd", Group(children...))
}

// Mark represents an HTML mark element.
func Mark(children ...Node) Node {
	return Element("mark", Group(children...))
}

// Q represents an HTML q element.
func Q(children ...Node) Node {
	return Element("q", Group(children...))
}

// S represents an HTML s element.
func S(children ...Node) Node {
	return Element("s", Group(children...))
}

// Samp represents an HTML samp element.
func Samp(children ...Node) Node {
	return Element("samp", Group(children...))
}

// Small represents an HTML small element.
func Small(children ...Node) Node {
	return Element("small", Group(children...))
}

// Strong represents an HTML strong element.
func Strong(children ...Node) Node {
	return Element("strong", Group(children...))
}

// Sub represents an HTML sub element.
func Sub(children ...Node) Node {
	return Element("sub", Group(children...))
}

// Sup represents an HTML sup element.
func Sup(children ...Node) Node {
	return Element("sup", Group(children...))
}

// Time represents an HTML time element.
func Time(children ...Node) Node {
	return Element("time", Group(children...))
}

// TitleElement represents an HTML title element.
func TitleElement(children ...Node) Node {
	return Element("title", Group(children...))
}

// Title represents an HTML title element.
func Title(children ...Node) Node {
	return Element("title", Group(children...))
}

// U represents an HTML u element.
func U(children ...Node) Node {
	return Element("u", Group(children...))
}

// Var represents an HTML var element.
func Var(children ...Node) Node {
	return Element("var", Group(children...))
}

// Video represents an HTML video element.
func Video(children ...Node) Node {
	return Element("video", Group(children...))
}

// Template represents an HTML template element.
func Template(children ...Node) Node {
	return Element("template", Group(children...))
}

// Slot represents an HTML slot element.
func Slot(children ...Node) Node {
	return Element("slot", Group(children...))
}

// Track represents an HTML track element.
func Track(children ...Node) Node {
	return Element("track", Group(children...))
}
