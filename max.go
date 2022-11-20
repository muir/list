package list

type Number interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~float32 | ~float64 
}

func Max[N Number](x N, y ...N) N {
	for _, z := range y {
		if z > x {
			x = z
		}
	}
	return x
}

func Min[N Number](x N, y ...N) N {
	for _, z := range y {
		if z < x {
			x = z
		}
	}
	return x
}
