// Package htmx provides a set of functions for generating HTML attributes as nodes.
// These attributes can be used in conjunction with the Fiber framework and the htmx library.

package htmx

import (
	"fmt"
)

// HxEventType represents the type of htmx event.
type HxEventType string

// String returns the string representation of the HxEventType.
func (e HxEventType) String() string {
	return string(e)
}

// List of predefined htmx event types.
const (
	HxEventTypeAbort            HxEventType = "htmx:abort"
	HxEventTypeAfterLoad        HxEventType = "htmx:afterLoad"
	HxEventTypeAfterProcessNode HxEventType = "htmx:afterProcessNode"
	HxEventTypeAfterRequest     HxEventType = "htmx:afterRequest"
)

// HxBoost sets the hx-boost attribute to enable or disable boosting.
func HxBoost(v bool) Node {
	return Attribute("hx-boost", AsStr(v))
}

// HxGet sets the hx-get attribute to specify the URL for GET requests.
func HxGet(url string) Node {
	return Attribute("hx-get", url)
}

// HxPost sets the hx-post attribute to specify the URL for POST requests.
func HxPost(url string) Node {
	return Attribute("hx-post", url)
}

// HxPushUrl sets the hx-push-url attribute to enable or disable URL pushing.
func HxPushUrl(v bool) Node {
	return Attribute("hx-boost", AsStr(v))
}

// HxTarget sets the hx-target attribute to specify the target element for the response.
func HxTarget(target string) Node {
	return Attribute("hx-target", target)
}

// HxSelect sets the hx-select attribute to specify the target element for selection.
func HxSelect(target string) Node {
	return Attribute("hx-select", target)
}

// HxSelectOob sets the hx-select-oob attribute to specify the target element for out-of-band selection.
func HxSelectOob(target string) Node {
	return Attribute("hx-select-oob", target)
}

// HxSwap sets the hx-swap attribute to specify the target element for swapping.
func HxSwap(target string) Node {
	return Attribute("hx-swap", target)
}

// HxSwapOob sets the hx-swap-oob attribute to specify the target element for out-of-band swapping.
func HxSwapOob(target string) Node {
	return Attribute("hx-swap-oob", target)
}

// HxTrigger sets the hx-trigger attribute to specify the target element for triggering an event.
func HxTrigger(target string) Node {
	return Attribute("hx-trigger", target)
}

// HxConfirm sets the hx-confirm attribute to display a confirmation message.
func HxConfirm(msg string) Node {
	return Attribute("hx-confirm", msg)
}

// HxPrompt sets the hx-prompt attribute to display a prompt message.
func HxPrompt(msg string) Node {
	return Attribute("hx-prompt", msg)
}

// HxDelete sets the hx-delete attribute to specify the URL for DELETE requests.
func HxDelete(url string) Node {
	return Attribute("hx-delete", url)
}

// HxOn sets the hx-put-{target} attribute to specify the JavaScript code to execute on a PUT request.
func HxOn(target string, js string) Node {
	return Attribute(fmt.Sprintf("hx-on:%s", target), js)
}

// HxPut sets the hx-put attribute to specify the URL for PUT requests.
func HxPut(url string) Node {
	return Attribute("hx-put", url)
}

// HxPath sets the hx-patch attribute to specify the URL for PATCH requests.
func HxPath(url string) Node {
	return Attribute("hx-patch", url)
}

// HxIndicator sets the hx-indicator attribute to specify the target element for showing an indicator.
func HxIndicator(target string) Node {
	return Attribute("hx-indicator", target)
}

// HxEncoding sets the hx-encoding attribute to specify the encoding type for form submission.
func HxEncoding(enc string) Node {
	return Attribute("hx-encoding", enc)
}

// HxExt sets the hx-ext attribute to specify the file extension for file uploads.
func HxExt(ext string) Node {
	return Attribute("hx-ext", ext)
}

// HxDisable sets the hx-disable attribute to disable htmx functionality.
func HxDisable() Node {
	return Attribute("hx-disable")
}

// HxDisabledElt sets the hx-disable-elt attribute to disable the target element.
func HxDisabledElt(target string) Node {
	return Attribute("hx-disabled-elt", target)
}

// HxValidate sets the hx-validate attribute to enable or disable form validation.
func HxValidate(v bool) Node {
	return Attribute("hx-validate", AsStr(v))
}

// HxInclude sets the hx-include attribute to specify the target element for inclusion.
func HxInclude(target string) Node {
	return Attribute("hx-include", target)
}

// Async sets the async attribute for script elements.
func Async() Node {
	return Attribute("async")
}

// AutoFocus sets the autofocus attribute for form elements.
func AutoFocus() Node {
	return Attribute("autofocus")
}

// AutoPlay sets the autoplay attribute for media elements.
func AutoPlay() Node {
	return Attribute("autoplay")
}

// Checked sets the checked attribute for input elements.
func Checked() Node {
	return Attribute("checked")
}

// Controls sets the controls attribute for media elements.
func Controls() Node {
	return Attribute("controls")
}

// Defer sets the defer attribute for script elements.
func Defer() Node {
	return Attribute("defer")
}

// Disabled sets the disabled attribute for form elements.
func Disabled() Node {
	return Attribute("disabled")
}

// Loop sets the loop attribute for media elements.
func Loop() Node {
	return Attribute("loop")
}

// Multiple sets the multiple attribute for input elements.
func Multiple() Node {
	return Attribute("multiple")
}

// Muted sets the muted attribute for media elements.
func Muted() Node {
	return Attribute("muted")
}

// PlaysInline sets the playsinline attribute for media elements.
func PlaysInline() Node {
	return Attribute("playsinline")
}

// ReadOnly sets the readonly attribute for form elements.
func ReadOnly() Node {
	return Attribute("readonly")
}

// Required sets the required attribute for form elements.
func Required() Node {
	return Attribute("required")
}

// Selected sets the selected attribute for option elements.
func Selected() Node {
	return Attribute("selected")
}

// Accept sets the accept attribute for file input elements.
func Accept(v string) Node {
	return Attribute("accept", v)
}

// Action sets the action attribute for form elements.
func Action(v string) Node {
	return Attribute("action", v)
}

// Alt sets the alt attribute for image elements.
func Alt(v string) Node {
	return Attribute("alt", v)
}

// Aria sets the aria-{name} attribute for elements.
func Aria(name, v string) Node {
	return Attribute("aria-"+name, v)
}

// As sets the as attribute for link elements.
func As(v string) Node {
	return Attribute("as", v)
}

// AutoComplete sets the autocomplete attribute for form elements.
func AutoComplete(v string) Node {
	return Attribute("autocomplete", v)
}

// Charset sets the charset attribute for meta elements.
func Charset(v string) Node {
	return Attribute("charset", v)
}

// Class sets the class attribute for elements.
func Class(v string) Node {
	return Attribute("class", v)
}

// Cols sets the cols attribute for textarea elements.
func Cols(v string) Node {
	return Attribute("cols", v)
}

// ColSpan sets the colspan attribute for table cells.
func ColSpan(v string) Node {
	return Attribute("colspan", v)
}

// Content sets the content attribute for meta elements.
func Content(v string) Node {
	return Attribute("content", v)
}

// DataAttribute sets the data-{name} attribute for elements.
func DataAttribute(name, v string) Node {
	return Attribute("data-"+name, v)
}

// For sets the for attribute for label elements.
func For(v string) Node {
	return Attribute("for", v)
}

// FormAttribute sets the form attribute for elements.
func FormAttribute(v string) Node {
	return Attribute("form", v)
}

// Height sets the height attribute for elements.
func Height(v string) Node {
	return Attribute("height", v)
}

// Href sets the href attribute for anchor elements.
func Href(v string) Node {
	return Attribute("href", v)
}

// ID sets the id attribute for elements.
func ID(v string) Node {
	return Attribute("id", v)
}

// Lang sets the lang attribute for elements.
func Lang(v string) Node {
	return Attribute("lang", v)
}

// Loading sets the loading attribute for elements.
func Loading(v string) Node {
	return Attribute("loading", v)
}

// Max sets the max attribute for input elements.
func Max(v string) Node {
	return Attribute("max", v)
}

// MaxLength sets the maxlength attribute for input elements.
func MaxLength(v string) Node {
	return Attribute("maxlength", v)
}

// Method sets the method attribute for form elements.
func Method(v string) Node {
	return Attribute("method", v)
}

// Min sets the min attribute for input elements.
func Min(v string) Node {
	return Attribute("min", v)
}

// MinLength sets the minlength attribute for input elements.
func MinLength(v string) Node {
	return Attribute("minlength", v)
}

// Name sets the name attribute for elements.
func Name(v string) Node {
	return Attribute("name", v)
}

// Pattern sets the pattern attribute for input elements.
func Pattern(v string) Node {
	return Attribute("pattern", v)
}

// Placeholder sets the placeholder attribute for input elements.
func Placeholder(v string) Node {
	return Attribute("placeholder", v)
}

// Poster sets the poster attribute for video elements.
func Poster(v string) Node {
	return Attribute("poster", v)
}

// Preload sets the preload attribute for media elements.
func Preload(v string) Node {
	return Attribute("preload", v)
}

// Rel sets the rel attribute for link elements.
func Rel(v string) Node {
	return Attribute("rel", v)
}

// Role sets the role attribute for elements.
func Role(v string) Node {
	return Attribute("role", v)
}

// Rows sets the rows attribute for textarea elements.
func Rows(v string) Node {
	return Attribute("rows", v)
}

// RowSpan sets the rowspan attribute for table cells.
func RowSpan(v string) Node {
	return Attribute("rowspan", v)
}

// Src sets the src attribute for elements.
func Src(v string) Node {
	return Attribute("src", v)
}

// SrcSet sets the srcset attribute for elements.
func SrcSet(v string) Node {
	return Attribute("srcset", v)
}

// Step sets the step attribute for input elements.
func Step(v string) Node {
	return Attribute("step", v)
}

// StyleAttribute sets the style attribute for elements.
func StyleAttribute(v string) Node {
	return Attribute("style", v)
}

// TabIndex sets the tabindex attribute for elements.
func TabIndex(v string) Node {
	return Attribute("tabindex", v)
}

// Target sets the target attribute for elements.
func Target(v string) Node {
	return Attribute("target", v)
}

// TitleAttribute sets the title attribute for elements.
func TitleAttribute(v string) Node {
	return Attribute("title", v)
}

// Type sets the type attribute for elements.
func Type(v string) Node {
	return Attribute("type", v)
}

// Value sets the value attribute for elements.
func Value(v string) Node {
	return Attribute("value", v)
}

// Width sets the width attribute for elements.
func Width(v string) Node {
	return Attribute("width", v)
}

// EncType sets the enctype attribute for form elements.
func EncType(v string) Node {
	return Attribute("enctype", v)
}

// OnClick sets the onclick attribute for elements.
func OnClick(v string) Node {
	return Attribute("onclick", v)
}
