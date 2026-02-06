package list

// FilterEmptySlices removes empty slices from a slice of slices
func FilterEmptySlices[T any, U ~[]T, S ~[]U](sliceOfSlices S) S {
	for i, slice := range sliceOfSlices {
		if len(slice) == 0 {
			n := make([]U, i, len(sliceOfSlices)-1)
			copy(n, []U(sliceOfSlices[:i]))
			for _, slice := range sliceOfSlices[i+1:] {
				if len(slice) != 0 {
					n = append(n, slice)
				}
			}
			return S(n)
		}
	}
	return sliceOfSlices
}
