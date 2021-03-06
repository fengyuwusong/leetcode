package medium

import (
	"fmt"
	"testing"
)

/**
给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。



示例:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

输出:  [1,2,4,7,5,3,6,8,9]

解释:



说明:

给定矩阵中的元素总数不会超过 100000 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diagonal-traverse
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//解法1 对角线迭代并翻转
// 先将问题简化成简单的问题 在附加上额外的条件变为题目要求即可
// 1. 可将问题简化为求二维数组各个对角线元素 那么从 1 / 2 4 / 3 5 7 / 6 8 /9
// 可看出 给定元素 m[i][j] 那么其对角线上的元素必定为 m[i-1][j+1] 或 m[i+1][j-1]
// 2. 由观可知整个二维数组从右到左的对角开始元素分别为m[0][j] 和 m[i][col] 即第一行和最后一列的起始元素
// 有了以上两个推论我们即可以将二维数组的对角线依次输出，观察题目要求 对角线遍历要求奇数对角线顺序 偶数对角线翻转 则按要求翻转插入结果即可
// 时间复杂度 O(N* M) 序遍历二维数全部元素 空间复杂度O(min(N,M)) 需额外空间存储偶数对角线 并进行翻转

// 解法2 找规律模拟对角线遍历
// 1. 每一趟对角线中元素的坐标(x, y)相加的和是递增的 如
// 第一趟 1(0, 0) x+y=0
// 第二趟 2(0, 1) 4(1, 0) x+y=1
// ...
// 2. 每一趟都是x或y其中一个从大到小(每次-1), 另一个从小到大(每次+1)
// 第二趟 2(0, 1) 4(1, 0) x+1 y-1
// 第三趟 7(2, 0) 5(1, 1) 3(0, 2) x-1 y+1
// 3. 确定单次循环初始值。 如奇数趟是x从大到小，x则取最大，当初始值超过x的上限时，不足的部分加到y 偶数趟反之 y值取最小
// 第二趟 初始值 2(0, 1) x+y=1 y取最大 1<m (二维矩阵宽m=3)
// 第四趟 初始值 6(1, 2) x+y=3 y取最大值 2<m x=3-y=1
// 4. 确定单次循环结束值 奇数趟是x从大到小 则结束判断为x减到0或y加到上限 反之则为加到上限值
// 第二趟 2(0, 1) 4(1, 0) x 1 加到 x+y=1上限值
// 第四趟 6(1, 2) 8(2, 1) x 2 加到 x<n (n=3)最大值
// 5. 确定加减顺序。 奇数趟是x从大到小，偶数趟必定是从小到大 将两次加减循序看为一个循环
// 第一趟 1(0, 0) 从大到小
// 第二趟 2(0, 1) 4(1, 0) 从小到大
// 第三趟 7(2, 0) 5(1, 1) 3(0, 2) 从大到小
// ...
// 根据以上5点即可编写代码
// 时间复杂度 O(N*M) 空间复杂度O(1) 输出空间不计入空间复杂度
func findDiagonalOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return nil
	}
	col := len(matrix[0])
	if col == 0 {
		return nil
	}
	bxFlag := true // 用于判断奇偶数趟次 初始化为奇数趟 true 偶数false
	var ans []int
	// 当i = x+y< row+col 时进行遍历
	for i := 0; i < row+col; i++ {
		// 根据当前是奇数还是偶数趟初始化x y
		var x, y int
		if bxFlag {
			// 奇数
			x = row - 1
			if i < row {
				x = i
			}
			y = i - x
			// 当x<=0 或 y>=col时循环结束 反之 x>=0&&y<col时继续
			for x >= 0 && y < col {
				ans = append(ans, matrix[x][y])
				x--
				y++
			}
		} else {
			// 偶数
			y = col - 1
			if i < col {
				y = i
			}
			x = i - y
			// 当x>=row 或 y<0时循环结束 反之 x<row&&y>=0时继续
			for x < row && y >= 0 {
				ans = append(ans, matrix[x][y])
				x++
				y--
			}
		}

		// 更改奇偶趟flag
		bxFlag = !bxFlag
	}
	return ans
}

func TestFindDiagonalOrder(t *testing.T) {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
