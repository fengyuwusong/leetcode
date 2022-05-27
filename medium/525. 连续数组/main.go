package main

// 0直接使用-1表示，则当preSum[i] - preSum[j] = 0 => preSum[i] = preSum[j] 说明有相同个数的0和1
func findMaxLength(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	count := make(map[int]int)
	// 构建前缀和数组
	for i := 1; i <= n; i++ {
		v := 1
		if nums[i-1] == 0 {
			v = -1
		}
		preSum[i] = preSum[i-1] + v
	}

	// 遍历记录每个前缀和第一次出现的最早索引
	for i := 0; i <= n; i++ {
		if _, ok := count[preSum[i]]; !ok {
			count[preSum[i]] = i
		}
	}

	res := 0

	for i := 1; i <= n; i++ {
		need := preSum[i]
		if j, ok := count[need]; ok && i-j > res {
			res = i - j
		}
	}
	return res
}

func main() {
	println(findMaxLength([]int{0, 0, 1, 0, 0, 0, 1, 1}))
}
