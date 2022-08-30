package main

import (
	"fmt"
)

//给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时
//，返回 true；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//输出：true
//解释：我们可以按以下顺序执行：
//push(1), push(2), push(3), push(4), pop() -> 4,
//push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//
//
// 示例 2：
//
//
//输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//输出：false
//解释：1 不能在 2 之前弹出。
//
//
//
//
// 提示：
//
//
// 1 <= pushed.length <= 1000
// 0 <= pushed[i] <= 1000
// pushed 的所有元素 互不相同
// popped.length == pushed.length
// popped 是 pushed 的一个排列
//
// Related Topics栈 | 数组 | 模拟
//
// 👍 258, 👎 0
//
//
//
//

// fengyuwusong 2022-08-31 00:30:07
//leetcode submit region begin(Prohibit modification and deletion)
func validateStackSequences(pushed []int, popped []int) bool {
	var ans []int

	for len(pushed) != 0 || len(popped) != 0 {
		if len(ans) == 0 && len(pushed) > 0 {
			ans = append(ans, pushed[0])
			pushed = pushed[1:]
			continue
		}
		if len(ans) > 0 && len(popped) > 0 && ans[len(ans)-1] == popped[0] {
			ans = ans[:len(ans)-1]
			popped = popped[1:]
			continue
		}
		if len(ans) > 0 && len(popped) > 0 && len(pushed) > 0 && ans[len(ans)-1] != popped[0] {
			ans = append(ans, pushed[0])
			pushed = pushed[1:]
			continue
		}
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
