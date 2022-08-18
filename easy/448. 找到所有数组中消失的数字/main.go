package main

import (
	"fmt"
	"math"
)

//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数
//字，并以数组的形式返回结果。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[5,6]
//
//
// 示例 2：
//
//
//输入：nums = [1,1]
//输出：[2]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁵
// 1 <= nums[i] <= n
//
//
// 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
// Related Topics数组 | 哈希表
//
// 👍 1054, 👎 0
//
//
//
//

// fengyuwusong 2022-08-19 01:06:31
//leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] < 0 {
			// 之前已经吧对应索引的元素变为负数了
			//	说明nums重复出现了两次 让索引继续保持负数
		} else {
			nums[index] = -nums[index]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			// 说明没有元素和这个索引对应 说明这个索引是缺失的元素
			ans = append(ans, i+1)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
