package util

// InSliceString Determine whether the string is in the slice.
func InSliceString(needle string, slices []string) bool {
	for _, value := range slices {
		if value == needle {
			return true
		}
	}
	return false
}
