package htmx

import (
	"strconv"
	"strings"
)

// AsBool ...
func AsBool(str string) bool {
	return strings.EqualFold(str, "true")
}

// AsStr ...
func AsStr(v bool) string {
	return strconv.FormatBool(v)
}

// IntAsString ...
func IntAsString(v int) string {
	return strconv.Itoa(v)
}
