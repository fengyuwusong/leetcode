package main

import (
	"fmt"
	"math"
	"sort"
)

//给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
//
// 现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
//
// 给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
//
//
//
// 示例：
//
//
//输入：[[1,2], [2,3], [3,4]]
//输出：2
//解释：最长的数对链是 [1,2] -> [3,4]
//
//
//
//
// 提示：
//
//
// 给出数对的个数在 [1, 1000] 范围内。
//
// Related Topics贪心 | 数组 | 动态规划 | 排序
//
// 👍 297, 👎 0
//
//
//
//

// fengyuwusong 2022-09-03 12:53:44
//leetcode submit region begin(Prohibit modification and deletion)
// 贪心: 挑选第一个满足前提的瞧见下，挑选第二个最小的
func findLongestChain(pairs [][]int) int {
	// pairs按照第2位进行升序排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	cur := math.MinInt
	var ans int
	for _, p := range pairs {
		// 可选
		if p[0] > cur {
			cur = p[1]
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
}
