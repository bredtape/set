package set

import (
	"fmt"
	"sort"
	"strings"
)

type Set[T comparable] map[T]struct{}

// instansiate
func New[T comparable](optionalCapacity ...int) Set[T] {
	if len(optionalCapacity) == 0 {
		return make(Set[T], 0)
	}

	return make(Set[T], optionalCapacity[0])
}

func NewValues[T comparable](xs ...T) Set[T] {
	set := New[T](len(xs))
	for _, x := range xs {
		set.Add(x)
	}
	return set
}

// alter

// add elements from list
func (set Set[T]) Add(xs ...T) {
	for _, x := range xs {
		set[x] = struct{}{}
	}
}

// add elements from other set
func (set Set[T]) AddSets(sets ...Set[T]) {
	for _, other := range sets {
		for x := range other {
			set[x] = struct{}{}
		}
	}
}

// remove elements from list
func (set Set[T]) Remove(xs ...T) {
	for _, x := range xs {
		delete(set, x)
	}
}

// remove elements from other sets
func (set Set[T]) RemoveSets(sets ...Set[T]) {
	for _, other := range sets {
		for x := range other {
			delete(set, x)
		}
	}
}

// create deep copy
func (set Set[T]) Copy() Set[T] {
	set2 := New[T](len(set))
	for x := range set {
		set2.Add(x)
	}
	return set2
}

// number of elements
func (set Set[T]) Count() int {
	return len(set)
}

// whether the set contains any elements
func (set Set[T]) Any() bool {
	return len(set) != 0
}

// whether the set contains all of the listed items.
// a set always contains the empty set
func (set Set[T]) Contains(xs ...T) bool {
	for _, x := range xs {
		if _, exists := set[x]; !exists {
			return false
		}
	}
	return true
}

func (set Set[T]) String() string {
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, fmt.Sprint(k))
	}
	sort.Strings(keys)
	return "[" + strings.Join(keys, ", ") + "]"
}

func (set Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}

func (set Set[T]) ToSortedSlice(less func(T, T) bool) []T {
	keys := set.ToSlice()

	sort.Slice(keys, func(i, j int) bool {
		return less(keys[i], keys[j])
	})
	return keys
}

func (set1 Set[T]) IsSubset(set2 Set[T]) bool {
	if len(set1) < len(set2) {
		return false
	}
	for x := range set2 {
		if !set1.Contains(x) {
			return false
		}
	}
	return true
}

func (set1 Set[T]) Equals(set2 Set[T]) bool {
	return len(set1) == len(set2) && set1.IsSubset(set2)
}

// intersect returns the common values in both set1 and set2
func (set1 Set[T]) Intersect(set2 Set[T]) Set[T] {
	same := New[T](len(set1))

	for k := range set2 {
		if set1.Contains(k) {
			same.Add(k)
		}
	}
	return same
}

// Diff returns a triple of sets
//
//	less: only in the 1st set (left-hand)
//	same: in both sets
//	more: only in the 2nd set
func (set1 Set[T]) Diff(set2 Set[T]) (less, same, more Set[T]) {
	less, same, more = New[T](0), New[T](len(set1)), New[T](0)

	for k := range set2 {
		if set1.Contains(k) {
			same.Add(k)
		} else {
			more.Add(k)
		}
	}

	for k := range set1 {
		if !same.Contains(k) {
			less.Add(k)
		}
	}
	return
}
