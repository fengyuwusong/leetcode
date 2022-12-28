package main

import (
	"fmt"
	"strings"
)

//给你三个整数数组 nums1、nums2 和 nums3 ，请你构造并返回一个 元素各不相同的 数组，且由 至少 在 两个 数组中出现的所有值组成。数组中的
//元素可以按 任意 顺序排列。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
//输出：[3,2]
//解释：至少在两个数组中出现的所有值为：
//- 3 ，在全部三个数组中都出现过。
//- 2 ，在数组 nums1 和 nums2 中出现过。
//
//
// 示例 2：
//
//
//输入：nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
//输出：[2,3,1]
//解释：至少在两个数组中出现的所有值为：
//- 2 ，在数组 nums2 和 nums3 中出现过。
//- 3 ，在数组 nums1 和 nums2 中出现过。
//- 1 ，在数组 nums1 和 nums3 中出现过。
//
//
// 示例 3：
//
//
//输入：nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
//输出：[]
//解释：不存在至少在两个数组中出现的值。
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length, nums3.length <= 100
// 1 <= nums1[i], nums2[j], nums3[k] <= 100
//
//
// Related Topics 数组 哈希表 👍 22 👎 0

// author: fengyuwusong
// create time: 2022-12-29 00:29:20
// leetcode submit region begin(Prohibit modification and deletion)
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	// todo 可通过位运算优化
	memo := make(map[int]string)
	for _, i := range nums1 {
		memo[i] = "1"
	}
	for _, i := range nums2 {
		if memo[i] == "" {
			memo[i] = "01"
		} else if memo[i] == "1" {
			memo[i] = "11"
		}
	}

	for _, i := range nums3 {
		if memo[i] == "" {
			memo[i] = "001"
		} else if len(memo[i]) == 2 {
			memo[i] = memo[i] + "1"
		} else if len(memo[i]) == 1 {
			memo[i] = memo[i] + "01"
		}
	}
	var ans []int
	for k, v := range memo {
		if strings.Count(v, "1") >= 2 {
			ans = append(ans, k)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(twoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))
	fmt.Println(twoOutOfThree([]int{3, 1}, []int{2, 3}, []int{1, 2}))
	fmt.Println(twoOutOfThree([]int{1, 2, 2}, []int{4, 3, 3}, []int{5}))
}
