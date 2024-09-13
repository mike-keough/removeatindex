package removeatindex

type Slice []interface{}

func RemoveAtIndex(s Slice, index int) Slice {
	if index < 0 || index >= len(s) {
		return s // Index out of bounds, return original slice
	}

	// Remove the element at the specified index
	return append(s[:index], s[index+1:]...)
}
