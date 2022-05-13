package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1104 lang=golang
 *
 * [1104] 二叉树寻路
 */

// @lc code=start
// 完全二叉树可以通过子节点/2得到父节点
// 之字形的完全二叉树可以通过
// 当 ii 是奇数时，第 ii 行为按从左到右的顺序进行标记，因此该节点的「从左到右标号」的父节点 = lable/2
// 当 ii 是偶数时，第 ii 行为按从右到左的顺序进行标记，将整行的标号左右翻转之后得到按从左到右的顺序进行标记的标号
// 因此实际lable=2^(i-1)+2^i-1-lable/2
func getReverse(label, row int) int {
	return 1<<(row-1) + 1<<row - 1 - label
}

func pathInZigZagTree(label int) (path []int) {
	row, rowStart := 1, 1
	for rowStart*2 <= label {
		row++
		rowStart *= 2
	}
	// 第一个节点偶数层无需翻转 故先提前翻转避免for循环翻转
	if row%2 == 0 {
		label = getReverse(label, row)
	}
	for row > 0 {
		if row%2 == 0 {
			path = append(path, getReverse(label, row))
		} else {
			path = append(path, label)
		}
		row--
		label >>= 1
	}
	for i, n := 0, len(path); i < n/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
	}
	return
}

// @lc code=end

func main() {
	fmt.Printf("pathInZigZagTree(11): %v\n", pathInZigZagTree(11))
}
