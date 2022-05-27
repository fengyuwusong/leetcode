package main

// 前缀和 (preSum[i] - preSum[j]) % k == 0 等价于 preSum[i] % k == preSum[j] % k。
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	// 前缀和数组
	preSum := make([]int, n+1)
	// 前缀和与k的余数到索引的映射 方便查找所需前缀和
	count := make(map[int]int)
	// 构造前缀和
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	for i := 0; i < len(preSum); i++ {
		// 哈希表记录余数
		val := preSum[i] % k
		// 如该余数无索引 则记录下来
		if _, ok := count[val]; !ok {
			count[val] = i
		}
		// 已有索引这跳过 最长数组记录相同索引起点应尽可能小
	}

	for i := 1; i < len(preSum); i++ {
		need := preSum[i] % k
		if j, ok := count[need]; ok {
			if i-j >= 2 {
				return true
			}
		}
	}
	return false
}

func main() {
	println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
	println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
}
