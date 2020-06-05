package medium

import (
	"fmt"
	"testing"
)

/**
  每日温度
根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
*/

// 解法1: 暴力解
// 解法2： 使用递减栈记录值 当当前温度大于栈顶时出栈并计算差值 否则将当天日期入栈
func dailyTemperatures(T []int) []int {
	if len(T) == 0 {
		return nil
	}
	ans := make([]int, len(T))
	// 利用递减栈辅助解
	var stack []int
	for i, v := range T {
		// 遍历栈 将栈元素出栈 计算栈元素对应差值日期保存结果
		// 当栈为空或者当前温度大于栈顶元素温度时, 退出遍历将索引入栈
		for len(stack) != 0 && T[stack[len(stack)-1]] < v {
			index := stack[len(stack)-1]
			ans[index] = i - index
			stack = stack[:len(stack)-1]
		}
		// 将当前元素入栈
		stack = append(stack, i)
	}
	return ans
}

func TestDailyTemperatures(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{73, 74}))
}
