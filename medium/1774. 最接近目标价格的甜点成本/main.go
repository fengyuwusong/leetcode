package main

import "fmt"

//你打算做甜点，现在需要购买配料。目前共有 n 种冰激凌基料和 m 种配料可供选购。而制作甜点需要遵循以下几条规则：
//
//
// 必须选择 一种 冰激凌基料。
// 可以添加 一种或多种 配料，也可以不添加任何配料。
// 每种类型的配料 最多两份 。
//
//
// 给你以下三个输入：
//
//
// baseCosts ，一个长度为 n 的整数数组，其中每个 baseCosts[i] 表示第 i 种冰激凌基料的价格。
// toppingCosts，一个长度为 m 的整数数组，其中每个 toppingCosts[i] 表示 一份 第 i 种冰激凌配料的价格。
// target ，一个整数，表示你制作甜点的目标价格。
//
//
// 你希望自己做的甜点总成本尽可能接近目标价格 target 。
//
// 返回最接近 target 的甜点成本。如果有多种方案，返回 成本相对较低 的一种。
//
//
//
// 示例 1：
//
//
//输入：baseCosts = [1,7], toppingCosts = [3,4], target = 10
//输出：10
//解释：考虑下面的方案组合（所有下标均从 0 开始）：
//- 选择 1 号基料：成本 7
//- 选择 1 份 0 号配料：成本 1 x 3 = 3
//- 选择 0 份 1 号配料：成本 0 x 4 = 0
//总成本：7 + 3 + 0 = 10 。
//
//
// 示例 2：
//
//
//输入：baseCosts = [2,3], toppingCosts = [4,5,100], target = 18
//输出：17
//解释：考虑下面的方案组合（所有下标均从 0 开始）：
//- 选择 1 号基料：成本 3
//- 选择 1 份 0 号配料：成本 1 x 4 = 4
//- 选择 2 份 1 号配料：成本 2 x 5 = 10
//- 选择 0 份 2 号配料：成本 0 x 100 = 0
//总成本：3 + 4 + 10 + 0 = 17 。不存在总成本为 18 的甜点制作方案。
//
//
// 示例 3：
//
//
//输入：baseCosts = [3,10], toppingCosts = [2,5], target = 9
//输出：8
//解释：可以制作总成本为 8 和 10 的甜点。返回 8 ，因为这是成本更低的方案。
//
//
// 示例 4：
//
//
//输入：baseCosts = [10], toppingCosts = [1], target = 1
//输出：10
//解释：注意，你可以选择不添加任何配料，但你必须选择一种基料。
//
//
//
// 提示：
//
//
// n == baseCosts.length
// m == toppingCosts.length
// 1 <= n, m <= 10
// 1 <= baseCosts[i], toppingCosts[i] <= 10⁴
// 1 <= target <= 10⁴
//
// Related Topics数组 | 动态规划 | 回溯
//
// 👍 135, 👎 0
//
//
//
//

// fengyuwusong 2022-12-04 23:43:03
//leetcode submit region begin(Prohibit modification and deletion)
// 1. dp[i] 表示该花费是否可被组成
// 2. 大于upper = 2*target-min的花费, 其与target的差值一定大于target和min的差值(大了一个target min<=target必成立) 故dp区间选择[min, upper-1]
// 3. 如果baseCost[i] < target 那么 dp[baseCosts[i]] = true
// 4. dp数组迭代辅料 如果dp[i] = true 那么加上该辅料1-2份的花费小于upper, 则dp[i] = true
// 5. dp数组需倒序遍历，避免重复成本的时候重复
// 6. 最后遍历dp[i]找到最接近的即可
// 时间复杂度O(len(toppingCosts) * n)
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	min := baseCosts[0]
	for _, c := range baseCosts {
		min = minFun(min, c)
	}
	if min >= target {
		return min
	}
	// 考虑[min, upper-1] 方案是否存在
	upper := 2*target - min
	dp := make([]bool, upper)
	// 选择基料
	for _, baseCost := range baseCosts {
		// 不考虑区间外的方案
		if baseCost < upper {
			dp[baseCost] = true
		}
	}
	// dp迭代 选择辅料
	for _, toppingCost := range toppingCosts {
		// 倒叙遍历dp
		for j := upper - 1; j >= min; j-- {
			// 每种辅料可以1-2，更新对应dp为true 区间外不考虑
			if dp[j] && j+toppingCost < upper {
				dp[j+toppingCost] = true
			}
			if dp[j] && j+2*toppingCost < upper {
				dp[j+2*toppingCost] = true
			}
		}
	}

	// 在[min, upper-1] 所有存在的方案中找到与target最接近的方案
	ans := min
	for i := min + 1; i < upper; i++ {
		if dp[i] {
			// 当前值更接近target
			if abs(i-target) < abs(ans-target) {
				ans = i
			} else if abs(i-target) == abs(ans-target) {
				// 同样接近 选择更小成本
				ans = minFun(ans, i)
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minFun(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

/**
// 回溯解法
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	ans := baseCosts[0]
	for _, c := range baseCosts {
		ans = min(ans, c)
	}
	var dfs func(int, int)
	dfs = func(p int, curCost int) {
		// 当前方案的差值大于已有方案的差值
		// 后续选择只会使得成本更大，与target的差值更大 直接剪枝
		if abs(ans-target) < curCost-target {
			return
		} else if abs(ans-target) > abs(curCost-target) {
			// 当前方案差值小于已有最优方案 则当前方案最优
			ans = curCost
		} else if abs(ans-target) == abs(curCost-target) {
			// 当前方案差值等于已有最优方案 则选取最小的继续回溯
			ans = min(ans, curCost)
		}
		if p == len(toppingCosts) {
			return
		}
		dfs(p+1, curCost+toppingCosts[p]*2)
		dfs(p+1, curCost+toppingCosts[p])
		dfs(p+1, curCost)
	}

	for _, c := range baseCosts {
		dfs(0, c)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
*/

func main() {
	fmt.Println(closestCost([]int{1, 7}, []int{3, 4}, 10))
}
