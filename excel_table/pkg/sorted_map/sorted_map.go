package sorted_map

import (
	"cmp"
	"slices"
)

type Object[K cmp.Ordered, V any] struct {
	Key   K
	Value V
}

type SortedMap[K cmp.Ordered, V any] interface {
	Get(K) (V, bool)
	Set(k K, v V)
	Len() int
	Ordered() []V
	ReNew()
	Range(ordered bool) []Object[K, V]
	History() History[K, V]
}

// Map[K]V(capacity) can return sorted V by K
func Map[K cmp.Ordered, V any](capacity int) SortedMap[K, V] {
	return &sortedMap[K, V]{
		keys:    make([]K, 0, capacity),
		storage: make(map[K]V, capacity),
		history: make(map[K]V, capacity),
	}
}

type sortedMap[K cmp.Ordered, V any] struct {
	keys    []K     // store keys for sort
	storage map[K]V // store values by keys
	history map[K]V // store all keys and values [no delete in ReNew]
}

func (sm *sortedMap[K, V]) Get(k K) (V, bool) {
	v, ok := sm.storage[k]
	return v, ok
}

func (sm *sortedMap[K, V]) Set(k K, v V) {
	if _, ok := sm.storage[k]; !ok {
		sm.keys = append(sm.keys, k)
	}
	sm.storage[k] = v
	sm.history[k] = v
}

func (sm *sortedMap[K, V]) Len() int {
	return len(sm.keys)
}

func (sm *sortedMap[K, V]) Ordered() []V {
	slices.Sort(sm.keys)

	s := make([]V, len(sm.keys))

	for i, k := range sm.keys {
		s[i] = sm.storage[k]
	}

	return s
}

func (sm *sortedMap[K, V]) ReNew() {
	capacity := len(sm.keys)
	sm.keys = make([]K, 0, capacity)
	sm.storage = make(map[K]V, capacity)
	// history will save
}

func (sm *sortedMap[K, V]) Range(ordered bool) []Object[K, V] {
	if ordered {
		slices.Sort(sm.keys)
	}

	s := make([]Object[K, V], len(sm.keys))

	for i, k := range sm.keys {
		s[i] = Object[K, V]{
			Key:   k,
			Value: sm.storage[k],
		}
	}

	return s
}

func (sm *sortedMap[K, V]) History() History[K, V] {
	return sm.history
}

type History[K cmp.Ordered, V any] map[K]V

func (h History[K, V]) Len() int {
	return len(h)
}

func (h History[K, V]) Get(k K) (V, bool) {
	v, ok := h[k]
	return v, ok
}

func (h History[K, V]) Range(ordered bool) []Object[K, V] {
	his := make([]Object[K, V], len(h))

	if !ordered {
		i := 0
		for k, v := range h {
			his[i] = Object[K, V]{
				Key:   k,
				Value: v,
			}
			i++
		}
		return his
	}

	// if need sorted
	keys := make([]K, 0, len(h))

	// insert in ascending order
	for k := range h {
		inserted := false
		for i := 0; i < len(keys); i++ {
			if k < keys[i] {
				keys = append(keys[:i], append([]K{k}, keys[i:]...)...)
				inserted = true
				break
			}
		}
		// for greater write last
		if !inserted {
			keys = append(keys, k)
		}
	}

	// write history
	for i, k := range keys {
		his[i] = Object[K, V]{
			Key:   k,
			Value: h[k],
		}
	}

	return his
}
