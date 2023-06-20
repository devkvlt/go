package math

import (
	"golang.org/x/exp/constraints"
)

// Min returns the minimum of the given numbers. If no arguments are given it
// panics.
func Min[T constraints.Ordered](nums ...T) T {
	var min T
	n := len(nums)
	if n == 0 {
		panic("no arguments given")
	}
	min = nums[0]
	for i := 1; i < n; i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	return min
}

// Max returns the maximum of the given numbers. If no arguments are given, it
// panics.
func Max[T constraints.Ordered](nums ...T) T {
	var max T
	n := len(nums)
	if n == 0 {
		panic("no arguments given")
	}
	max = nums[0]
	for i := 1; i < n; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

// Sum returns the sum of the given numbers. If no arguments are given, it returns
// zero.
func Sum[T constraints.Integer | constraints.Float | constraints.Complex](nums ...T) T {
	var s T
	for _, v := range nums {
		s += v
	}
	return s
}
