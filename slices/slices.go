package slices

// Equal reports if two slices are equal.
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// Unique filters out repeated elements in a slice. It returns a new slice.
func Unique[T comparable](s []T) []T {
	occurred := map[T]struct{}{}
	unique := []T{}
	for _, v := range s {
		if _, ok := occurred[v]; !ok {
			occurred[v] = struct{}{}
			unique = append(unique, v)
		}
	}
	return unique
}

// Filter filters out the elements of the slice that don't satisfy the predicate
// f. It returns a new slice.
func Filter[T any](s []T, f func(T) bool) []T {
	filtered := []T{}
	for _, v := range s {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Map applies the function f to the elements of a slice and returns a slice
// with the resulting images.
func Map[A, B any](s []A, f func(A) B) []B {
	image := make([]B, len(s))
	for i, v := range s {
		image[i] = f(v)
	}
	return image
}
