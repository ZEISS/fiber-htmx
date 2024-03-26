package htmx

import (
	"strconv"
	"strings"
)

// AsBool is a helper function that returns a boolean value based on the provided string.
func AsBool(str string) bool {
	return strings.EqualFold(str, "true")
}

// AsStr is a helper function that returns a string value based on the provided boolean.
func AsStr(v bool) string {
	return strconv.FormatBool(v)
}

// IntAsString is a helper function that returns a string value based on the provided integer.
func IntAsString(v int) string {
	return strconv.Itoa(v)
}

// Merge returns a new ClassNames object that is the result of merging the provided ClassNames objects.
func Merge(classNames ...ClassNames) ClassNames {
	merged := ClassNames{}

	for _, c := range classNames {
		for k, v := range c {
			merged[k] = v
		}
	}

	return merged
}
