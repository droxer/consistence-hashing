package ch

import (
	"sort"
	"strconv"
)

type hash func(data []byte) uint32

type Map struct {
	hashFunc    hash
	numReplicas int
	nodes       []int
	hashMap     map[int]string
}

func NewMap(numReplicas int, hashFunc hash) *Map {
	return &Map{
		hashFunc:    hashFunc,
		numReplicas: numReplicas,
		hashMap:     make(map[int]string),
	}
}

func (m *Map) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < m.numReplicas; i++ {
			hash := int(m.hashFunc([]byte(strconv.Itoa(i) + node)))
			m.nodes = append(m.nodes, hash)
			m.hashMap[hash] = node
		}
	}

	sort.Ints(m.nodes)
}

func (m *Map) Remove(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < m.numReplicas; i++ {
			hash := int(m.hashFunc([]byte(strconv.Itoa(i) + node)))
			idx := sort.Search(len(m.nodes), func(i int) bool {
				return m.nodes[i] >= hash
			})
			m.nodes = append(m.nodes[:idx], m.nodes[idx+1:]...)
			delete(m.hashMap, hash)
		}
	}
}

func (m *Map) Get(key string) string {
	if len(m.nodes) == 0 {
		return ""
	}

	hash := int(m.hashFunc([]byte(key)))
	idx := sort.Search(len(m.nodes), func(i int) bool { return m.nodes[i] >= hash })

	if idx == len(m.nodes) {
		idx = 0
	}

	return m.hashMap[m.nodes[idx]]
}
