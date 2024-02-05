package htmx

import (
	"fmt"
)

// HxEventType ...
type HxEventType string

// String ...
func (e HxEventType) String() string {
	return string(e)
}

const (
	HxEventTypeAbort            HxEventType = "htmx:abort"
	HxEventTypeAfterLoad        HxEventType = "htmx:afterLoad"
	HxEventTypeAfterProcessNode HxEventType = "htmx:afterProcessNode"
	HxEventTypeAfterRequest     HxEventType = "htmx:afterRequest"
)

func HxBoost(v bool) Node {
	return Attribute("hx-boost", AsStr(v))
}

func HxGet(url string) Node {
	return Attribute("hx-get", url)
}

func HxPost(url string) Node {
	return Attribute("hx-post", url)
}

func HxPushUrl(v bool) Node {
	return Attribute("hx-boost", AsStr(v))
}

func HxTarget(target string) Node {
	return Attribute("hx-target", target)
}

func HxSelect(target string) Node {
	return Attribute("hx-select", target)
}

func HxSelectOob(target string) Node {
	return Attribute("hx-select-oob", target)
}

func HxSwap(target string) Node {
	return Attribute("hx-swap", target)
}

func HxSwapOob(target string) Node {
	return Attribute("hx-swap-oob", target)
}

func HxTrigger(target string) Node {
	return Attribute("hx-trigger", target)
}

func HxConfirm(msg string) Node {
	return Attribute("hx-confirm", msg)
}

func HxPrompt(msg string) Node {
	return Attribute("hx-prompt", msg)
}

func HxDelete(url string) Node {
	return Attribute("hx-delete", url)
}

func HxOn(target string, js string) Node {
	return Attribute(fmt.Sprintf("hx-put-%s", target), js)
}

func HxPut(url string) Node {
	return Attribute("hx-put", url)
}

func HxPath(url string) Node {
	return Attribute("hx-patch", url)
}

func HxIndicator(target string) Node {
	return Attribute("hx-indicator", target)
}

func HxEncoding(enc string) Node {
	return Attribute("hx-encoding", enc)
}

func HxExt(ext string) Node {
	return Attribute("hx-ext", ext)
}

func HxDisable() Node {
	return Attribute("hx-disable")
}

func HxValidate(v bool) Node {
	return Attribute("hx-validate", AsStr(v))
}

// func HxVals(vals string) Node {
// 	return Attribute("hx-vals", vals)
// }

func Async() Node {
	return Attribute("async")
}

func AutoFocus() Node {
	return Attribute("autofocus")
}

func AutoPlay() Node {
	return Attribute("autoplay")
}

func Checked() Node {
	return Attribute("checked")
}

func Controls() Node {
	return Attribute("controls")
}

func Defer() Node {
	return Attribute("defer")
}

func Disabled() Node {
	return Attribute("disabled")
}

func Loop() Node {
	return Attribute("loop")
}

func Multiple() Node {
	return Attribute("multiple")
}

func Muted() Node {
	return Attribute("muted")
}

func PlaysInline() Node {
	return Attribute("playsinline")
}

func ReadOnly() Node {
	return Attribute("readonly")
}

func Required() Node {
	return Attribute("required")
}

func Selected() Node {
	return Attribute("selected")
}

func Accept(v string) Node {
	return Attribute("accept", v)
}

func Action(v string) Node {
	return Attribute("action", v)
}

func Alt(v string) Node {
	return Attribute("alt", v)
}

func Aria(name, v string) Node {
	return Attribute("aria-"+name, v)
}

func As(v string) Node {
	return Attribute("as", v)
}

func AutoComplete(v string) Node {
	return Attribute("autocomplete", v)
}

func Charset(v string) Node {
	return Attribute("charset", v)
}

func Class(v string) Node {
	return Attribute("class", v)
}

func Cols(v string) Node {
	return Attribute("cols", v)
}

func ColSpan(v string) Node {
	return Attribute("colspan", v)
}

func Content(v string) Node {
	return Attribute("content", v)
}

func DataAttribute(name, v string) Node {
	return Attribute("data-"+name, v)
}

func For(v string) Node {
	return Attribute("for", v)
}

func FormAttribute(v string) Node {
	return Attribute("form", v)
}

func Height(v string) Node {
	return Attribute("height", v)
}

func Href(v string) Node {
	return Attribute("href", v)
}

func ID(v string) Node {
	return Attribute("id", v)
}

func Lang(v string) Node {
	return Attribute("lang", v)
}

func Loading(v string) Node {
	return Attribute("loading", v)
}

func Max(v string) Node {
	return Attribute("max", v)
}

func MaxLength(v string) Node {
	return Attribute("maxlength", v)
}

func Method(v string) Node {
	return Attribute("method", v)
}

func Min(v string) Node {
	return Attribute("min", v)
}

func MinLength(v string) Node {
	return Attribute("minlength", v)
}

func Name(v string) Node {
	return Attribute("name", v)
}

func Pattern(v string) Node {
	return Attribute("pattern", v)
}

func Placeholder(v string) Node {
	return Attribute("placeholder", v)
}

func Poster(v string) Node {
	return Attribute("poster", v)
}

func Preload(v string) Node {
	return Attribute("preload", v)
}

func Rel(v string) Node {
	return Attribute("rel", v)
}

func Role(v string) Node {
	return Attribute("role", v)
}

func Rows(v string) Node {
	return Attribute("rows", v)
}

func RowSpan(v string) Node {
	return Attribute("rowspan", v)
}

func Src(v string) Node {
	return Attribute("src", v)
}

func SrcSet(v string) Node {
	return Attribute("srcset", v)
}

func Step(v string) Node {
	return Attribute("step", v)
}

func StyleAttribute(v string) Node {
	return Attribute("style", v)
}

func TabIndex(v string) Node {
	return Attribute("tabindex", v)
}

func Target(v string) Node {
	return Attribute("target", v)
}

func TitleAttribute(v string) Node {
	return Attribute("title", v)
}

func Type(v string) Node {
	return Attribute("type", v)
}

func Value(v string) Node {
	return Attribute("value", v)
}

func Width(v string) Node {
	return Attribute("width", v)
}

func EncType(v string) Node {
	return Attribute("enctype", v)
}
