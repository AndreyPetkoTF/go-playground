package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
[null,null,   null,  null,  null,  null,  -1,  null,  19, 17, null,  -1, null,  null, null,  -1,  null,  -1, 5,  -1, 12, null,  null, 3,  5,   5,   null,  null, 1,  null,   -1, null,  30, 5,  30, null,  null,   null,  -1,  null,  -1, 24,  null,  null,  18, null, null,  null,  null, -1, null,  null,  18, null,  null,  -1,  null,  null,  null,   null,  null,  18, null, null,   24,  null,  4,  -1, -1, null, 12, -1,  null,  null,   null,  null,   29, null,  null, null,  null,  17, 22, 18, null,   null,  null,   24,  null,  null,  null,  20, null,   null, null,  29, 18, 18, null, null,  null, null,  20, null, null, null,  null,  null,  null,   null]


**/
func main() {
	cache := NewLRUCache(10)

	cache.Put(10, 13)
	cache.Put(3, 17)
	cache.Put(6, 11)
	cache.Put(10, 5)
	cache.Put(9, 10)

	fmt.Println(cache.Get(13)) // -1

	cache.Put(2, 19)

	cache.Get(2)
	cache.Get(3)

	cache.Put(5, 25)

	fmt.Println(cache.Get(8))

	cache.Put(9, 22)
	cache.Put(5, 5)
	cache.Put(1, 30)
	cache.Get(11)
	cache.Put(9, 12)
	cache.Get(7)
	cache.Get(5)
	cache.Get(8)
	cache.Get(9)
	cache.Put(4, 30)
	cache.Put(9, 3)
	cache.Get(9)
	cache.Get(10)
	cache.Get(10)

	//cache.Put(3, 3)

	//fmt.Println(cache.Get(2)) // -1

}

type Node struct {
	Val  int
	Next *Node
}

func (lru *LRUCache) String() string {
	var parts []string
	currentNode := lru.getOrderList

	for currentNode != nil {
		parts = append(parts, strconv.Itoa(currentNode.Val)+"("+strconv.Itoa(lru.cache[currentNode.Val])+")")
		currentNode = currentNode.Next
	}

	return strings.Join(parts, "->")
}

type LRUCache struct {
	cache        map[int]int
	capacity     int
	getOrderList *Node
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{make(map[int]int), capacity, nil}
}

func (lru *LRUCache) Get(key int) int {
	if lru.getOrderList == nil {
		return -1
	}

	if lru.getOrderList.Val == key {
		return lru.cache[key]
	}

	item := lru.getOrderList

	var preItem *Node
	for item.Next != nil {
		if item.Next.Val == key {
			preItem = item
			break
		}

		item = item.Next
	}

	if preItem == nil {
		return -1
	}

	keyItem := preItem.Next
	preItem.Next = preItem.Next.Next

	//keyItem.Next = head
	lru.getOrderList = keyItem

	return lru.cache[key]
}

func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.cache[key]; ok {
		lru.cache[key] = value
		lru.Get(key)
	} else {
		if lru.getOrderList != nil {
			item := &Node{
				Val:  key,
				Next: lru.getOrderList,
			}

			lru.getOrderList = item
		} else {
			lru.getOrderList = &Node{
				Val:  key,
				Next: nil,
			}
		}

		lru.cache[key] = value

		if len(lru.cache) > lru.capacity {
			// remove element
			item := lru.getOrderList

			for item.Next != nil && item.Next.Next != nil {
				item = item.Next
			}

			item.Next = nil
			// remove element
		}
	}
}
