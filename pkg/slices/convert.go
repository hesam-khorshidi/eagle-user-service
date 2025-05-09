package slices

import "fmt"

func Convert[T, T2 any](inputSlice []T, caster func(T) T2) []T2 {
	s := make([]T2, len(inputSlice))
	for i, element := range inputSlice {
		s[i] = caster(element)
	}

	return s
}

func ToString[T fmt.Stringer](inputSlice []T) []string {
	return Convert(inputSlice, func(t T) string {
		return t.String()
	})
}

func ToAny[T fmt.Stringer](inputSlice []T) []any {
	return Convert(inputSlice, func(t T) any {
		return t
	})
}
