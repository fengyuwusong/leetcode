package main

import "fmt"

//给你一个正整数数组 nums，请你帮忙从该数组中找出能满足下面要求的 最长 前缀，并返回该前缀的长度：
//
//
// 从前缀中 恰好删除一个 元素后，剩下每个数字的出现次数都相同。
//
//
// 如果删除这个元素后没有剩余元素存在，仍可认为每个数字都具有相同的出现次数（也就是 0 次）。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,2,1,1,5,3,3,5]
//输出：7
//解释：对于长度为 7 的子数组 [2,2,1,1,5,3,3]，如果我们从中删去 nums[4] = 5，就可以得到 [2,2,1,1,3,3]，里面每个数
//字都出现了两次。
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1,2,2,2,3,3,3,4,4,4,5]
//输出：13
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
// Related Topics数组 | 哈希表
//
// 👍 64, 👎 0
//
//
//
//

// fengyuwusong 2022-08-18 00:48:06
//leetcode submit region begin(Prohibit modification and deletion)
func maxEqualFreq(nums []int) int {
	var ans int
	freq := map[int]int{}
	count := map[int]int{}
	maxFreq := 0
	for i, num := range nums {
		if count[num] > 0 {
			freq[count[num]]--
		}
		count[num]++
		maxFreq = max(maxFreq, count[num])
		freq[count[num]]++
		// 所有数出现次数均是1
		if maxFreq == 1 ||
			// 所有数的出现次数都是maxFreq 或者 maxFreq-1, 并且最大出现次数的数只有一个 (即仅有一个数出现maxFreq或maxFreq-1次, 那么删除一个数， 剩下的数的出现次数都是maxFreq-1次)
			freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 ||
			// 除开一个数 其他所有数的出现次数都是maxFreq 并且该数的出现次数为1
			freq[maxFreq]*maxFreq+1 == i+1 && freq[1] == 1 {
			ans = max(ans, i+1)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(maxEqualFreq([]int{2, 2, 1, 1, 5, 3, 3, 5}))
	fmt.Println(maxEqualFreq([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5}))
}
