package main

import "fmt"

/**
输入：initialEnergy = 5, initialExperience = 3, energy = [1,4,3,2], experience = [2,6,3,1]
输出：8
解释：在 6 小时训练后，你可以将精力提高到 11 ，并且再训练 2 个小时将经验提高到 5 。
按以下顺序与对手比赛：
- 你的精力与经验都超过第 0 个对手，所以获胜。
  精力变为：11 - 1 = 10 ，经验变为：5 + 2 = 7 。
- 你的精力与经验都超过第 1 个对手，所以获胜。
  精力变为：10 - 4 = 6 ，经验变为：7 + 6 = 13 。
- 你的精力与经验都超过第 2 个对手，所以获胜。
  精力变为：6 - 3 = 3 ，经验变为：13 + 3 = 16 。
- 你的精力与经验都超过第 3 个对手，所以获胜。
  精力变为：3 - 2 = 1 ，经验变为：16 + 1 = 17 。
在比赛前进行了 8 小时训练，所以返回 8 。
*/

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	var ans int
	// 计算需要的最大精力
	sumEnergy := 0
	for i := 0; i < len(energy); i++ {
		sumEnergy += energy[i]
	}
	// 如所需精力总和大于等于初始精力 则计算额外所需的精力
	if sumEnergy >= initialEnergy {
		ans += sumEnergy - initialEnergy + 1
	}
	// 计算需要的最大经验
	for i := 0; i < len(experience); i++ {
		if initialExperience <= experience[i] {
			ans += experience[i] - initialExperience + 1
			initialExperience = experience[i] + 1
		}
		initialExperience += experience[i]
	}
	return ans
}

func main() {
	fmt.Println(minNumberOfHours(5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1}))
	fmt.Println(minNumberOfHours(1, 1, []int{1, 1, 1, 1}, []int{1, 1, 1, 50}))
	fmt.Println(minNumberOfHours(5, 3, []int{1, 4}, []int{2, 5}))
}
