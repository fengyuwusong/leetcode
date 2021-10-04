package main

import "fmt"

/**
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。



进阶：

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？


示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]


提示：

1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
0 <= k <= 105


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1：拆分合并数组
// 解法2：环状替换
// 解法3：翻转数组 1. 全部翻转 2. 翻转 [0, k mod n - 1] 3. 翻转[k mod n, n - 1]

// 解法2
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	// 使用单独的变量 count 跟踪当前已经访问的元素数量，当 count=n 时，结束遍历过程。
	for start, count := 0, 0; count < n; start++ {
		pre, cur := nums[start], start
		// 回到起点则停止循环
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			// 对调 cur next 同时将cur置为next
			nums[next], pre, cur = pre, nums[next], next
			count++
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
