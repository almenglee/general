package general

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

type Map[K, V any] struct {
	_map map[uint64]_hash[K, V]
}

type _hash[K, V any] struct {
	Key K
	Val V
}

func NewMap[K, V any]() *Map[K, V] {
	return &Map[K, V]{
		_map: make(map[uint64]_hash[K, V]),
	}
}

func (m *Map[K, V]) Each(f func(K, V)) {
	for _, v := range m._map {
		f(v.Key, v.Val)
	}
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0)
	m.Each(func(k K, _ V) {
		keys = append(keys, k)
	})
	return keys
}

func (m *Map[K, V]) Values() []V {
	vals := make([]V, 0)
	m.Each(func(_ K, v V) {
		vals = append(vals, v)
	})
	return vals
}

func (m Map[K, V]) Get(key K) (V, bool) {
	hash := HashKey(key)
	if v, ok := m._map[hash]; ok {
		return v.Val, true
	}
	return *new(V), false
}

func (m *Map[K, V]) Set(key K, val V) {
	hash := HashKey(key)
	m._map[hash] = _hash[K, V]{key, val}
}

func (m *Map[K, V]) HashMap() map[uint64]_hash[K, V] {
	return m._map
}

// ============================================================

func HashKey(key any) uint64 {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		println("Map: Key Hashing failed")
		panic(err)
	}
	return fingerprint(buf.Bytes())
}

func fingerprint(b []byte) uint64 {
	hash := fnv.New64a()
	_, err := hash.Write(b)
	if err != nil {
		return 0
	}
	return hash.Sum64()
}
