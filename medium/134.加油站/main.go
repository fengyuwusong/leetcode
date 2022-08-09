package main

import "fmt"

// author: fengyuwusong date: 2022-08-08 11:37:09

// 134. 加油站
//在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
// 给定两个整数数组 gas 和 cost ，如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
//
//
//
// 示例 1:
//
//
//输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
//输出: 3
//解释:
//从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
//开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
//开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
//开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
//开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
//开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
//因此，3 可为起始索引。
//
// 示例 2:
//
//
//输入: gas = [2,3,4], cost = [3,4,3]
//输出: -1
//解释:
//你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
//我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
//开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
//开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
//你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
//因此，无论怎样，你都不可能绕环路行驶一周。
//
//
//
// 提示:
//
//
// gas.length == n
// cost.length == n
// 1 <= n <= 10⁵
// 0 <= gas[i], cost[i] <= 10⁴
//
// Related Topics贪心 | 数组
//
// 👍 1005, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	// 当前油量 最小油量 最小油量的下一个加油站
	var sum, minSum, start int
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
		if sum < minSum {
			minSum = sum
			start = i + 1
		}
	}
	// 一轮遍历后 当前油量小于0 说明无解
	if sum < 0 {
		return -1
	}
	// 返回起点即可
	if start == n {
		return 0
	}
	return start
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
