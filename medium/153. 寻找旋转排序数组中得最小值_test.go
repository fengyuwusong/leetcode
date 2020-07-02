package medium

import (
	"fmt"
	"testing"
)

/**

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0
*/

// 双指针二分查找缩小范围 时间复杂度O(logn) 空间复杂度O(1)
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		// 此处仅能判断右指针 与中间数对比 左指针判断有反例
		if nums[r] < nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

func TestFindMin(t *testing.T) {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}
