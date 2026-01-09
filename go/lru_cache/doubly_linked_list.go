package lrucache

import "fmt"

type node struct {
	data int16
	next *node
	prev *node
	key  int16
}

func newNode(data int16, key int16) *node {
	return &node{
		data: data,
		next: nil,
		prev: nil,
		key:  key,
	}
}

type DllNode struct {
	head *node
	tail *node
}

func NewDlNode() *DllNode {
	return &DllNode{}
}

func (dlNode *DllNode) InsertAtHead(key int16, data int16) *node {
	newNode := newNode(data, key)

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

func (dlNode *DllNode) RemoveNode(node *node) bool {
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

func (dlNode *DllNode) Print() {
	fmt.Println("Printing list: ")
	curr := dlNode.head
	for curr != nil {
		fmt.Print(curr.data, " <-> ")
		curr = curr.next
	}
	fmt.Println("nil")
}
