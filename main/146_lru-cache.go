package main

// import (
// 	"fmt"
// )

// /**
// 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
// 写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
// 进阶:
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？
// 示例:
// LRUCache cache = new LRUCache( 2 ); //缓存容量
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // 返回  1
// cache.put(3, 3);    // 该操作会使得密钥 2 作废
// cache.get(2);       // 返回 -1 (未找到)
// cache.put(4, 4);    // 该操作会使得密钥 1 作废
// cache.get(1);       // 返回 -1 (未找到)
// cache.get(3);       // 返回  3
// cache.get(4);       // 返回  4
// */
// // LRUCache lru缓存
// type LRUCache struct {
// 	Cap  int              // 容量
// 	KMp  map[int]*DuLNode // 存储缓存key
// 	data *DuLNode
// }

// // DuLNode 双向链表，带首位空指针的双向链表
// type DuLNode struct {
// 	Head *DuLNode
// 	Tail *DuLNode
// 	Pre  *DuLNode
// 	Next *DuLNode
// 	Len  int
// 	Val  int
// }

// func (l LRUCache) string() string {
// 	tmp := l.data
// 	vals := []int{}
// 	for tmp != nil {
// 		vals = append(vals, tmp.Val)
// 		tmp = tmp.Next
// 	}
// 	return fmt.Sprintf("%d %v %v", l.Cap, l.KMp, vals)
// }

// // Remove ...
// func (root *DuLNode) Remove(node *DuLNode) error {
// 	if node.Pre != nil && node.Next != nil {
// 		fmt.Println("remove:", node.Val)
// 		node.Pre.Next = node.Next
// 		node.Next.Pre = node.Pre
// 		root.Len--
// 	} else {
// 		panic("can't remove")
// 	}
// 	return nil
// }

// // Add ...
// func (root *DuLNode) Add(node *DuLNode) *DuLNode {
// 	fmt.Println("add:", node.Val)
// 	node.Head = root.Head
// 	node.Tail = root.Tail
// 	if root.Head
// 	node.Pre = root.Tail.Pre
// 	root.Tail.Pre.Next = node
// 	node.Next = root.Tail
// 	root.Tail.Pre = node
// 	root.Len++
// 	return node
// }

// // Constructor ...
// func Constructor(capacity int) LRUCache {
// 	head := &DuLNode{}
// 	tail := &DuLNode{}
// 	head.Next = tail
// 	tail.Pre = head
// 	return LRUCache{
// 		Cap: capacity,
// 		KMp: map[int]*DuLNode{},
// 		data: &DuLNode{
// 			Head: head,
// 			Tail: tail,
// 		},
// 	}
// }

// // Get ...
// func (this *LRUCache) Get(key int) int {
// 	if node, ok := this.KMp[key]; ok {
// 		this.data.Remove(node)
// 		this.data.Add(node)
// 		return node.Val
// 	}
// 	return -1
// }

// // Put ...
// func (this *LRUCache) Put(key int, value int) {
// 	if node, ok := this.KMp[key]; ok {
// 		this.data.Remove(node)
// 		node.Val = value
// 		this.data.Add(node)
// 	}
// 	if this.data.Len >= this.Cap {
// 		delete(this.KMp, this.data.Head.Next.Val)
// 		this.data.Remove(this.data.Head.Next)
// 	}
// 	node := &DuLNode{
// 		Val: value,
// 	}
// 	node = this.data.Add(node)
// 	this.KMp[key] = node
// }

// /**
//  * Your LRUCache object will be instantiated and called as such:
//  * obj := Constructor(capacity);
//  * param_1 := obj.Get(key);
//  * obj.Put(key,value);
//  */
