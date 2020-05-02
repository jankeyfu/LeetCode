package src

import "container/list"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (46.92%)
 * Likes:    527
 * Dislikes: 0
 * Total Accepted:    49.9K
 * Total Submissions: 105.6K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 *
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 *
 *
 *
 * 进阶:
 *
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 *
 *
 *
 * 示例:
 *
 * LRUCache cache = new LRUCache( 2 缓存容量  )
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 *
 *
*/

// @lc code=start
type LRUCache struct {
	mp       map[int]*list.Element
	data     *list.List
	capacity int
}

// KV ...
type KV struct {
	k int
	v int
}

// Constructor ....
func Constructor2(capacity int) LRUCache {
	return LRUCache{
		mp:       make(map[int]*list.Element, capacity),
		data:     list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.mp[key]; ok {
		this.data.MoveToFront(e)
		return (e.Value).(KV).v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.mp[key]; ok {
		this.data.MoveToFront(e)
		e.Value = KV{k: key, v: value}
		return
	}
	if len(this.mp) == this.capacity {
		e := this.data.Back()
		this.data.Remove(e)
		delete(this.mp, e.Value.(KV).k)
	}
	this.mp[key] = this.data.PushFront(KV{k: key, v: value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
