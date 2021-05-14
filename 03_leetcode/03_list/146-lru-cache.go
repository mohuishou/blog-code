package list

import (
	"fmt"
)

type LRUCache struct {
	root     *LRUCacheNode
	length   int
	capacity int
}

type LRUCacheNode struct {
	next *LRUCacheNode
	prev *LRUCacheNode
	val  int
	key  int
}

func Constructor(capacity int) LRUCache {
	n := &LRUCacheNode{key: -1, val: -1}
	n.next = n
	n.prev = n
	return LRUCache{
		root:     n,
		capacity: capacity,
	}
}

func (this *LRUCache) find(key int) *LRUCacheNode {
	n := this.root.next
	for n != this.root {
		if n.key == key {
			// 将当前节点从现在的位置上摘除
			n.prev.next = n.next
			n.next.prev = n.prev
			// 将当前节点移动到第一个
			head := this.root.next
			this.root.next = n
			n.prev = this.root
			n.next = head
			head.prev = n
			return n
		}
		n = n.next
	}
	return nil
}

func (this *LRUCache) Get(key int) int {
	node := this.find(key)
	if node != nil {
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := this.find(key)
	if node != nil {
		node.val = value
		return
	}

	node = &LRUCacheNode{key: key, val: value}
	head := this.root.next
	this.root.next = node
	node.prev = this.root
	node.next = head
	head.prev = node

	if this.capacity == this.length {
		last := this.root.prev
		last.prev.next = this.root
		this.root.prev = last.prev
		return
	}

	this.length++
}

func (this *LRUCache) debug() {
	n := this.root.next
	fmt.Printf("head -> ")
	for n != this.root {
		fmt.Printf("%d:%d -> ", n.key, n.val)
		n = n.next
	}
	fmt.Printf("\n")
}
