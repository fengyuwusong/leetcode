package easy

import (
	"testing"
)

/**
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-insert-position
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for index, num := range nums {
		if num == target {
			return index
		}
		if num > target {
			return index
		}
	}
	return len(nums)
}

func TestSearchInsert(t *testing.T) {
	println(searchInsert([]int{1, 3, 5, 6}, 7))
}
