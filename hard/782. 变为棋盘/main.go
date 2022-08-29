package main

import "fmt"

//一个 n x n 的二维网络 board 仅由 0 和 1 组成 。每次移动，你能任意交换两列或是两行的位置。
//
// 返回 将这个矩阵变为 “棋盘” 所需的最小移动次数 。如果不存在可行的变换，输出 -1。
//
// “棋盘” 是指任意一格的上下左右四个方向的值均与本身不同的矩阵。
//
//
//
// 示例 1:
//
//
//
//
//输入: board = [[0,1,1,0],[0,1,1,0],[1,0,0,1],[1,0,0,1]]
//输出: 2
//解释:一种可行的变换方式如下，从左到右：
//第一次移动交换了第一列和第二列。
//第二次移动交换了第二行和第三行。
//
//
// 示例 2:
//
//
//
//
//输入: board = [[0, 1], [1, 0]]
//输出: 0
//解释: 注意左上角的格值为0时也是合法的棋盘，也是合法的棋盘.
//
//
// 示例 3:
//
//
//
//
//输入: board = [[1, 0], [1, 0]]
//输出: -1
//解释: 任意的变换都不能使这个输入变为合法的棋盘。
//
//
//
//
// 提示：
//
//
// n == board.length
// n == board[i].length
// 2 <= n <= 30
// board[i][j] 将只包含 0或 1
//
// Related Topics位运算 | 数组 | 数学 | 矩阵
//
// 👍 55, 👎 0
//
//
//
//

// fengyuwusong 2022-08-23 00:22:09
//leetcode submit region begin(Prohibit modification and deletion)
func movesToChessboard(board [][]int) int {
	n := len(board)
	// 分别存储第一行和第一列出现1的个数
	rowSum := 0
	colSum := 0
	// 分别存储第一行和列同于当前索引的次数
	rowDiff := 0
	colDiff := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 进行异或运算判断是否符合棋盘标准（每个子矩阵的四个角为分别是 两个0和1 || 四个1 || 四个0） 由于出现次数均为偶数 故异或1...1(n次)
			if itob(board[0][0] ^ board[i][0] ^ board[0][j] ^ board[i][j]) {
				return -1
			}
		}
	}
	// 计算相关变量值
	for i := 0; i < n; i++ {
		rowSum += board[0][i]
		colSum += board[i][0]
		rowDiff += btoi(board[i][0] == i%2)
		colDiff += btoi(board[0][i] == i%2)
	}
	// 判断行和列出现 1 的次数是否符合棋盘标准 (偶数行列 1或0出现的次数是n/2; 奇数行列 1或0出现的次数是n/2或(n+1)/2)
	if n/2 > rowSum || rowSum > (n+1)/2 {
		return -1
	}
	if n/2 > colSum || colSum > (n+1)/2 {
		return -1
	}
	// 计算移动所需最小步数
	// 第一行列符合棋盘标准 则每一行列都符合棋盘标准
	// 符合条件的棋盘每一行列应该为01010 或1010 即偶数[1010, 0101] 则1和0相同 奇数行则1和0出现分别为奇数和偶数次[10101, 01010]

	// 如果棋盘为奇数次 则需要得到更改偶数的位差  这时的变更最小
	if itob(n % 2) {
		if itob(rowDiff % 2) {
			rowDiff = n - rowDiff
		}
		if itob(colDiff % 2) {
			colDiff = n - colDiff
		}
	} else {
		// 如果棋盘为偶数 则计算不同于行列的次数即可 这里需要取n-Diff的最小值 是因为本身可能就是正确的行列了
		// 如rowDiff计算的是0101的, board原本是1010也符合 则需取4-4=0
		rowDiff = min(n-rowDiff, rowDiff)
		colDiff = min(n-colDiff, colDiff)
	}
	return (rowDiff + colDiff) / 2
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(b int) bool {
	return b != 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	board := [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}}
	fmt.Println(movesToChessboard(board))
}
