package lrucache

import "fmt"

type node[K comparable, V any] struct {
	data V
	next *node[K, V]
	prev *node[K, V]
	key  K
}

func newNode[K comparable, V any](data V, key K) *node[K, V] {
	return &node[K, V]{
		data: data,
		next: nil,
		prev: nil,
		key:  key,
	}
}

type DllNode[K comparable, V any] struct {
	head *node[K, V]
	tail *node[K, V]
}

func NewDlNode[K comparable, V any]() *DllNode[K, V] {
	return &DllNode[K, V]{}
}

func (dlNode *DllNode[K, V]) InsertAtHead(key K, data V) *node[K, V] {
	var newNode *node[K, V] = newNode(data, key)

	if dlNode.head == nil {
		dlNode.head = newNode
		dlNode.tail = newNode
		return newNode
	}

	newNode.next = dlNode.head
	dlNode.head.prev = newNode
	dlNode.head = newNode

	return newNode
}

func (dlNode *DllNode[K, V]) RemoveNode(node *node[K, V]) bool {
	if node == nil {
		return false
	}
	if dlNode.head == nil {
		return false
	}
	if dlNode.head == dlNode.tail && dlNode.head == node {
		dlNode.head = nil
		dlNode.tail = nil
		return true
	}
	if node == dlNode.head {
		dlNode.head = dlNode.head.next
		dlNode.head.prev = nil
		return true
	}
	if node == dlNode.tail {
		dlNode.tail = dlNode.tail.prev
		dlNode.tail.next = nil
		return true
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	return true
}

func (dlNode *DllNode[K, V]) Print() {
	fmt.Println("Printing list: ")
	curr := dlNode.head
	for curr != nil {
		fmt.Print(curr.data, " <-> ")
		curr = curr.next
	}
	fmt.Println("nil")
}
