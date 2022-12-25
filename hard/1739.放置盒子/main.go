package main

import "fmt"

//有一个立方体房间，其长度、宽度和高度都等于 n 个单位。请你在房间里放置 n 个盒子，每个盒子都是一个单位边长的立方体。放置规则如下：
//
//
// 你可以把盒子放在地板上的任何地方。
// 如果盒子 x 需要放置在盒子 y 的顶部，那么盒子 y 竖直的四个侧面都 必须 与另一个盒子或墙相邻。
//
//
// 给你一个整数 n ，返回接触地面的盒子的 最少 可能数量。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 3
//输出：3
//解释：上图是 3 个盒子的摆放位置。
//这些盒子放在房间的一角，对应左侧位置。
//
//
// 示例 2：
//
//
//
//
//输入：n = 4
//输出：3
//解释：上图是 3 个盒子的摆放位置。
//这些盒子放在房间的一角，对应左侧位置。
//
//
// 示例 3：
//
//
//
//
//输入：n = 10
//输出：6
//解释：上图是 10 个盒子的摆放位置。
//这些盒子放在房间的一角，对应后方位置。
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁹
//
//
// Related Topics 贪心 数学 二分查找 👍 35 👎 0

// author: fengyuwusong
// create time: 2022-12-25 00:21:29
// leetcode submit region begin(Prohibit modification and deletion)
// 3 6 10
func minimumBoxes(n int) int {
	// curr 目前所放置的盒子数 i 为当前 j 层的盒子数 j 为层数
	curr, i, j := 0, 0, 0
	// 计算出需要多少层 最大层数所能带来的盒子数 = 1 + (1+2) + (1+2+3) + ... + (1+...+i)
	for curr < n {
		j++
		i += j
		curr += i
	}

	// 如果正好等于 则 i 就是底层所放置的盒子数
	if curr == n {
		return i
	}

	// 因为 curr > n 说明放多 要减去放的 i 同时 i 退回上一层
	curr -= i
	i -= j
	// j 次数代表继续放置盒子第几次 第 j 次可以放 j 个 每次增加盒子数等于1+2+...+j
	j = 0
	for curr < n {
		j++
		curr += j
	}
	// 一共放了 j 次最底层 +j 个盒子
	return i + j
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minimumBoxes(3))
	fmt.Println(minimumBoxes(4))
	fmt.Println(minimumBoxes(14))
}
