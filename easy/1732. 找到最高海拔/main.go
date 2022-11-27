package main

import "fmt"

//有一个自行车手打算进行一场公路骑行，这条路线总共由 n + 1 个不同海拔的点组成。自行车手从海拔为 0 的点 0 开始骑行。
//
// 给你一个长度为 n 的整数数组 gain ，其中 gain[i] 是点 i 和点 i + 1 的 净海拔高度差（0 <= i < n）。请你返回 最高点的
//海拔 。
//
//
//
// 示例 1：
//
//
//输入：gain = [-5,1,5,0,-7]
//输出：1
//解释：海拔高度依次为 [0,-5,-4,1,1,-6] 。最高海拔为 1 。
//
//
// 示例 2：
//
//
//输入：gain = [-4,-3,-2,-1,4,3,2]
//输出：0
//解释：海拔高度依次为 [0,-4,-7,-9,-10,-6,-3,-1] 。最高海拔为 0 。
//
//
//
//
// 提示：
//
//
// n == gain.length
// 1 <= n <= 100
// -100 <= gain[i] <= 100
//
// Related Topics数组 | 前缀和
//
// 👍 62, 👎 0
//
//
//
//

// fengyuwusong 2022-11-22 01:27:37
//leetcode submit region begin(Prohibit modification and deletion)
func largestAltitude(gain []int) int {
	var ans int
	currentAltitude := 0
	for i := 0; i < len(gain); i++ {
		currentAltitude += gain[i]
		ans = max(ans, currentAltitude)
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
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
