package sort

import (
	"fmt"
	"testing"
)

// 插入排序
func insert(nums []int) []int {
	var i, j int
	// 有序区 [0, ...,  i-1]
	for i = 1; i < len(nums); i++ {
		// 无序区 [i, ...., len(nums)-1]
		for j = i - 1; j >= 0; j-- {
			// 将无序区第一个数插入有序区 找到有序区适合的位置
			if nums[j] < nums[i] {
				break
			}
		}

		if j != i-1 {
			// 将有序区[j, ...., i]数组前移 将nums[i] 放入nums[j]位置
			// k > j+1 因为nums[k-1] = nums[j] 否则会越界
			for k := i; k > j+1; k-- {
				nums[k], nums[k-1] = nums[k-1], nums[k]
			}
		}
	}
	return nums
}

func TestInsert(t *testing.T) {
	fmt.Printf("%v\n", insert([]int{3, 2, 5, 1, 7, 4}))
}
