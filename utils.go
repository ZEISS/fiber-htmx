package htmx

import "encoding/json"

// Merge returns a new ClassNames object that is the result of merging the provided ClassNames objects.
func Merge(classNames ...ClassNames) ClassNames {
	merged := ClassNames{}

	for i := 0; i < len(classNames); i++ {
		for k, v := range classNames[i] {
			merged[k] = v
		}
	}

	return merged
}

// JsonSerializeOrEmpty returns a JSON serialized string of the provided data or an empty string if the serialization fails.
func JsonSerializeOrEmpty(data any) string {
	if data == nil {
		return ""
	}

	serialized, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(serialized)
}
