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
func (set Set[T]) Add(xs ...T) {
	for _, x := range xs {
		set[x] = struct{}{}
	}
}

func (set Set[T]) Remove(xs ...T) {
	for _, x := range xs {
		delete(set, x)
	}
}

// queries
func (set Set[T]) Copy() Set[T] {
	set2 := New[T](len(set))
	for x := range set {
		set2.Add(x)
	}
	return set2
}

func (set Set[T]) Count() int {
	return len(set)
}

func (set Set[T]) Any() bool {
	return len(set) != 0
}

func (set Set[T]) Has(xs ...T) bool {
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

func (set1 Set[T]) IsSubset(set2 Set[T]) bool {
	if len(set1) < len(set2) {
		return false
	}
	for x := range set2 {
		if !set1.Has(x) {
			return false
		}
	}
	return true
}

func (set1 Set[T]) Equals(set2 Set[T]) bool {
	return len(set1) == len(set2) && set1.IsSubset(set2)
}

// Diff returns a triple of sets
//
//	less: only in the 1st set (left-hand)
//	same: in both sets
//	more: only in the 2nd set
func (set1 Set[T]) Diff(set2 Set[T]) (less, same, more Set[T]) {
	less, same, more = New[T](0), New[T](len(set1)), New[T](0)

	for k := range set2 {
		if set1.Has(k) {
			same.Add(k)
		} else {
			more.Add(k)
		}
	}

	for k := range set1 {
		if !same.Has(k) {
			less.Add(k)
		}
	}
	return
}
