package algorithm

// 前缀和算法

func preSum(nums []int) {
	// 前缀和数组
	preSum := make([]int, len(nums)+1)
	// 哈希表记录前缀和出现次数  用于一些需要统计前缀出现次数的场景
	count := make(map[int]int)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]

		count[preSum[i]]++

		// 计算某个数组和总数出现的次数
		//need := preSum[i] - k
		// preSum[need]的次数即为总数和出现的次数
	}

	// 查询闭区间 [i, j] 的累加和
	// prefix[j+1] - prefix[i]
}
