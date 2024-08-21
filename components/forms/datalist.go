package forms

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/pkg/utilx"
)

// DatalistProps is the type of the props for the Datalist component
type DatalistProps struct {
	// ID is the id of the datalist
	ID string
	// URL is the url of the datalist
	URL string
	// ClassNames is a map of class names to conditionally add to the component
	ClassNames htmx.ClassNames
	// Name is the name of the datalist
	Name string
	// Placeholder is the placeholder of the datalist
	Placeholder string
	// Disabled is whether the datalist is disabled
	Disabled bool
	// Indicator is the indicator of the datalist
	Indicator string
}

// Datalist is a component that displays a datalist.
func Datalist(props DatalistProps, children ...htmx.Node) htmx.Node {
	return htmx.Fragment(
		htmx.Script(htmx.Raw(`function checkUserKeydown(event) {return event instanceof KeyboardEvent}`)),
		TextInputBordered(
			TextInputProps{
				ClassNames:  props.ClassNames,
				Name:        props.Name,
				Placeholder: props.Placeholder,
				Disabled:    props.Disabled,
			},
			htmx.List(props.Name),
			htmx.HxGet(props.URL),
			htmx.HxTarget(fmt.Sprintf("#%s", props.ID)),
			htmx.HxTrigger("load, keyup[checkUserKeydown.call(this, event)] changed delay:350ms"),
			htmx.HxIndicator(utilx.IfElse(utilx.Empty(props.Indicator), htmx.HxClssNameIndicatorSelector, props.Indicator)),
		),
		htmx.DataList(
			htmx.ID(props.ID),
		),
		htmx.Group(children...),
	)
}
