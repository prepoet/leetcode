package leetcode

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	keys     map[int]*list.Element
	list     *list.List
}

type pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		keys:     make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.keys[key]; ok {
		this.list.MoveToFront(ele)
		return ele.Value.(pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.keys[key]; ok {
		ele.Value = pair{key: key, value: value}
		this.list.MoveToFront(ele)
	} else {
		ele := this.list.PushFront(pair{key: key, value: value})
		this.keys[key] = ele
	}
	if this.list.Len() > this.capacity {
		ele := this.list.Back()
		this.list.Remove(ele)
		delete(this.keys, ele.Value.(pair).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
