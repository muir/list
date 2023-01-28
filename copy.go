package list

// Copy makes a shallow copy of it's input.
// Given nil, it returns nil.
func Copy[T any](orig []T) []T {
	if orig == nil {
		return nil
	}
	c := make([]T, len(orig))
	copy(c, orig)
	return c
}
