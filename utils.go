package htmx

import "strings"

// AsBool ...
func AsBool(str string) bool {
	return strings.EqualFold(str, "true")
}
