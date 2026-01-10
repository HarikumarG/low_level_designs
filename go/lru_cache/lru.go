package lrucache

import (
	"fmt"
	"sync"
)

type LRUCache[K comparable, V any] struct {
	cache       map[K]*node[K, V]
	dllNode     *DllNode[K, V]
	capacity    int16
	currentSize int16
	mu          sync.Mutex
}

func NewLRUCache[K comparable, V any](capacity int16) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		cache:       map[K]*node[K, V]{},
		dllNode:     NewDlNode[K, V](),
		capacity:    capacity,
		currentSize: 0,
	}
}

func (l *LRUCache[K, V]) Put(k K, v V) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.canInsert() {
		l.insert(k, v)
		l.currentSize++
	} else {
		l.remove(l.dllNode.tail.key, l.dllNode.tail)
		l.insert(k, v)
	}
}

func (l *LRUCache[K, V]) Get(k K) V {
	l.mu.Lock()
	defer l.mu.Unlock()
	fetchNode := l.cache[k]
	if fetchNode == nil {
		var z V
		return z
	}
	l.remove(k, fetchNode)
	l.insert(k, fetchNode.data)
	return fetchNode.data
}

func (l *LRUCache[K, V]) canInsert() bool {
	return l.currentSize < l.capacity
}

func (l *LRUCache[K, V]) insert(k K, v V) {
	l.cache[k] = l.dllNode.InsertAtHead(k, v)
}

func (l *LRUCache[K, V]) remove(k K, node *node[K, V]) {
	if node != nil {
		l.dllNode.RemoveNode(node)
		delete(l.cache, k)
	}
}

func (l *LRUCache[K, V]) Print() {
	cacheKeys := make([]K, 0, len(l.cache))
	for k := range l.cache {
		cacheKeys = append(cacheKeys, k)
	}
	fmt.Printf("Cache keys :%v\n", cacheKeys)
	l.dllNode.Print()
}
