package easy

import (
	"fmt"
	"testing"
)

/**
给定一个整数数组，判断是否存在重复元素。

如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。



示例 1:

输入: [1,2,3,1]
输出: true
示例 2:

输入: [1,2,3,4]
输出: false
示例 3:

输入: [1,1,1,3,3,4,3,2,4,2]
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 二维遍历 时间复杂度O(n^2) 空间复杂度O(1)
// 解法2 排序遍历 时间复杂度O(nlogn) 空间复杂度O(1)
// 解法3 hash记录 时间复杂度O(n) 空间复杂度O(n)
func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]interface{}, len(nums))
	for _, n := range nums {
		if _, ok := hashMap[n]; ok {
			return true
		}
		hashMap[n] = struct{}{}
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
