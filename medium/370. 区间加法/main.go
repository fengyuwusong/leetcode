package main

import "fmt"

type Difference struct {
	diffArray []int // 差分数组
}

func NewDifference(nums []int) *Difference {
	// 构造差分数组
	diffArray := make([]int, len(nums))
	diffArray[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diffArray[i] = nums[i] - diffArray[i-1]
	}
	return &Difference{diffArray: diffArray}
}

func (d *Difference) increment(i, j, val int) {
	d.diffArray[i] += val
	if j+1 < len(d.diffArray) {
		d.diffArray[j+1] -= val
	}
}

func (d *Difference) result() []int {
	// 根据差分数组反推原数组并返回
	res := make([]int, len(d.diffArray))
	res[0] = d.diffArray[0]
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + d.diffArray[i]
	}
	return res
}

// 差分数组: 用于对数组区间元素进行批量操作 同时时间复杂度O(1)
func getModifiedArray(length int, updates [][]int) []int {
	// 初始化数组为0
	nums := make([]int, length)
	// 构造差分解法
	df := NewDifference(nums)

	for _, update := range updates {
		i := update[0]
		j := update[1]
		val := update[2]
		df.increment(i, j, val)
		fmt.Printf("%v\n", df.result())
	}
	return df.result()
}

func main() {
	fmt.Printf("%v\n", getModifiedArray(5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}))
}
