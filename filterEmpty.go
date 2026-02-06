package list

// FilterEmptySlices removes empty slices from a slice of slices
func FilterEmptySlices[T any](sliceOfSlices [][]T) [][]T {
	for i, slice := range sliceOfSlices {
		if len(slice) == 0 {
			n := make([][]T, i, len(sliceOfSlices)-1)
			copy(n, sliceOfSlices[:i])
			for _, slice := range sliceOfSlices[i+1:] {
				if len(slice) != 0 {
					n = append(n, slice)
				}
			}
			return n
		}
	}
	return sliceOfSlices
}
