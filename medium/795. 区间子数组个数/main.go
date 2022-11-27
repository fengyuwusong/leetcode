package main

import "fmt"

//给你一个整数数组 nums 和两个整数：left 及 right 。找出 nums 中连续、非空且其中最大元素在范围 [left, right] 内的子数组
//，并返回满足条件的子数组的个数。
//
// 生成的测试用例保证结果符合 32-bit 整数范围。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,1,4,3], left = 2, right = 3
//输出：3
//解释：满足条件的三个子数组：[2], [2, 1], [3]
//
//
// 示例 2：
//
//
//输入：nums = [2,9,2,5,6], left = 2, right = 8
//输出：7
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁹
// 0 <= left <= right <= 10⁹
//
// Related Topics数组 | 双指针
//
// 👍 248, 👎 0
//
//
//
//

// fengyuwusong 2022-11-24 11:06:50
//leetcode submit region begin(Prohibit modification and deletion)
func numSubarrayBoundedMax(nums []int, left, right int) (ans int) {
	i0, i1 := -1, -1
	for i, x := range nums {
		if x > right {
			i0 = i
		}
		if x >= left {
			i1 = i
		}
		ans += i1 - i0
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
}
