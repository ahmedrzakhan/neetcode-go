package main

import "fmt"

// https://leetcode.com/problems/lru-cache/

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

// TC - O(1), SC - O(capacity)
func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, found := this.cache[key]; found {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, found := this.cache[key]; found {
		node.value = value
		this.moveToFront(node)
		return
	}

	if len(this.cache) == this.capacity {
		delete(this.cache, this.tail.key)
		this.removeTail()
	}

	newNode := &Node{key: key, value: value}
	this.cache[key] = newNode
	this.addToFront(newNode)
}

func (this *LRUCache) addToFront(node *Node) {
	node.next = this.head
	node.prev = nil

	if this.head != nil {
		this.head.prev = node
	}
	this.head = node

	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) moveToFront(node *Node) {
	if node == this.head {
		return
	}

	// Detach the node
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	// If node is the tail, update the tail
	if node == this.tail {
		this.tail = node.prev
	}

	// Move the node to the front
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node

	// If the list was empty, also set the tail
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) removeTail() {
	if this.tail == nil {
		return
	}

	if this.tail.prev != nil {
		this.tail.prev.next = nil
	} else {
		this.head = nil
	}
	this.tail = this.tail.prev
}

func mainLR() {
	cache := ConstructorLRU(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	fmt.Println(cache.Get(1)) // Returns 1
	cache.Put(4, 4)           // Evicts key 2
	fmt.Println(cache.Get(2)) // Returns -1 (not found)
	fmt.Println(cache.Get(3)) // Returns 3
	fmt.Println(cache.Get(4)) // Returns 4
}
