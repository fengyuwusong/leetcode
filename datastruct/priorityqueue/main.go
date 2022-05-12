package main

import "fmt"

// 最大堆实现优先级队列（Priority Queue） https://labuladong.github.io/algo/2/21/62/
// 优先级队列这种数据结构有一个很有用的功能，你插入或者删除元素的时候，
// 元素会自动排序，这底层的原理就是二叉堆的操作。
// 数据结构的功能无非增删查改，优先级队列有两个主要 API，
// 分别是 insert 插入一个元素和 delMax 删除最大元素（如果底层用最小堆，那么就是 delMin）。

// 最大堆的性质是：每个节点都大于等于它的两个子节点。
// 对于最大堆，会破坏堆性质的有两种情况：
// 1、如果某个节点 A 比它的子节点（中的一个）小，那么 A 就不配做父节点，应该下去，下面那个更大的节点上来做父节点，这就是对 A 进行下沉。
// 2、如果某个节点 A 比它的父节点大，那么 A 不应该做子节点，应该把父节点换下来，自己去做父节点，这就是对 A 的上浮。

// 插入和删除元素的时间复杂度为 O(logK)，K 为当前二叉堆（优先级队列）中的元素总数。

type PriorityQueue struct {
	arr  []int
	size int
}

// 堆的性质
// 每个节点的父节点索引=当前索引/2
// 左孩子索引 = 当前索引*2
// 右孩子索引 = 当前索引*2+1

// 父节点的索引
func (p *PriorityQueue) parent(root int) int {
	return root / 2
}

// 左孩子的索引
func (p *PriorityQueue) left(root int) int {
	return root * 2
}

// 右孩子的索引
func (p *PriorityQueue) right(root int) int {
	return root*2 + 1
}

// 索引0不使用 多分配一个空间
func NewPriorityQueue(cap int) *PriorityQueue {
	return &PriorityQueue{
		arr: make([]int, cap+1),
	}
}

// 返回队列中的最大元素
func (p *PriorityQueue) max() int {
	return p.arr[1]
}

// 插入元素 加入最后让其上浮
func (p *PriorityQueue) insert(i int) {
	p.size++
	// 先把新元素加到最后
	p.arr[p.size] = i
	// 然后让它上浮到正确的位置
	p.swim(p.size)
}

// 删除并返回队列中的最大元素 将堆顶元素A和堆底元素B对调 删除A后让B下沉到正确位置
func (p *PriorityQueue) delMax() int {
	// 对调堆顶和堆底
	max := p.max()
	p.swap(1, p.size)
	// 删除堆底
	p.arr[p.size] = 0
	p.size--

	// 让堆顶下沉
	p.sink(1)
	return max
}

// 上浮第x个元素 以维护最大堆性质
func (p *PriorityQueue) swim(x int) {
	// 如果浮到堆顶，就不能再上浮了
	for x > 1 && p.less(p.parent(x), x) {
		// 如果第 x 个元素比上层大
		// 将 x 换上去
		p.swap(p.parent(x), x)
		x = p.parent(x)
	}
}

// 下沉第x个元素 以维护最大堆性质
func (p *PriorityQueue) sink(x int) {
	// 如果沉到堆底 则不继续下沉
	for p.left(x) < p.size {
		// 获取左右节点最小孩子
		min := p.left(x)
		if p.less(p.right(x), min) {
			min = p.right(x)
		}

		// x比左右孩子大 无需下沉
		if p.less(min, x) {
			break
		}

		// 下沉x
		p.swap(x, min)
		x = min
	}
}

// 交换数组的两个元素
func (p *PriorityQueue) swap(i, j int) {
	p.arr[i], p.arr[j] = p.arr[j], p.arr[i]
}

// 判断 arr[i] 是否比 arr[j] 小
func (p *PriorityQueue) less(i, j int) bool {
	return p.arr[i] < p.arr[j]
}

func main() {
	pq := NewPriorityQueue(5)
	pq.insert(3)
	pq.insert(20)
	pq.insert(5)
	pq.insert(1)
	pq.insert(9)

	fmt.Printf("pq.arr: %v\n", pq.arr)
	fmt.Printf("pq.delMax(): %v\n", pq.delMax())
	pq.insert(7)
	fmt.Printf("pq.arr: %v\n", pq.arr)
}
