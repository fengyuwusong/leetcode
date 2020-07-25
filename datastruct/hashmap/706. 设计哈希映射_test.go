package easy

import (
	"fmt"
	"testing"
)

// 使用 []int 作为bucket得元素 key value 分别依序存储再数组中，那么偶数索引取值即是key 奇数索引取值即时value
type MyHashMap struct {
	val [][]int
}

func hashFunc(key int) int {
	return key % 2069
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{val: make([][]int, 2069)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	hashNum := hashFunc(key)
	if this.val[hashNum] == nil {
		this.val[hashNum] = []int{key, value}
		return
	}
	for i := 0; i < len(this.val[hashNum]); i += 2 {
		if this.val[hashNum][i] == key {
			this.val[hashNum][i+1] = value
			return
		}
	}
	// 找不着
	this.val[hashNum] = append(this.val[hashNum], key, value)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hashNum := hashFunc(key)
	if this.val[hashNum] == nil {
		return -1
	}
	for i := 0; i < len(this.val[hashNum]); i += 2 {
		if this.val[hashNum][i] == key {
			return this.val[hashNum][i+1]
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	hashNum := hashFunc(key)
	if this.val[hashNum] == nil {
		return
	}
	for i := 0; i < len(this.val[hashNum]); i += 2 {
		if this.val[hashNum][i] == key {
			if i+2 < len(this.val[hashNum]) {
				this.val[hashNum] = append(this.val[hashNum][:i], this.val[hashNum][i+2:]...)
			} else {
				this.val[hashNum] = this.val[hashNum][:i]
			}
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func TestHashMap(t *testing.T) {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
	obj.Remove(2)
	fmt.Println(obj.Get(2))
}
