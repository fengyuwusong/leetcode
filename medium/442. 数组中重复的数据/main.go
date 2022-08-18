package main

import "fmt"

//给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次
//的整数，并以数组形式返回。
//
// 你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[2,3]
//
//
// 示例 2：
//
//
//输入：nums = [1,1,2]
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[]
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
// nums 中的每个元素出现 一次 或 两次
//
// Related Topics数组 | 哈希表
//
// 👍 652, 👎 0
//
//
//
//

// fengyuwusong 2022-08-19 01:24:42
//leetcode submit region begin(Prohibit modification and deletion)
func findDuplicates(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] < 0 {
			ans = append(ans, index+1)
		}
		nums[index] = -nums[index]
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates([]int{1, 1, 2}))
	fmt.Println(findDuplicates([]int{1}))
}
