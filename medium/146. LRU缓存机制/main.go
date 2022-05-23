package main

import "container/list"

type LRUCache struct {
	max   int
	l     *list.List
	cache map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		max:   capacity,
		l:     list.New(),
		cache: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		this.l.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.cache[key]; ok {
		this.l.MoveToFront(ele)
		kv := ele.Value.(*entry)
		kv.value = value
	} else {
		ele := this.l.PushFront(&entry{
			key:   key,
			value: value,
		})
		this.cache[key] = ele
	}
	if len(this.cache) > this.max {
		back := this.l.Back()
		this.l.Remove(back)
		kv := back.Value.(*entry)
		delete(this.cache, kv.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	println(obj.Get(1))
	obj.Put(3, 3)
	println(obj.Get(1))
}
