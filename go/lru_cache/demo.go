package lrucache

import "fmt"

func LruDemo() {
	operations := []string{"LRUCache", "put", "put", "get", "put", "get", "get"}
	values := [][]int16{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {1}}

	var lru *LRUCache

	for index, ops := range operations {
		switch ops {
		case "LRUCache":
			lru = NewLRUCache(values[index][0])
			fmt.Println("Created LRUCache capacity", values[index][0])
		case "put":
			lru.Put(values[index][0], values[index][1])
			fmt.Printf("\nput(%d) -> %d\n", values[index][0], values[index][1])
			lru.Print()
		case "get":
			res := lru.Get(values[index][0])
			fmt.Printf("\nget(%d) -> %d\n", values[index][0], res)
			lru.Print()
		}
	}
}
