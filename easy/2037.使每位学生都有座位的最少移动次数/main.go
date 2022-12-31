package main

import (
	"fmt"
	"sort"
)

//一个房间里有 n 个座位和 n 名学生，房间用一个数轴表示。给你一个长度为 n 的数组 seats ，其中 seats[i] 是第 i 个座位的位置。同时给
//你一个长度为 n 的数组 students ，其中 students[j] 是第 j 位学生的位置。
//
// 你可以执行以下操作任意次：
//
//
// 增加或者减少第 i 位学生的位置，每次变化量为 1 （也就是将第 i 位学生从位置 x 移动到 x + 1 或者 x - 1）
//
//
// 请你返回使所有学生都有座位坐的 最少移动次数 ，并确保没有两位学生的座位相同。
//
// 请注意，初始时有可能有多个座位或者多位学生在 同一 位置。
//
//
//
// 示例 1：
//
// 输入：seats = [3,1,5], students = [2,7,4]
//输出：4
//解释：学生移动方式如下：
//- 第一位学生从位置 2 移动到位置 1 ，移动 1 次。
//- 第二位学生从位置 7 移动到位置 5 ，移动 2 次。
//- 第三位学生从位置 4 移动到位置 3 ，移动 1 次。
//总共 1 + 2 + 1 = 4 次移动。
//
//
// 示例 2：
//
// 输入：seats = [4,1,5,9], students = [1,3,2,6]
//输出：7
//解释：学生移动方式如下：
//- 第一位学生不移动。
//- 第二位学生从位置 3 移动到位置 4 ，移动 1 次。
//- 第三位学生从位置 2 移动到位置 5 ，移动 3 次。
//- 第四位学生从位置 6 移动到位置 9 ，移动 3 次。
//总共 0 + 1 + 3 + 3 = 7 次移动。
//
//
// 示例 3：
//
// 输入：seats = [2,2,6,6], students = [1,3,2,6]
//输出：4
//解释：学生移动方式如下：
//- 第一位学生从位置 1 移动到位置 2 ，移动 1 次。
//- 第二位学生从位置 3 移动到位置 6 ，移动 3 次。
//- 第三位学生不移动。
//- 第四位学生不移动。
//总共 1 + 3 + 0 + 0 = 4 次移动。
//
//
//
//
// 提示：
//
//
// n == seats.length == students.length
// 1 <= n <= 100
// 1 <= seats[i], students[j] <= 100
//
//
// Related Topics 数组 排序 👍 52 👎 0

// author: fengyuwusong
// create time: 2022-12-31 22:42:30
// leetcode submit region begin(Prohibit modification and deletion)
func minMovesToSeat(seats []int, students []int) int {
	// 将两个数组都进行排序，则每个对应就是其需要移动的最小次数
	sort.Ints(seats)
	sort.Ints(students)
	var ans int
	for i := 0; i < len(seats); i++ {
		ans += abs(students[i] - seats[i])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
	fmt.Println(minMovesToSeat([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
	fmt.Println(minMovesToSeat([]int{2, 2, 6, 6}, []int{1, 3, 2, 6}))
}
