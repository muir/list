package list

func Copy[T any](orig []T) []T {
	if orig == nil {
		return nil
	}
	c := make([]T, len(orig))
	copy(c, orig)
	return c
}
