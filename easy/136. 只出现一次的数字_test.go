package easy

import (
	"fmt"
	"testing"
)

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 暴力
// 解法2 排序
// 解法3 hash
// 解法4 异或 时间复杂度O(n) 空间复杂度O(1)
// 异或运算: b ^ b = 0 => a ^ b ^ b = a 则可将整个数组遍历进行异或运算 最后得出的元素即仅出现一次的元素
func singleNumber(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}

func TestSingleNumber(t *testing.T) {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
