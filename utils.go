package htmx

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
