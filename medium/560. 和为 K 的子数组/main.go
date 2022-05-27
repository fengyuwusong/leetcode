package main

func subarraySum(nums []int, k int) int {
	n := len(nums)
	// 前缀和数组
	preSum := make([]int, n+1)
	// 前缀和到索引的映射 方便快速查找所需的前缀和
	count := make(map[int]int)
	count[0] = 1
	// 记录和=k的子数组个数
	res := 0

	// 计算nums的前缀和
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		// 如果之前存在值为 need 的前缀和
		// 说明存在以 nums[i-1] 结尾的子数组和=k
		need := preSum[i] - k
		if val, ok := count[need]; ok {
			res += val
		}

		// 将当前前缀和存入哈希表
		count[preSum[i]]++
	}
	return res
}

func main() {
	print(subarraySum([]int{1, 1, 1}, 2))
}
