// Package set implements the set data structure and some of the fundamental
// operations on sets. A set is defined as a collection of unique elements of
// the same type. The type must implement comparable.
package sets

import (
	"fmt"
	"strings"

	"github.com/devkvlt/go/slices"
)

// Set represents a set of any comparable type T.
type Set[T comparable] map[T]struct{}

// New creates a new set containing the given elements.
// New can also be used to convert a slice to a set:
//
//	s := []int{1,2,3}
//	set := sets.New(s...)
func New[T comparable](elems ...T) Set[T] {
	set := Set[T]{}
	for _, e := range elems {
		set.Insert(e)
	}
	return set
}

// Equal reports whether two sets are equal. Two sets are equal if they have the
// same cardinal (number of elements) and one is a subset of the other. In
// particular, the nil set and the empty set are equal.
func Equal[T comparable](s1, s2 Set[T]) bool {
	if len(s1) != len(s2) {
		return false
	}
	return s1.Contains(s2)
}

// Disjoint reports whether two sets are disjoint, i.e. they don't have any
// element in common.
func Disjoint[T comparable](s1, s2 Set[T]) bool {
	// We want to range over the smaller set.
	var smaller, larger Set[T]
	if len(s1) > len(s2) {
		smaller, larger = s2, s1
	} else {
		smaller, larger = s1, s2
	}
	for e := range smaller {
		if larger.Has(e) {
			return false
		}
	}
	return true
}

// Contains reports whether the set s contains the set t.
func (s Set[T]) Contains(t Set[T]) bool {
	return s.Forall(func(e T) bool {
		return t.Has(e)
	})
}

// String implements the fmt.Stringer interface.
// TODO: use strings.Builder
func (s Set[T]) String() string {
	ss := s.Slice()
	toStr := func(e T) string { return fmt.Sprintf("%v", e) }
	strs := slices.Map(ss, toStr)
	return "Set{" + strings.Join(strs, " ") + "}"
}

// Insert inserts the given elements into the set s.
func (s Set[T]) Insert(e T) {
	s[e] = struct{}{}
}

// Delete removes the given elements from the set s.
func (s Set[T]) Delete(elems ...T) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Pop removes a random element from a non empty set and returns it. If the set
// is empty the zero value of the set type is returned.
func (s Set[T]) Pop() T {
	var x T
	for e := range s {
		x = e
		s.Delete(e)
		break
	}
	return x
}

// Has reports whether the set s has the given elements. If no arguments are
// given, true is returned.
func (s Set[T]) Has(e T) bool {
	_, ok := s[e]
	return ok
}

// Slice returns a slice containing the elements of a set.
func (s Set[T]) Slice() []T {
	if s == nil {
		return []T(nil)
	}
	ss := make([]T, len(s))
	i := 0
	for e := range s {
		ss[i] = e
		i++
	}
	return ss
}

// Forall reports whether every element of the set s satisfies the predicate
// function f.
func (s Set[T]) Forall(f func(T) bool) bool {
	for e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

// Exists reports whether there is at least one element in the set s that
// satisfies the predicate function f.
func (s Set[T]) Exists(f func(T) bool) bool {
	for e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

// Filter returns a subset of s containing only the elements of s that satisfy
// the predicate f.
func Filter[T comparable](s Set[T], f func(T) bool) Set[T] {
	filtered := New[T]()
	for e := range s {
		if f(e) {
			filtered.Insert(e)
		}
	}
	return filtered
}

// Map returns a set containing the elements obtained by applying f to the
// elements of s.
func Map[A, B comparable](s Set[A], f func(A) B) Set[B] {
	image := make(Set[B], len(s))
	for e := range s {
		image.Insert(f(e))
	}
	return image
}

// Clone returns a copy of s.
func Clone[T comparable](s Set[T]) Set[T] {
	ss := make(Set[T], len(s))
	for e := range s {
		ss.Insert(e)
	}
	return ss
}

// Union returns a set containing all the elements in the two sets.
func Union[T comparable](s1, s2 Set[T]) Set[T] {
	union := Clone(s1)
	for e := range s2 {
		union.Insert(e)
	}
	return union
}

// DisjointUnion
func DisjointUnion[T comparable](s1, s2 Set[T]) Set[T] {
	dis := make(Set[T])
	for e := range s1 {
		if !s2.Has(e) {
			dis.Insert(e)
		}
	}
	return dis
}

// Intersection
func Intersection[T comparable](s1, s2 Set[T]) Set[T] {
	inter := make(Set[T])
	for e := range s1 {
		if s2.Has(e) {
			inter.Insert(e)
		}
	}
	return inter
}

// Difference returns a set containing the elements of s1 that don't exist in s2.
// Example:
//
//	s1 := sets.New(1, 2, 3, 4)
//	s2 := sets.New(3, 4, 5)
//	fmt.Println(sets.Difference(s1, s2)) // out: Set{1 2}
func Difference[T comparable](s1, s2 Set[T]) Set[T] {
	diff := make(Set[T])
	for e := range s1 {
		if !s2.Has(e) {
			diff.Insert(e)
		}
	}
	return diff
}
