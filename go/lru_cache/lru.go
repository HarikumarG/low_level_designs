package lrucache

import "fmt"

type LRUCache struct {
	cache       map[int16]*node
	dllNode     *DllNode
	capacity    int16
	currentSize int16
}

func NewLRUCache(capacity int16) *LRUCache {
	return &LRUCache{
		cache:       map[int16]*node{},
		dllNode:     NewDlNode(),
		capacity:    capacity,
		currentSize: 0,
	}
}

func (l *LRUCache) Put(k int16, v int16) {
	if l.canInsert() {
		l.insert(k, v)
		l.currentSize++
	} else {
		l.remove(l.dllNode.tail.key, l.dllNode.tail)
		l.insert(k, v)
	}
}

func (l *LRUCache) Get(k int16) int16 {
	fetchNode := l.cache[k]
	if fetchNode == nil {
		return -1
	}
	l.remove(k, fetchNode)
	l.insert(k, fetchNode.data)
	return fetchNode.data
}

func (l *LRUCache) canInsert() bool {
	return l.currentSize < l.capacity
}

func (l *LRUCache) insert(k, v int16) {
	l.cache[k] = l.dllNode.InsertAtHead(k, v)
}

func (l *LRUCache) remove(k int16, node *node) {
	if node != nil {
		l.dllNode.RemoveNode(node)
		delete(l.cache, k)
	}
}

func (l *LRUCache) Print() {
	cacheKeys := make([]int16, 0, len(l.cache))
	for k := range l.cache {
		cacheKeys = append(cacheKeys, k)
	}
	fmt.Printf("Cache keys :%v\n", cacheKeys)
	l.dllNode.Print()
}
