package easy

import (
	"fmt"
	"testing"
)

/**
Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example 1:

Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums being 2.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,1,2,2,3,0,4,2], val = 2,

Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.

Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return -1
	}
	var lastIndex int = len(nums)
	for i := 0; i < len(nums); i++ {
		// 跳出循环条件
		if lastIndex == i {
			break
		}
		// 相等则前后索引元素替换 lastIndex、i 后退重新遍历当前元素
		if nums[i] == val {
			lastIndex--
			nums[i] = nums[lastIndex]
			i--
		}
	}
	return lastIndex
}
func TestRemoveElement(t *testing.T) {

	f := func(nums []int, length int) {
		println(length)
		for i := 0; i < length; i++ {
			fmt.Printf("%d=>", nums[i])
		}
		println()
	}
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	length := removeElement(nums, 2)
	f(nums, length)
	nums = []int{1}
	length = removeElement(nums, 1)
	f(nums, length)
}
