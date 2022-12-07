package main

import "fmt"

//给你两个长度可能不等的整数数组 nums1 和 nums2 。两个数组中的所有值都在 1 到 6 之间（包含 1 和 6）。
//
// 每次操作中，你可以选择 任意 数组中的任意一个整数，将它变成 1 到 6 之间 任意 的值（包含 1 和 6）。
//
// 请你返回使 nums1 中所有数的和与 nums2 中所有数的和相等的最少操作次数。如果无法使两个数组的和相等，请返回 -1 。
//
//
//
// 示例 1：
//
// 输入：nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
//输出：3
//解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
//- 将 nums2[0] 变为 6 。 nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2] 。
//- 将 nums1[5] 变为 1 。 nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2] 。
//- 将 nums1[2] 变为 2 。 nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2] 。
//
//
// 示例 2：
//
// 输入：nums1 = [1,1,1,1,1,1,1], nums2 = [6]
//输出：-1
//解释：没有办法减少 nums1 的和或者增加 nums2 的和使二者相等。
//
//
// 示例 3：
//
// 输入：nums1 = [6,6], nums2 = [1]
//输出：3
//解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
//- 将 nums1[0] 变为 2 。 nums1 = [2,6], nums2 = [1] 。
//- 将 nums1[1] 变为 2 。 nums1 = [2,2], nums2 = [1] 。
//- 将 nums2[0] 变为 4 。 nums1 = [2,2], nums2 = [4] 。
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 10⁵
// 1 <= nums1[i], nums2[i] <= 6
//
// Related Topics贪心 | 数组 | 哈希表 | 计数
//
// 👍 58, 👎 0
//
//
//
//

// fengyuwusong 2022-12-07 00:19:16
//leetcode submit region begin(Prohibit modification and deletion)
func help(h1 [7]int, h2 [7]int, diff int) (res int) {
	h := [7]int{}
	for i := 1; i < 7; i++ {
		// 计算h1[i]变为6能增大的数量对应个数
		h[6-i] += h1[i]
		// 计算h2[i]变为1能增大的数量对应个数
		h[i-1] += h2[i]
	}
	for i := 5; i > 0 && diff > 0; i-- {
		// 计算diff所需的变为对应数量的最小个数
		t := min((diff+i-1)/i, h[i])
		res += t
		diff -= t * i
	}
	return res
}

func minOperations(nums1 []int, nums2 []int) (ans int) {
	n, m := len(nums1), len(nums2)
	// 6*n所能表示最大的数 m所能表示最小的数
	if 6*n < m || 6*m < n {
		return -1
	}
	// 记录两个数组1-6的个数
	var cnt1, cnt2 [7]int
	// 记录diff差距值
	diff := 0
	for _, i := range nums1 {
		cnt1[i]++
		diff += i
	}
	for _, i := range nums2 {
		cnt2[i]++
		diff -= i
	}
	// 没有差距 返回0
	if diff == 0 {
		return 0
	}
	// 计算步数
	if diff > 0 {
		return help(cnt2, cnt1, diff)
	}
	return help(cnt1, cnt2, -diff)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minOperations([]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2}))
}
