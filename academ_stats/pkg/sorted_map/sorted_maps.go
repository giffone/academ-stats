package sorted_map

import (
	"cmp"
	"slices"
	"strconv"
)

type Storage[K cmp.Ordered, V any] interface {
	Storage() Maps[K, V]
	ReadByDepth(depth int) []Maps[K, V]
	Len() int
}

type Maps[K cmp.Ordered, V any] interface {
	Get(key K) (*Data[K, V], bool)
	Set(key K, value V)
	GetOrSet(key K, value V) *Data[K, V]
	Read() ([]K, []*Data[K, V])
	Range() []*Data[K, V]
	Len() int
	Sort()
	SortFunc(f func(a, b K) int)
	Parent() Maps[K, V]
}

// NewStorage[K]V(capacity) can return sorted V by K
func NewStorage[K cmp.Ordered, V any](capacity int) Storage[K, V] {
	h := head[K, V]{
		depthStorage: make(map[int][]Maps[K, V]),
		capacity:     capacity,
	}
	// create storage
	h.mapStorage = h.newStorage(0)

	return &h
}

func (h *head[K, V]) newStorage(depthKey int) *storage[K, V] {
	return &storage[K, V]{
		depthKey: depthKey,
		ordKeys:  make([]K, 0, h.capacity),
		data:     make(map[K]*Data[K, V], h.capacity),
		head:     h,
	}
}

func (h *head[K, V]) addStorageDepth(s *storage[K, V]) {
	s.depthFlag = true
	h.depthStorage[s.depthKey] = append(h.depthStorage[s.depthKey], s)
}

type head[K cmp.Ordered, V any] struct {
	mapStorage   Maps[K, V]           // map[key]value
	depthStorage map[int][]Maps[K, V] // map[depth][]map[key]value - saved by n-iteration in depth
	capacity     int                  // storage capacity
}

func (h *head[K, V]) Storage() Maps[K, V] {
	return h.mapStorage
}

func (h *head[K, V]) ReadByDepth(depth int) []Maps[K, V] {
	return h.depthStorage[depth]
}

func (h *head[K, V]) Len() int {
	return len(h.depthStorage)
}

type Data[K cmp.Ordered, V any] struct {
	Value V
	Prev  Maps[K, V]
	Next  Maps[K, V]
}

type storage[K cmp.Ordered, V any] struct {
	depthKey  int               // number of cycles in depth
	depthFlag bool              // struct added to storage_depth
	ordKeys   []K               // store keys for sort
	data      map[K]*Data[K, V] // store values by keys
	parent    *storage[K, V]    // upstream struct reference
	head      *head[K, V]       // home struct reference
	sort      bool              // if need sort
	sortF     func(a, b K) int  //
}

func (m *storage[K, V]) Get(key K) (*Data[K, V], bool) {
	v, ok := m.data[key]
	if !ok {
		return nil, false
	}
	return v, ok
}

func (m *storage[K, V]) Set(key K, value V) {
	if _, ok := m.data[key]; !ok {
		// save new key
		m.ordKeys = append(m.ordKeys, key)

		// create next storage
		next := m.head.newStorage(m.depthKey + 1)
		next.parent = m

		// create data
		m.data[key] = &Data[K, V]{
			Prev: m.parent,
			Next: next,
		}

		// add to depth_storage
		if !m.depthFlag {
			m.head.addStorageDepth(m)
		}
	}

	m.data[key].Value = value
}

func (m *storage[K, V]) GetOrSet(key K, value V) *Data[K, V] {
	v, ok := m.data[key]
	if !ok {
		m.Set(key, value)
		// update
		v = m.data[key]
	}
	return v
}

func (m *storage[K, V]) Read() ([]K, []*Data[K, V]) {
	return m.ordKeys, m.Range()
}

func (m *storage[K, V]) Range() []*Data[K, V] {
	if m.sort {
		m.order()
	}

	vals := make([]*Data[K, V], len(m.ordKeys))

	for i, k := range m.ordKeys {
		vals[i] = m.data[k]
	}

	return vals
}

func (m *storage[K, V]) Len() int {
	return len(m.ordKeys)
}

func (m *storage[K, V]) Sort() {
	m.sort = true
}

func (m *storage[K, V]) SortFunc(f func(a, b K) int) {
	m.sortF = f
	m.sort = true
}

func (m *storage[K, V]) Parent() Maps[K, V] {
	return m.parent
}

func (m *storage[K, V]) order() {
	if len(m.ordKeys) == 0 {
		return
	}
	switch any(m.ordKeys[0]).(type) {
	case string:
		if m.sortF == nil {
			m.sortF = func(a, b K) int {
				num1, err := strconv.ParseFloat(any(a).(string), 64)
				if err == nil {
					if num2, err := strconv.ParseFloat(any(b).(string), 64); err == nil {
						return cmp.Compare(num1, num2)
					}
				}
				return cmp.Compare(a, b)
			}
		}
		slices.SortFunc(m.ordKeys, m.sortF)
	default:
		slices.Sort(m.ordKeys)
	}
}
