package ch

import (
	"sort"
	"strconv"
)

type hash func(data []byte) uint32

type Map struct {
	hashFunc    hash
	numReplicas int
	keys        []int
	hashMap     map[int]string
}

func NewMap(numReplicas int, hashFunc hash) *Map {
	return &Map{
		hashFunc:    hashFunc,
		numReplicas: numReplicas,
		hashMap:     make(map[int]string),
	}
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.numReplicas; i++ {
			hash := int(m.hashFunc([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}

	sort.Ints(m.keys)
}

func (m *Map) Remove(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.numReplicas; i++ {
			hash := int(m.hashFunc([]byte(strconv.Itoa(i) + key)))
			idx := sort.Search(len(m.keys), func(i int) bool {
				return m.keys[i] >= hash
			})
			m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
			delete(m.hashMap, hash)
		}
	}
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hashFunc([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })

	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}
