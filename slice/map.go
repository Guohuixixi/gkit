package slice

// Map 将一个切片映射为另一个切片
func Map[Src any, Dst any](src []Src, fn func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = fn(i, s)
	}
	return dst
}
