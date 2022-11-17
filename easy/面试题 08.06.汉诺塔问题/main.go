package main

import "fmt"

// author: fengyuwusong date: 2022-09-14 14:53:19

// 面试题 08.06. 汉诺塔问题
//在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上(即每一个盘子只
//能放在更大的盘子上面)。移动圆盘时受到以下限制:
//(1) 每次只能移动一个盘子;
//(2) 盘子只能从柱子顶端滑出移到下一根柱子;
//(3) 盘子只能叠在比它大的盘子上。
//
// 请编写程序，用栈将所有盘子从第一根柱子移到最后一根柱子。
//
// 你需要原地修改栈。
//
// 示例1:
//
//  输入：A = [2, 1, 0], B = [], C = []
// 输出：C = [2, 1, 0]
//
//
// 示例2:
//
//  输入：A = [1, 0], B = [], C = []
// 输出：C = [1, 0]
//
//
// 提示:
//
//
// A中盘子的数目不大于14个。
//
// Related Topics递归 | 数组
//
// 👍 177, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func hanota(A []int, B []int, C []int) []int {
	var move func(from, to, other *[]int, n int)
	move = func(from, to, other *[]int, n int) {
		if n == 1 {
			cur := (*from)[len(*from)-1]
			*to = append(*to, cur)
			*from = (*from)[:len(*from)-1]
			return
		}
		// 把n-1借助to移动到other
		move(from, other, to, n-1)
		cur := (*from)[len(*from)-1]
		// 把from最底层移动到to
		*to = append(*to, cur)
		*from = (*from)[:len(*from)-1]
		// 把n-1从other移动到from
		move(other, to, from, n-1)
	}
	move(&A, &C, &B, len(A))
	return C
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(hanota([]int{1, 2, 3}, []int{}, []int{}))
}
