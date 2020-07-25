package 剑指offer

import (
	"fmt"
	"testing"
)

/**
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法同 leetcode 217
// 解法4 原地置换 由于所有数字均小于n-1 那么可以按位数遍历找到当前位数对应得值为止  时间复杂度 O(n) 空间复杂度 O(1)
func findRepeatNumber(nums []int) int {
	for i, n := range nums {
		// 如当前索引值不等于当前索引 则置换当前值到对应索引直到与当前索引相等为止
		for i != n {
			if n == nums[n] {
				return n
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

// 解法3 哈希表
func findRepeatNumber2(nums []int) int {
	hashMap := make(map[int]interface{}, len(nums))
	for _, n := range nums {
		if _, ok := hashMap[n]; ok {
			return n
		}
		hashMap[n] = struct{}{}
	}
	return -1
}

func TestFindRepeatNumber(t *testing.T) {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}
