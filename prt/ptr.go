package prt

// ToPtr 泛型转指针类型
func ToPtr[T any](t T) *T {
	return &t
}

// ToPrtSlice 泛型切片转泛型指针切片类型
func ToPrtSlice[T any](src []T) []*T {
	dst := make([]*T, len(src))
	for srcIndex, srcValue := range src {
		dst[srcIndex] = ToPtr[T](srcValue)
	}
	return dst
}
